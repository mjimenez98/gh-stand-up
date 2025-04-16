package main

import (
	"fmt"

	github "github.com/mjimenez98/gh-stand-up/internal"
)

func main() {
	client := github.NewClient()
	user := client.GetUser()
	openedIssues := client.GetOpenedIssues(user.Login)

	fmt.Printf("Hi %s 🙌\n\n", user.Login)
	fmt.Println("This is what you did yesterday:")
	for _, issue := range openedIssues {
		fmt.Printf("- You opened the issue: %s\n", issue.URL)
	}
}
