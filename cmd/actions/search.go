package actions

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"slices"
	"sort"
	"strings"

	"github.com/atticus64/dona/cmd/models"
	"github.com/fatih/color"
	"github.com/gernest/wow"
	"github.com/gernest/wow/spin"

	"github.com/spf13/cobra"
)

func SearchDotfiles(query string, page int) ([]models.Repository, error) {

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
	fmt.Println(color.YellowString("Found"), result.Total_count, "dotfiles repositories")
	fmt.Println(color.MagentaString("Page"), page, "of", result.Total_count / 10 )
	fmt.Println()

	return result.Items, nil
}

func SortViaStars(repositories []models.Repository) {
	sort.Slice(repositories, func(i, j int) bool {
		return repositories[i].Stargazers_count > repositories[j].Stargazers_count
	})
}

func SortViaName(repositories []models.Repository) {
	sort.Slice(repositories, func(i, j int) bool {
		return repositories[i].FullName < repositories[j].FullName
	})
}

var SearchCmd = &cobra.Command{
	Use:   "search [query string]",
	Short: "Search across dotfiles in github",
	Long:  `Search in github repositories the dotfiles repos with match query`,
	Args:  cobra.MinimumNArgs(1),
	SuggestFor: []string{
		"find", "lookup",
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
		type_sort, err := cmd.Flags().GetString("sort")
		if err != nil {
			fmt.Println(err)
			return
		}

		sorts := []string{"stars", "name"}

		if type_sort != "" && !slices.Contains(sorts, type_sort) {
			fmt.Println(color.RedString("Error:"), "Sort invalid")
			return
		}

		repositories, error := SearchDotfiles(args[0], page)

		if type_sort != "" {
			switch type_sort {
			case "name":
				SortViaName(repositories)
			case "stars":
				SortViaStars(repositories)
			default:
				fmt.Println(color.RedString("Error:"), "Sort invalid")
				return
			}
		}

		if error != nil {
			fmt.Println(error)
			return
		}

		for _, repo := range repositories {
			fmt.Println(color.RedString("Name:"), repo.FullName)
			fmt.Println(color.BlueString("Url:"), repo.Html_url)
			fmt.Println(color.GreenString("Description:"), repo.Description)
			fmt.Println(color.YellowString("Stars:"), repo.Stargazers_count)
			fmt.Println()
		}
	},
}
