package github

import (
	"fmt"
	"net/url"
)

type PullRequestReview struct {
	PullRequest PullRequest
	Repository  string
}

// GetPullRequestsReviewed retrieves the pull requests reviewed the day before by the given user.
func (c *Client) GetPullRequestsReviewed(userLogin string) ([]PullRequest, error) {
	query := url.Values{}
	query.Add("q", fmt.Sprintf("is:pr reviewed-by:%s", userLogin))
	query.Add("per_page", "30")
	path := fmt.Sprintf("search/issues?%s", query.Encode())

	var response PullRequestSearch
	err := c.Client.Get(path, &response)
	if err != nil {
		return nil, fmt.Errorf("Error searching pull request reviews: %w", err)
	}

	return response.PullRequests, nil
}
