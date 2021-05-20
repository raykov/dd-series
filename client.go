package ddseries

import (
	"compress/gzip"
	"context"
	"errors"
	"io"
	"net/http"
	"os"
)

// NewClient returns a new client
func NewClient(httpClient *http.Client) *Client {
	return &Client{
		HttpClient: httpClient,

		Host:      os.Getenv("DATADOG_HOST"),
		Subdomain: "app",
		AppKey:    os.Getenv("DATADOG_APP_KEY"),
		ApiKey:    os.Getenv("DATADOG_API_KEY"),
	}
}

// Client holds info about client
type Client struct {
	HttpClient *http.Client

	Host      string
	Subdomain string
	AppKey    string
	ApiKey    string
}

// Do executes request
func (c *Client) Do(ctx context.Context, queries ...Query) ([]byte, error) {
	rb := RequestBuilder{
		Host:      c.Host,
		Subdomain: c.Subdomain,
		AppKey:    c.AppKey,
		ApiKey:    c.ApiKey,
	}

	r, err := rb.Call(ctx, queries...)
	if err != nil {
		return nil, err
	}

	resp, err := c.HttpClient.Do(r)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		errMessage, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		return nil, errors.New(string(errMessage))
	}

	gr, err := gzip.NewReader(resp.Body)
	if err != nil {
		return nil, err
	}
	defer gr.Close()

	body, err := io.ReadAll(gr)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// DoWithParsing executes request and parse it to Body struct
func (c *Client) DoWithParsing(ctx context.Context, queries ...Query) (*Body, error) {
	body, err := c.Do(ctx, queries...)
	if err != nil {
		return nil, err
	}

	b := &Body{}
	err = b.Unmarshall(body)
	if err != nil {
		return nil, err
	}

	return b, nil
}
