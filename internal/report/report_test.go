package report

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mjimenez98/gh-stand-up/internal/github"
)

func TestDoesNotGenerateReportWhenNoActivity(t *testing.T) {
	generatedReport := GenerateYesterdayReport([]github.Issue{}, []github.PullRequest{}, []github.PullRequest{})

	assert.Equal(t, "No activity detected yesterday.\n", generatedReport, "Generated report does not match expected output when no activity")
}

func TestGenerateYesterdayReportWithOpenIssues(t *testing.T) {
	openedIssues := []github.Issue{
		{Title: "Open issue", URL: "https://api.github.com/example/repo/issues/1"},
		{Title: "Another issue", URL: "https://api.github.com/example/repo/issues/2"},
	}

	generatedReport := GenerateYesterdayReport(openedIssues, []github.PullRequest{}, []github.PullRequest{})

	assert.Contains(t, generatedReport, "This is what you did yesterday:\n", "Generated report does not contain the expected header")
	assert.Contains(t, generatedReport, "📋 Opened Open issue - https://api.github.com/example/repo/issues/1", "Generated report does not contain the expected opened issue")
	assert.Contains(t, generatedReport, "📋 Opened Another issue - https://api.github.com/example/repo/issues/2", "Generated report does not contain the expected opened issue")
}

func TestGenerateYesterdayReportWithOpenPullRequests(t *testing.T) {
	openedPullRequests := []github.PullRequest{
		{Title: "Open PR", URL: "https://api.github.com/example/repo/pull/1"},
		{Title: "Another PR", URL: "https://api.github.com/example/repo/pull/2"},
	}

	generatedReport := GenerateYesterdayReport([]github.Issue{}, openedPullRequests, []github.PullRequest{})

	assert.Contains(t, generatedReport, "This is what you did yesterday:\n", "Generated report does not contain the expected header")
	assert.Contains(t, generatedReport, "🔀 Opened Open PR - https://api.github.com/example/repo/pull/1", "Generated report does not contain the expected opened pull request")
	assert.Contains(t, generatedReport, "🔀 Opened Another PR - https://api.github.com/example/repo/pull/2", "Generated report does not contain the expected opened pull request")
}

func TestGenerateYesterdayReportWithReviewedPullRequests(t *testing.T) {
	pullRequestsReviewed := []github.PullRequest{
		{Title: "Reviewed PR", URL: "https://api.github.com/example/repo/pull/3"},
	}

	generatedReport := GenerateYesterdayReport([]github.Issue{}, []github.PullRequest{}, pullRequestsReviewed)

	assert.Contains(t, generatedReport, "This is what you did yesterday:\n", "Generated report does not contain the expected header")
	assert.Contains(t, generatedReport, "🔍 Reviewed Reviewed PR - https://api.github.com/example/repo/pull/3", "Generated report does not contain the expected reviewed pull request")
}
