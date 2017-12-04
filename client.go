package graph

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"

	"golang.org/x/oauth2"
)

const (
	HOST_URL        = "https://graph.microsoft.com"
	DEFAULT_VERSION = "1.0"
)

type Client struct {
	native  *http.Client
	version string
	headers map[string]string
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

func (c *Client) SetHeaders(headers map[string]string) {
	c.headers = headers
}

func (c *Client) getAbsoluteUrl(path string) string {
	return fmt.Sprintf("%s/v%s/%s", HOST_URL, c.version, path)
}

func (c *Client) getRequest(path string, v interface{}) error {
	return c.doRequest("GET", path, nil, v)
}

func (c *Client) doRequest(method string, path string, body io.Reader, v interface{}) error {
	pathUrl, err := url.Parse(path)
	if err != nil {
		return err
	}

	var link string
	if !pathUrl.IsAbs() {
		link = c.getAbsoluteUrl(path)
	} else {
		link = path
	}

	req, err := http.NewRequest(method, link, body)
	for header, header_value := range c.headers {
		req.Header.Set(header, header_value)
	}

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

	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return graphError(link, resp.Body)
	} else {
		if v != nil {
			return parseGraphResult(resp.Body, v)
		} else {
			return nil
		}

	}
}

func (c *Client) readGetIntoFunc(path string, singularValue interface{}, newItemFn func(interface{})) error {
	wrp := new(ValueWrapper)

	sr := reflect.ValueOf(singularValue).Type()
	slcr := reflect.SliceOf(sr)
	valuesReflect := reflect.New(slcr)
	vals := valuesReflect.Interface()
	wrp.Value = &vals

	err := c.getRequest(path, wrp)
	if err != nil {
		return err
	}

	refSlc := valuesReflect.Elem()
	for i := 0; i < refSlc.Len(); i++ {
		newItemFn(refSlc.Index(i).Interface())
	}

	if wrp.NextLink == "" {
		return nil
	} else {
		return c.readGetIntoFunc(wrp.NextLink, singularValue, newItemFn)
	}
}
