package media

import (
	"context"
	"fmt"
	"net/url"
	"os"
	"sync"

	httptransport "github.com/go-openapi/runtime/client"
	apiclient "github.com/qcserestipy/instagram-media-insights-go-client/pkg/sdk/v24.0/client"
)

var (
	ClientInstance *apiclient.InstagramMediaInsightsAPI
	ClientOnce     sync.Once
	ClientErr      error
)

func GetClient() (*apiclient.InstagramMediaInsightsAPI, error) {
	ClientOnce.Do(func() {
		// read access token from environment variable or configuration
		accessToken := os.Getenv("ACCESS_TOKEN")
		if accessToken == "" {
			ClientErr = fmt.Errorf("ACCESS_TOKEN environment variable is not set")
			return
		}
		// Create the client configuration
		cfg := apiclient.Config{
			URL: &url.URL{
				Scheme: "https",
				Host:   "graph.facebook.com",
				Path:   "/v24.0",
			},
			// Pass the access token as a query parameter
			AuthInfo: httptransport.APIKeyAuth("access_token", "query", accessToken),
		}
		// Create the client
		ClientInstance = apiclient.New(cfg)
	})
	return ClientInstance, ClientErr
}

func ContextWithClient() (context.Context, *apiclient.InstagramMediaInsightsAPI, error) {
	client, err := GetClient()
	if err != nil {
		return nil, nil, err
	}
	ctx := context.Background()
	return ctx, client, nil
}
