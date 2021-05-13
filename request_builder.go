package ddseries

import (
	"context"
	"fmt"
	"net/http"
	"strings"
)

// RequestBuilder holds data for request
type RequestBuilder struct {
	Host      string
	Subdomain string
	AppKey    string
	ApiKey    string
}

// Call builds request from the input
func (b *RequestBuilder) Call(ctx context.Context, queries ...Query) (*http.Request, error) {
	body := "resp_version=2"

	for index, q := range queries {
		body += "&"
		body += q.String(index)
	}
	url := fmt.Sprintf("https://%s.%s/series/batch_query", b.Subdomain, b.Host)

	r, err := http.NewRequestWithContext(ctx, http.MethodPost, url, strings.NewReader(body))
	if err != nil {
		return nil, err
	}

	r.Header.Set("DD-APPLICATION-KEY", b.AppKey)
	r.Header.Set("DD-API-KEY", b.ApiKey)
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Set("Accept-Encoding", "gzip, deflate, br")

	return r, nil
}
