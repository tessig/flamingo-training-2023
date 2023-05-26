package infrastructure

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type (
	// APIClient for openweather
	APIClient struct {
		baseURL    string
		apiKey     string
		httpClient *http.Client
	}
)

// Inject dependencies
func (c *APIClient) Inject(
	httpClient *http.Client,
	cfg *struct {
		BaseURL string `inject:"config:openweather.apiURL"`
		APIKey  string `inject:"config:openweather.apiKey"`
	},
) *APIClient {
	c.httpClient = httpClient
	if cfg != nil {
		c.baseURL = cfg.BaseURL
		c.apiKey = cfg.APIKey
	}

	return c
}

func (c *APIClient) request(ctx context.Context, method, path string, body io.Reader) (*http.Response, error) {
	path = strings.TrimRight(c.baseURL, "/") + "/" + strings.TrimLeft(path, "/")

	u, err := url.Parse(path)
	if err != nil {
		return nil, err
	}

	query := u.Query()
	query.Add("appid", c.apiKey)
	query.Add("units", "metric")

	u.RawQuery = query.Encode()

	request, err := http.NewRequest(method, u.String(), body)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json")

	response, err := c.httpClient.Do(request.WithContext(ctx))

	return response, err
}
