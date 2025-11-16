package account

import (
	"github.com/qcserestipy/instagram-api-go-client/pkg/client"
	"github.com/qcserestipy/instagram-api-go-client/pkg/sdk/v24.0/account/client/stories"
)

func GetStoriesByUserID(accountId string) (*stories.GetStoriesByUserIDOK, error) {
	ctx, instagramClient, err := client.ContextWithClient()
	if err != nil {
		return nil, err
	}
	response, err := instagramClient.Account.Stories.GetStoriesByUserID(ctx, &stories.GetStoriesByUserIDParams{
		InstagramAccountID: accountId,
	})
	if err != nil {
		return nil, err
	}
	return response, nil
}
