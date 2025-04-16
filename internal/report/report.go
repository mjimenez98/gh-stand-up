// Package report
package report

import (
	"fmt"

	"github.com/mjimenez98/gh-stand-up/internal/github"
)

// GenerateYestedayReport generates a report of the user's activity for the previous day.
func GenerateYestedayReport(openedIssues []github.Issue, openedPullRequests []github.PullRequest) string {
	update := "This is what you did yesterday:\n"

	for _, issue := range openedIssues {
		update += fmt.Sprintf("📋 Opened %s - %s\n", issue.Title, issue.URL)
	}
	for _, pull_request := range openedPullRequests {
		update += fmt.Sprintf("🔀 Opened %s - %s\n", pull_request.Title, pull_request.URL)
	}

	return update
}
