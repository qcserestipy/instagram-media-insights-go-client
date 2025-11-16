package media

import (
	"context"
	"sync"

	"github.com/qcserestipy/instagram-media-insights-go-client/pkg/config"
	apiclient "github.com/qcserestipy/instagram-media-insights-go-client/pkg/sdk/v24.0/client"
)

var (
	ClientInstance *apiclient.InstagramMediaInsightsAPI
	ClientOnce     sync.Once
	ClientErr      error
)

func GetClient() (*apiclient.InstagramMediaInsightsAPI, error) {
	ClientOnce.Do(func() {
		apiURL, authInfo, err := config.CreateClientConfig()
		if err != nil {
			ClientErr = err
			return
		}

		// Create the client configuration
		cfg := apiclient.Config{
			URL:      apiURL,
			AuthInfo: authInfo,
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
