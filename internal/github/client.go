// Package github provides a client for interacting with the GitHub API
// using the go-gh library.
package github

import (
	"fmt"

	"github.com/cli/go-gh/v2/pkg/api"
)

// Client is a wrapper around the GitHub API client.
type Client struct {
	Client *api.RESTClient
}

// NewClient creates a new GitHub API client.
func NewClient() (*Client, error) {
	client, err := api.DefaultRESTClient()
	if err != nil {
		return nil, fmt.Errorf("failed to create REST client: %w", err)
	}

	return &Client{
		Client: client,
	}, nil
}
