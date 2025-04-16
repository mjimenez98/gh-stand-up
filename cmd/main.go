package main

import (
	"fmt"

	github "github.com/mjimenez98/gh-stand-up/internal"
)

func main() {
	client := github.NewClient()
	if client == nil {
		fmt.Println("Failed to create GitHub client")
		return
	}

	user := client.GetUser()
	openedIssues := client.GetOpenedIssues(user.Login)

	fmt.Printf("Hi %s 🙌\n\n", user.Login)
	fmt.Println("This is what you did yesterday:")
	for _, issue := range openedIssues {
		fmt.Printf("- You opened the issue: %s\n", issue.URL)
	}
}
