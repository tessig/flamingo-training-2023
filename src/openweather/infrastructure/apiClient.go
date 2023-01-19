package infrastructure

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"

	"flamingo.me/flamingo/v3/framework/flamingo"
	"flamingo.me/flamingo/v3/framework/opencensus"
	"go.opencensus.io/stats"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/trace"
)

type (
	// APIClient for openweather
	APIClient struct {
		baseURL    string
		apiKey     string
		httpClient *http.Client
		logger     flamingo.Logger
	}
)

var (
	errorCounter = stats.Int64("openweather/api/error_count", "number of errors on the API", stats.UnitDimensionless)
)

func init() {
	err := opencensus.View("openweather/api/error_count", errorCounter, view.Count())
	if err != nil {
		panic(err)
	}
}

// Inject dependencies
func (c *APIClient) Inject(
	httpClient *http.Client,
	logger flamingo.Logger,
	cfg *struct {
		BaseURL string `inject:"config:openweather.apiURL"`
		APIKey  string `inject:"config:openweather.apiKey"`
	},
) *APIClient {
	c.httpClient = httpClient
	c.logger = logger
	if cfg != nil {
		c.baseURL = cfg.BaseURL
		c.apiKey = cfg.APIKey
	}

	return c
}

func (c *APIClient) request(ctx context.Context, method, path string, body io.Reader) (*http.Response, error) {
	ctx, span := trace.StartSpan(ctx, "openweather/apiclient/request")
	defer span.End()

	path = strings.TrimRight(c.baseURL, "/") + "/" + strings.TrimLeft(path, "/")

	u, err := url.Parse(path)
	if err != nil {
		stats.Record(ctx, errorCounter.M(1))
		return nil, err
	}

	query := u.Query()
	query.Add("appid", c.apiKey)
	query.Add("units", "metric")

	u.RawQuery = query.Encode()

	request, err := http.NewRequest(method, u.String(), body)
	if err != nil {
		stats.Record(ctx, errorCounter.M(1))
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json")

	response, err := c.httpClient.Do(request.WithContext(ctx))
	if err != nil {
		stats.Record(ctx, errorCounter.M(1))
		span.SetStatus(trace.Status{
			Code:    trace.StatusCodeInternal,
			Message: err.Error(),
		})
		c.logger.WithContext(ctx).
			WithField(flamingo.LogKeyApicall, "1").
			WithField(flamingo.LogKeyPath, u.Path).
			Error(err)
	}

	return response, err
}
