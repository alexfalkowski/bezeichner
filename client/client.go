package client

import (
	"context"

	v1 "github.com/alexfalkowski/bezeichner/api/bezeichner/v1"
)

// Client for bezeichner.
type Client struct {
	client v1.ServiceClient
}

// NewClient for bezeichner.
func NewClient(client v1.ServiceClient) *Client {
	return &Client{client: client}
}

// GenerateIdentifiers for client.
func (c *Client) GenerateIdentifiers(ctx context.Context, app string, count uint64) ([]string, error) {
	req := &v1.GenerateIdentifiersRequest{Application: app, Count: count}

	resp, err := c.client.GenerateIdentifiers(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp.GetIds(), nil
}

// MapIdentifiers for client.
func (c *Client) MapIdentifiers(ctx context.Context, ids []string) ([]string, error) {
	req := &v1.MapIdentifiersRequest{Ids: ids}

	resp, err := c.client.MapIdentifiers(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp.GetIds(), nil
}
