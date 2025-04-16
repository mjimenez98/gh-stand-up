package github

import (
	"fmt"
)

type User struct {
	Login string `json:"login"`
}

// GetUser retrieves the user information from the GitHub API.
func (c *Client) GetUser() User {
	var response User

	err := c.Client.Get("user", &response)
	if err != nil {
		fmt.Println(err)
		return User{}
	}

	return response
}
