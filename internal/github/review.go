package github

import (
	"fmt"
	"net/url"
	"time"

	"github.com/mjimenez98/gh-stand-up/internal/helpers"
)

// PullRequestReviewSearch represents the response from the GitHub API for pull request reviews.
type PullRequestReviewSearch struct {
	Items []PullRequestReview `json:"items"`
}

// PullRequestReview represents a review of a pull request.
type PullRequestReview struct {
	Author User `json:"user"`
	// The submittedAt field is in RFC3339 format, e.g. "2023-10-01T12:34:56Z".
	SubmittedAt string `json:"submitted_at"`
}

// GetPullRequestsReviewed retrieves the pull requests reviewed the day before by the given user.
func (c *Client) GetPullRequestsReviewed(userLogin string) ([]PullRequest, error) {
	query := url.Values{}
	query.Add("q", fmt.Sprintf("is:pr reviewed-by:%s sort:updated", userLogin))
	query.Add("per_page", "10")
	path := fmt.Sprintf("search/issues?%s", query.Encode())

	var response PullRequestSearch
	err := c.Client.Get(path, &response)
	if err != nil {
		return nil, fmt.Errorf("error searching pull request reviews: %w", err)
	}

	// Filter out pull requests that were not created yesterday.
	// GitHub API does not provide a direct way to filter by date reviewed.
	var filteredPullRequests []PullRequest
	for _, pullRequest := range response.PullRequests {
		ownerRepo := pullRequest.repoWithOwner()
		query := fmt.Sprintf("repos/%s/pulls/%d/reviews", ownerRepo, pullRequest.Number)

		var reviews []PullRequestReview
		err := c.Client.Get(query, &reviews)
		if err != nil {
			return nil, fmt.Errorf("error getting pull request for review: %w", err)
		}

		// GitHub API does not guarantee the order of reviews, so we need to check each review.
		for _, review := range reviews {
			if review.Author.Login == userLogin && review.submittedAtFormatted() == helpers.GetYesterdayDate() {
				filteredPullRequests = append(filteredPullRequests, pullRequest)
			}
		}
	}

	return filteredPullRequests, nil
}

// submittedAtDateTime returns the submittedAt date in the format "YYYY-MM-DD".
func (r PullRequestReview) submittedAtFormatted() string {
	parsedTime, err := time.Parse(time.RFC3339, r.SubmittedAt)
	if err != nil {
		return ""
	}

	return parsedTime.Format("2006-01-02")
}
