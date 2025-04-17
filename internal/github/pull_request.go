package github

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/mjimenez98/gh-stand-up/internal/helpers"
)

// PullRequestSearch represents the response from the GitHub API for searching pull requests.
type PullRequestSearch struct {
	TotalCount   int           `json:"total_count"`
	PullRequests []PullRequest `json:"items"`
}

// PullRequest represents a GitHub pull request.
type PullRequest struct {
	Author        User   `json:"user"`
	Number        int    `json:"number"`
	Title         string `json:"title"`
	RepositoryUrl string `json:"repository_url"`
	URL           string `json:"html_url"`
}

// Example repository URL: "https://api.github.com/repos/batterseapower/pinyin-toolkit"
const reposPrefix = "/repos/"

// GetOpenedPullRequests retrieves the pull requests opened the day before by the given user.
func (c *Client) GetOpenedPullRequests(userLogin string) ([]PullRequest, error) {
	query := url.Values{}
	query.Add("q", fmt.Sprintf("is:pr author:%s created:%s", userLogin, helpers.GetYesterdayDate()))
	query.Add("per_page", "30")
	path := fmt.Sprintf("search/issues?%s", query.Encode())

	var response PullRequestSearch
	err := c.Client.Get(path, &response)
	if err != nil {
		return nil, fmt.Errorf("error searching opened pull requests: %w", err)
	}

	return response.PullRequests, nil
}

// GetOpenPullRequests retrieves the open pull requests for the given user.
func (p *Client) GetOpenPullRequests(userLogin string) ([]PullRequest, error) {
	query := url.Values{}
	query.Add("q", fmt.Sprintf("is:pr author:%s is:open", userLogin))
	query.Add("per_page", "30")
	path := fmt.Sprintf("search/issues?%s", query.Encode())

	var response PullRequestSearch
	err := p.Client.Get(path, &response)
	if err != nil {
		return nil, fmt.Errorf("error searching open pull requests: %w", err)
	}

	return response.PullRequests, nil
}

// parseRepositoryUrl extracts the repository owner and name from the repository URL.
// It assumes the URL is in the format "https://api.github.com/repos/{owner}/{repo}".
func (p *PullRequest) repoWithOwner() string {
	index := strings.Index(p.RepositoryUrl, reposPrefix)
	if index != -1 {
		substring := p.RepositoryUrl[index+len(reposPrefix):]
		return substring
	}

	return ""
}
