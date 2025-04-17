// Package report provides functionality to generate reports based on GitHub activity.
package report

import (
	"fmt"

	"github.com/mjimenez98/gh-stand-up/internal/github"
)

// GenerateYestedayReport generates a report of the user's activity for the previous day.
func GenerateYestedayReport(openedIssues []github.Issue, openedPullRequests []github.PullRequest, pullRequestsReviewed []github.PullRequest) string {
	update := "This is what you did yesterday:\n"

	for _, pull_request := range pullRequestsReviewed {
		update += fmt.Sprintf("🔍 Reviewed %s - %s\n", pull_request.Title, pull_request.URL)
	}
	for _, pull_request := range openedPullRequests {
		update += fmt.Sprintf("🔀 Opened %s - %s\n", pull_request.Title, pull_request.URL)
	}
	for _, issue := range openedIssues {
		update += fmt.Sprintf("📋 Opened %s - %s\n", issue.Title, issue.URL)
	}

	return update
}
