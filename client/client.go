package client

import (
	"context"

	v1 "github.com/alexfalkowski/bezeichner/api/bezeichner/v1"
	v1c "github.com/alexfalkowski/bezeichner/client/v1/config"
	"github.com/alexfalkowski/go-service/time"
)

// Client for bezeichner.
type Client struct {
	client v1.ServiceClient
	config *v1c.Config
}

// NewClient for bezeichner.
func NewClient(client v1.ServiceClient, config *v1c.Config) *Client {
	return &Client{client: client, config: config}
}

// GenerateIdentifiers for client.
func (c *Client) GenerateIdentifiers(ctx context.Context, app string, count uint64) ([]string, error) {
	ctx, cancel := context.WithTimeout(ctx, time.MustParseDuration(c.config.Timeout))
	defer cancel()

	req := &v1.GenerateIdentifiersRequest{Application: app, Count: count}

	resp, err := c.client.GenerateIdentifiers(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp.GetIds(), nil
}

// MapIdentifiers for client.
func (c *Client) MapIdentifiers(ctx context.Context, ids []string) ([]string, error) {
	ctx, cancel := context.WithTimeout(ctx, time.MustParseDuration(c.config.Timeout))
	defer cancel()

	req := &v1.MapIdentifiersRequest{Ids: ids}

	resp, err := c.client.MapIdentifiers(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp.GetIds(), nil
}
