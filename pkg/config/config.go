package config

import (
	"fmt"
	"net/url"
	"os"
	"sync"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
)

var (
	accessToken     string
	accessTokenOnce sync.Once
	accessTokenErr  error
)

// GetAccessToken retrieves the access token from environment variable
func GetAccessToken() (string, error) {
	accessTokenOnce.Do(func() {
		accessToken = os.Getenv("ACCESS_TOKEN")
		if accessToken == "" {
			accessTokenErr = fmt.Errorf("ACCESS_TOKEN environment variable is not set")
		}
	})
	return accessToken, accessTokenErr
}

// CreateClientConfig creates a common configuration for Instagram API clients
func CreateClientConfig() (*url.URL, runtime.ClientAuthInfoWriter, error) {
	token, err := GetAccessToken()
	if err != nil {
		return nil, nil, err
	}

	apiURL := &url.URL{
		Scheme: "https",
		Host:   "graph.facebook.com",
		Path:   "/v24.0",
	}

	authInfo := httptransport.APIKeyAuth("access_token", "query", token)

	return apiURL, authInfo, nil
}
