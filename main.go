package main

import (
	"fmt"
	"net/url"
	"time"

	"github.com/cli/go-gh/v2/pkg/api"
)

func main() {
	// Create a new GitHub API client
	client, err := api.DefaultRESTClient()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Obtain username
	response := struct {
		Login string
	}{}
	err = client.Get("user", &response)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Search for issues created by the user
	issues := struct {
		TotalCount int `json:"total_count"`
		Items      []struct {
			Title string `json:"title"`
			URL   string `json:"html_url"`
		} `json:"items"`
	}{}

	currentTime := time.Now()
	yesterday := currentTime.AddDate(0, 0, -1).Format("2006-01-02")

	query := url.Values{}
	query.Add("q", fmt.Sprintf("is:issue author:%s created:%s", response.Login, yesterday))
	query.Add("per_page", "30")
	path := fmt.Sprintf("search/issues?%s", query.Encode())

	err = client.Get(path, &issues)
	if err != nil {
		fmt.Println("Error searching issues:", err)
		return
	}

	// Generate report
	fmt.Printf("Hi %s, here is your report for issues you created yesterday:\n", response.Login)
	fmt.Printf("Total Issues: %d\n", issues.TotalCount)
	for _, issue := range issues.Items {
		fmt.Printf("- %s (%s)\n", issue.Title, issue.URL)
	}
}

// For more examples of using go-gh, see:
// https://github.com/cli/go-gh/blob/trunk/example_gh_test.go
