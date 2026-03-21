package models

type Pin struct {
	Name string
	Tag  string
}

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
