package infrastructure

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"flamingo.me/flamingo/v3/core/cache"
	"flamingo.me/flamingo/v3/framework/flamingo"
	"flamingo.me/flamingo/v3/framework/opencensus"
	"go.opencensus.io/stats"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/tag"
	"go.opencensus.io/trace"
)

type (
	// APIClient for openweather
	APIClient struct {
		baseURL    string
		apiKey     string
		httpClient *http.Client
		logger     flamingo.Logger
		cache      *cache.HTTPFrontend
	}
)

var (
	requestCount    = stats.Int64("openweather/api/calls", "number of calls to the API", stats.UnitDimensionless)
	requestDuration = stats.Float64(
		"openweather/api/request_duration_seconds",
		"The time elapsing between the request and the response to openweather",
		stats.UnitSeconds,
	)

	statusCodeKey, _ = tag.NewKey("statusCode")
)

// register view in opencensus
func init() {
	err := opencensus.View("openweather/api/calls", requestCount, view.Count())
	if err != nil {
		panic(err)
	}
	_ = opencensus.View("openweather/api/request_duration_seconds", requestDuration, view.Distribution(
		0, 0.2, 0.5, 1.0, 2.0, 5.0, 7.5, 10.0,
	), statusCodeKey)
}

// Inject dependencies
func (c *APIClient) Inject(
	httpClient *http.Client,
	logger flamingo.Logger,
	cfg *struct {
		BaseURL string `inject:"config:openweather.apiURL"`
		APIKey  string `inject:"config:openweather.apiKey"`
	},
	ann *struct {
		Cache *cache.HTTPFrontend `inject:"openweather"`
	},
) *APIClient {
	c.httpClient = httpClient
	c.logger = logger
	if cfg != nil {
		c.baseURL = cfg.BaseURL
		c.apiKey = cfg.APIKey
	}
	if ann != nil {
		c.cache = ann.Cache
	}

	return c
}

func (c *APIClient) request(ctx context.Context, method, path string, body io.Reader) (*http.Response, error) {
	path = strings.TrimRight(c.baseURL, "/") + "/" + strings.TrimLeft(path, "/")

	loadData := func(ctx context.Context) (*http.Response, *cache.Meta, error) {
		ctx, span := trace.StartSpan(ctx, "openweather/apiclient/request")
		defer span.End()

		c.logger.WithContext(ctx).Debugf("openweather.apiClient.request: method: %q path: %q", method, path)

		u, err := url.Parse(path)
		if err != nil {
			return nil, nil, err
		}

		query := u.Query()
		query.Add("appid", c.apiKey)
		query.Add("units", "metric")

		u.RawQuery = query.Encode()

		c.logger.Info("Requesting", u.String())

		request, err := http.NewRequest(method, u.String(), body)
		if err != nil {
			return nil, nil, err
		}
		request.Header.Set("Content-Type", "application/json")

		start := time.Now()
		response, err := c.httpClient.Do(request.WithContext(ctx))

		if err != nil {
			ctx, _ = tag.New(ctx, tag.Upsert(statusCodeKey, fmt.Sprintf("%vxx", response.StatusCode/100)))
		}

		stats.Record(ctx, requestCount.M(1))
		stats.Record(ctx, requestDuration.M(float64(time.Since(start).Nanoseconds())/1000000000))

		return response, &cache.Meta{
			Lifetime:  time.Minute,
			Gracetime: 2 * time.Minute,
		}, err
	}
	response, err := c.cache.Get(ctx, path, loadData)

	return response, err
}
