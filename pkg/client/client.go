package client

import (
	"net/url"

	"github.com/go-openapi/runtime"
	"github.com/qcserestipy/instagram-api-go-client/pkg/config"
	accountclient "github.com/qcserestipy/instagram-api-go-client/pkg/sdk/v24.0/account/client"
	mediaclient "github.com/qcserestipy/instagram-api-go-client/pkg/sdk/v24.0/media/client"
	pageclient "github.com/qcserestipy/instagram-api-go-client/pkg/sdk/v24.0/page/client"
)

// InstagramClient holds both media and account API clients
type InstagramClient struct {
	Media   *mediaclient.InstagramMediaInsightsAPI
	Account *accountclient.InstagramAccountInsightsAPI
	Page    *pageclient.FacebookPageAPI
}

// Again, adjust this to your real auth type if you want an explicit constructor variant.
type AuthInfoType = any

// NewFromConfig creates a new InstagramClient from a given apiURL + authInfo.
// This is convenient if another project wants to fully control configuration.
func NewFromConfig(apiURL *url.URL, authInfo runtime.ClientAuthInfoWriter) *InstagramClient {
	mediaCfg := mediaclient.Config{
		URL:      apiURL,
		AuthInfo: authInfo,
	}
	accountCfg := accountclient.Config{
		URL:      apiURL,
		AuthInfo: authInfo,
	}
	pageCfg := pageclient.Config{
		URL:      apiURL,
		AuthInfo: authInfo,
	}

	return &InstagramClient{
		Media:   mediaclient.New(mediaCfg),
		Account: accountclient.New(accountCfg),
		Page:    pageclient.New(pageCfg),
	}
}

// NewDefault creates a new InstagramClient using your existing config.CreateClientConfig().
func NewDefault() (*InstagramClient, error) {
	apiURL, authInfo, err := config.CreateClientConfig()
	if err != nil {
		return nil, err
	}
	return NewFromConfig(apiURL, authInfo), nil
}
