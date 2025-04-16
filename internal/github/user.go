package github

import (
	"fmt"
)

type User struct {
	Login string `json:"login"`
}

// GetUser retrieves the user information from the GitHub API.
func (c *Client) GetUser() (User, error) {
	var response User

	err := c.Client.Get("user", &response)
	if err != nil {
		return User{}, fmt.Errorf("error retrieving user information: %w", err)
	}

	return response, nil
}
