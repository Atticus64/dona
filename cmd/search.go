package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/gernest/wow"
	"github.com/gernest/wow/spin"

	"github.com/spf13/cobra"
)

type Repository struct {
	Name             string `json:"name"`
	FullName         string `json:"full_name"`
	Html_url         string `json:"html_url"`
	Description      string `json:"description"`
	Stargazers_count int    `json:"stargazers_count"`
}

type Response struct {
	Items       []Repository `json:"items"`
	Total_count int          `json:"total_count"`
}

func searchDotfiles(query string) ([]Repository, error) {

	parsedQuery := strings.Join(strings.Split(query, " "), "+")
	q := fmt.Sprintf("dotfiles+%s", parsedQuery)
	uri := fmt.Sprintf("https://api.github.com/search/repositories?q=%s", q)
	w := wow.New(os.Stdout, spin.Get(spin.Moon), " Searching in github")
	w.Start()

	resp, err := http.Get(uri)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	var result Response
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
	Use:     "search [query string]",
	Short:   "Search across dotfiles in github",
	Long:    `Search in github repositories the dotfiles repos with match query`,
	Example: `dona search "arch linux aesthetic"`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) <= 0 {
			cmd.Help()
			return
		}

		result, error := searchDotfiles(args[0])

		if error != nil {
			fmt.Println(error)
			return
		}

		for _, repo := range result {
			fmt.Println(color.RedString("Name:"), repo.FullName)
			fmt.Println(color.BlueString("Url:"), repo.Html_url)
			fmt.Println(color.YellowString("Description:"), repo.Description)
			fmt.Println(color.GreenString("-----------------------------------------"))
			fmt.Println()
		}
	},
}
