package main

import (
	"fmt"
	"log"

	"github.com/mjimenez98/gh-stand-up/internal/github"
	"github.com/mjimenez98/gh-stand-up/internal/report"
)

func main() {
	// Create a new GitHub client.
	client, err := github.NewClient()
	if err != nil {
		log.Fatalf("Error creating GitHub client: %v", err)
	}

	// Get information to generate yesterday's report.
	user, err := client.GetUser()
	if err != nil {
		log.Fatalf("Error processing user information for yesterday's report: %v", err)
	}
	openedIssues, err := client.GetOpenedIssues(user.Login)
	if err != nil {
		log.Fatalf("Error processing issues opened for yesterday's report: %v", err)
	}
	openedPullRequests, err := client.GetOpenedPullRequests(user.Login)
	if err != nil {
		log.Fatalf("Error processing pull requests opened for yesterday's report: %v", err)
	}
	pullRequestsReviewed, err := client.GetPullRequestsReviewed(user.Login)
	if err != nil {
		log.Fatalf("Error processing pull request reviews for yesterday's report: %v", err)
	}

	// Generate yesterday's report.
	update := report.GenerateYesterdayReport(openedIssues, openedPullRequests, pullRequestsReviewed)
	fmt.Printf("%s\n", update)

	// Get information to generate today's report.
	openPullRequests, err := client.GetOpenPullRequests(user.Login)
	if err != nil {
		log.Fatalf("Error processing open issues for today's report: %v", err)
	}

	// Generate today's report.
	update = report.GenerateTodayReport(openPullRequests)
	fmt.Print(update)
}
