package account

import (
	"github.com/qcserestipy/instagram-api-go-client/pkg/client"
	"github.com/qcserestipy/instagram-api-go-client/pkg/sdk/v24.0/account/client/user"
)

func GetUserByID(params *user.GetInstagramUserByIDParams) (*user.GetInstagramUserByIDOK, error) {
	ctx, instagramClient, err := client.ContextWithClient()
	if err != nil {
		return nil, err
	}
	response, err := instagramClient.Account.User.GetInstagramUserByID(
		ctx,
		params,
	)
	if err != nil {
		return nil, err
	}
	return response, nil
}
