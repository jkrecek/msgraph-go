package graph

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"golang.org/x/oauth2"
	"io"
	"net/http"
)

const (
	HOST_URL        = "https://graph.microsoft.com"
	DEFAULT_VERSION = "1.0"
)

type Client struct {
	native  *http.Client
	version string
}

func NewClient(oauthConfig *oauth2.Config, oauthToken *oauth2.Token) *Client {
	ctx := context.Background()

	return &Client{
		native:  oauthConfig.Client(ctx, oauthToken),
		version: DEFAULT_VERSION,
	}
}

func (c *Client) SetVersion(version string) {
	c.version = version
}

func (c *Client) getUrl(path string) string {
	return fmt.Sprintf("%s/v%s/%s", HOST_URL, c.version, path)
}

func (c *Client) getRequest(path string, v interface{}) error {
	return c.doRequest("GET", path, nil, v)
}

func (c *Client) doRequest(method string, path string, body io.Reader, v interface{}) error {
	url := c.getUrl(path)
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := c.native.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode >= 400 {
		return graphError(url, resp.Body)
	} else {
		if v != nil {
			return parseGraphResult(resp.Body, v)
		} else {
			return nil
		}

	}
}
