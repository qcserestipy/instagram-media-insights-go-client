package client

import (
	"context"
	"sync"

	"github.com/qcserestipy/instagram-media-insights-go-client/pkg/config"
	accountclient "github.com/qcserestipy/instagram-media-insights-go-client/pkg/sdk-account/v24.0/client"
	mediaclient "github.com/qcserestipy/instagram-media-insights-go-client/pkg/sdk/v24.0/client"
)

// InstagramClient holds both media and account API clients
type InstagramClient struct {
	Media   *mediaclient.InstagramMediaInsightsAPI
	Account *accountclient.InstagramAccountInsightsAPI
}

var (
	instance *InstagramClient
	once     sync.Once
	initErr  error
)

// Get returns a singleton instance of the unified Instagram client with both APIs
func Get() (*InstagramClient, error) {
	once.Do(func() {
		apiURL, authInfo, err := config.CreateClientConfig()
		if err != nil {
			initErr = err
			return
		}

		// Create media client
		mediaCfg := mediaclient.Config{
			URL:      apiURL,
			AuthInfo: authInfo,
		}

		// Create account client
		accountCfg := accountclient.Config{
			URL:      apiURL,
			AuthInfo: authInfo,
		}

		instance = &InstagramClient{
			Media:   mediaclient.New(mediaCfg),
			Account: accountclient.New(accountCfg),
		}
	})
	return instance, initErr
}

// ContextWithClient returns a context and the unified Instagram client
func ContextWithClient() (context.Context, *InstagramClient, error) {
	client, err := Get()
	if err != nil {
		return nil, nil, err
	}
	ctx := context.Background()
	return ctx, client, nil
}
