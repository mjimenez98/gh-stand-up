package github

import (
	"fmt"
	"net/url"
	"time"
)

// IssueSearch represents the response from the GitHub API for searching issues.
type IssueSearch struct {
	TotalCount int     `json:"total_count"`
	Issues     []Issue `json:"items"`
}

// Issue represents a GitHub issue.
type Issue struct {
	Title string `json:"title"`
	URL   string `json:"html_url"`
}

// GetOpenedIssues retrieves the issues opened the day before by the given user.
func (c *Client) GetOpenedIssues(userLogin string) []Issue {
	currentTime := time.Now()
	yesterday := currentTime.AddDate(0, 0, -1).Format("2006-01-02")

	query := url.Values{}
	query.Add("q", fmt.Sprintf("is:issue author:%s created:%s", userLogin, yesterday))
	query.Add("per_page", "30")
	path := fmt.Sprintf("search/issues?%s", query.Encode())

	var response IssueSearch
	err := c.Client.Get(path, &response)
	if err != nil {
		fmt.Println("Error searching issues:", err)
		return nil
	}

	return response.Issues
}
