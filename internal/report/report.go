// Package report
package report

import (
	"fmt"

	"github.com/mjimenez98/gh-stand-up/internal/github"
)

// GenerateYestedayReport generates a report of the user's activity for the previous day.
func GenerateYestedayReport(user string, openedIssues []github.Issue, openedPullRequests []github.PullRequest) string {
	update := "This is what you did yesterday:\n"

	for _, issue := range openedIssues {
		update += fmt.Sprintf("- You opened the issue: %s\n", issue.URL)
	}
	for _, pull_request := range openedPullRequests {
		update += fmt.Sprintf("- You opened the pull request: %s\n", pull_request.URL)
	}

	return update
}
