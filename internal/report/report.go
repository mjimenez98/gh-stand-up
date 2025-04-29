// Package report provides functionality to generate reports based on GitHub activity.
package report

import (
	"fmt"

	"github.com/mjimenez98/gh-stand-up/internal/github"
)

// GenerateYesterdayReport generates a report of the user's activity for the previous day.
func GenerateYesterdayReport(openedIssues []github.Issue, openedPullRequests []github.PullRequest, pullRequestsReviewed []github.PullRequest) string {
	if len(openedIssues) == 0 && len(openedPullRequests) == 0 && len(pullRequestsReviewed) == 0 {
		return "No activity detected yesterday.\n"
	}

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

// GenerateTodayReport generates a report of what the user should address today.
func GenerateTodayReport(openPullRequests []github.PullRequest) string {
	if len(openPullRequests) == 0 {
		return "No open pull requests to work on.\n"
	}

	update := "This is work you have in progress or have yet to address:\n"

	for _, pull_request := range openPullRequests {
		update += fmt.Sprintf("🔨 Work on %s - %s\n", pull_request.Title, pull_request.URL)
	}

	return update
}
