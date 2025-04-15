package main

import (
	"fmt"

	github "github.com/mjimenez98/gh-stand-up/internal"
)

func main() {
	client := github.NewClient()
	user := client.GetUser()

	// Generate report
	fmt.Printf("Hi %s, your stand-up update is coming soon\n", user.Login)
}

// For more examples of using go-gh, see:
// https://github.com/cli/go-gh/blob/trunk/example_gh_test.go
