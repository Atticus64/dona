package actions

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/atticus64/dona/cmd/models"
	"github.com/fatih/color"
	"github.com/gernest/wow"
	"github.com/gernest/wow/spin"

	"github.com/spf13/cobra"
)

func searchDotfiles(query string, page int) ([]models.Repository, error) {

	parsedQuery := strings.Join(strings.Split(query, " "), "+")
	q := fmt.Sprintf("dotfiles+%s", parsedQuery)
	uri := fmt.Sprintf("https://api.github.com/search/repositories?q=%s&page=%d&per_page=10", q, page)
	w := wow.New(os.Stdout, spin.Get(spin.Moon), " Searching in github")
	w.Start()

	resp, err := http.Get(uri)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	var result models.Response
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	w.PersistWith(spin.Spinner{Frames: []string{"🐢"}}, " Nice!")

	return result.Items, nil
}

var SearchCmd = &cobra.Command{
	Use:   "search [query string]",
	Short: "Search across dotfiles in github",
	Long:  `Search in github repositories the dotfiles repos with match query`,
	Args:  cobra.MinimumNArgs(1),
	SuggestFor: []string{
		"find",
	},
	Example: `
	dona search "arch linux aesthetic"
	dona search fedora --page 4
	dona search cat -p 2
	`,
	Run: func(cmd *cobra.Command, args []string) {
		page, err := cmd.Flags().GetInt("page")
		if err != nil {
			fmt.Println(err)
			return
		}

		result, error := searchDotfiles(args[0], page)

		if error != nil {
			fmt.Println(error)
			return
		}

		for _, repo := range result {
			fmt.Println(color.RedString("Name:"), repo.FullName)
			fmt.Println(color.BlueString("Url:"), repo.Html_url)
			fmt.Println(color.GreenString("Description:"), repo.Description)
			fmt.Println(color.YellowString("Stars:"), repo.Stargazers_count)
			fmt.Println()
		}
	},
}
