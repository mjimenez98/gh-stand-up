package github

import (
	"fmt"
	"net/url"

	"github.com/mjimenez98/gh-stand-up/internal/helpers"
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
func (c *Client) GetOpenedIssues(userLogin string) ([]Issue, error) {
	query := url.Values{}
	query.Add("q", fmt.Sprintf("is:issue author:%s created:%s", userLogin, helpers.GetYesterdayDate()))
	query.Add("per_page", "30")
	path := fmt.Sprintf("search/issues?%s", query.Encode())

	var response IssueSearch
	err := c.Client.Get(path, &response)
	if err != nil {
		return nil, fmt.Errorf("Error searching issues: %w", err)
	}

	return response.Issues, nil
}
