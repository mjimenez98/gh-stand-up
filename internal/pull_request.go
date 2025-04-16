package github

import (
	"fmt"
	"net/url"
	"time"
)

// PullRequestSearch represents the response from the GitHub API for searching pull requests.
type PullRequestSearch struct {
	TotalCount   int           `json:"total_count"`
	PullRequests []PullRequest `json:"items"`
}

// PullRequest represents a GitHub pull request.
type PullRequest struct {
	Title string `json:"title"`
	URL   string `json:"html_url"`
}

// GetOpenedPullRequests retrieves the pull requests opened the day before by the given user.
func (c *Client) GetOpenedPullRequests(userLogin string) ([]PullRequest, error) {
	currentTime := time.Now()
	yesterday := currentTime.AddDate(0, 0, -1).Format("2006-01-02")

	query := url.Values{}
	query.Add("q", fmt.Sprintf("is:pr author:%s created:%s", userLogin, yesterday))
	query.Add("per_page", "30")
	path := fmt.Sprintf("search/issues?%s", query.Encode())

	var response PullRequestSearch
	err := c.Client.Get(path, &response)
	if err != nil {
		return nil, fmt.Errorf("Error searching pull requests: %w", err)
	}

	return response.PullRequests, nil
}
