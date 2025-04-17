package report

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mjimenez98/gh-stand-up/internal/github"
)

func TestGenerateYesterdayReport(t *testing.T) {
	t.Run("When no activity", func(t *testing.T) {
		generatedReport := GenerateYesterdayReport([]github.Issue{}, []github.PullRequest{}, []github.PullRequest{})

		assert.Equal(t, "No activity detected yesterday.\n", generatedReport)
	})

	t.Run("When issues were opened", func(t *testing.T) {
		openedIssues := []github.Issue{
			{Title: "Open issue", URL: "https://api.github.com/example/repo/issues/1"},
			{Title: "Another issue", URL: "https://api.github.com/example/repo/issues/2"},
		}

		generatedReport := GenerateYesterdayReport(openedIssues, []github.PullRequest{}, []github.PullRequest{})

		assert.Contains(t, generatedReport, "This is what you did yesterday:\n")
		assert.Contains(t, generatedReport, "📋 Opened Open issue - https://api.github.com/example/repo/issues/1")
		assert.Contains(t, generatedReport, "📋 Opened Another issue - https://api.github.com/example/repo/issues/2")
	})

	t.Run("When pull requests were opened", func(t *testing.T) {
		openedPullRequests := []github.PullRequest{
			{Title: "Open PR", URL: "https://api.github.com/example/repo/pull/1"},
			{Title: "Another PR", URL: "https://api.github.com/example/repo/pull/2"},
		}

		generatedReport := GenerateYesterdayReport([]github.Issue{}, openedPullRequests, []github.PullRequest{})

		assert.Contains(t, generatedReport, "This is what you did yesterday:\n")
		assert.Contains(t, generatedReport, "🔀 Opened Open PR - https://api.github.com/example/repo/pull/1")
		assert.Contains(t, generatedReport, "🔀 Opened Another PR - https://api.github.com/example/repo/pull/2")
	})

	t.Run("When pull requests were reviewed", func(t *testing.T) {
		pullRequestsReviewed := []github.PullRequest{
			{Title: "Reviewed PR", URL: "https://api.github.com/example/repo/pull/3"},
		}

		generatedReport := GenerateYesterdayReport([]github.Issue{}, []github.PullRequest{}, pullRequestsReviewed)

		assert.Contains(t, generatedReport, "This is what you did yesterday:\n")
		assert.Contains(t, generatedReport, "🔍 Reviewed Reviewed PR - https://api.github.com/example/repo/pull/3")
	})
}

func TestGenerateTodayReport(t *testing.T) {
	t.Run("When no activity", func(t *testing.T) {
		generatedReport := GenerateTodayReport([]github.PullRequest{})

		assert.Equal(t, "No open pull requests to work on.\n", generatedReport)
	})

	t.Run("When pull requests are open", func(t *testing.T) {
		openPullRequests := []github.PullRequest{
			{Title: "Open PR", URL: "https://api.github.com/example/repo/pull/1"},
			{Title: "Another PR", URL: "https://api.github.com/example/repo/pull/2"},
		}

		generatedReport := GenerateTodayReport(openPullRequests)

		assert.Contains(t, generatedReport, "This is work you have in progress or have yet to address:\n")
		assert.Contains(t, generatedReport, "🔨 Work on Open PR - https://api.github.com/example/repo/pull/1")
		assert.Contains(t, generatedReport, "🔨 Work on Another PR - https://api.github.com/example/repo/pull/2")
	})
}
