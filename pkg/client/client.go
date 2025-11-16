package client

import (
	"net/url"
	"os"
	"sync"

	"github.com/qcserstipy/instagram-media-insights-go-client/pkg/client"
	"github.com/sirupsen/logrus"
)

var (
	ClientInstance *InstagramMediaInsightsAPI
	ClientOnce     sync.Once
	ClientErr      error
)

func GetClient() *InstagramMediaInsightsAPI {
	ClientOnce.Do(func() {

		// read access token from environment variable or configuration
		accessToken := os.Getenv("ACCESS_TOKEN")
		if accessToken == "" {
			logrus.Fatal("ACCESS_TOKEN environment variable is not set")
		}

		// Create the client configuration
		cfg := client.Config{
			URL: &url.URL{
				Scheme: "https",
				Host:   "graph.facebook.com",
				Path:   "/v24.0",
			},
			// Pass the access token as a query parameter
			AuthInfo: httptransport.APIKeyAuth("access_token", "query", accessToken),
		}

		// Create the client
		ClientInstance := client.New(cfg)
		// config, err := GetCurrentHarborConfig()
		// if err != nil {
		// 	ClientErr = fmt.Errorf("failed to get current credential name: %v", err)
		// 	return
		// }
		// credentialName := config.CurrentCredentialName
		// if credentialName == "" {
		// 	ClientErr = fmt.Errorf("current-credential-name is not set in config file")
		// 	return
		// }

		// ClientInstance, ClientErr = GetClientByCredentialName(credentialName)
		// if ClientErr != nil {
		// 	log.Errorf("failed to initialize client: %v", ClientErr)
		// 	return
		// }
	})

	return ClientInstance
}
