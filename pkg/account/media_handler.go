package account

import (
	"github.com/qcserestipy/instagram-api-go-client/pkg/client"
	"github.com/qcserestipy/instagram-api-go-client/pkg/sdk/v24.0/account/client/media"
)

func GetMediaByUserID(params *media.GetMediaByUserIDParams) (*media.GetMediaByUserIDOK, error) {
	ctx, instagramClient, err := client.ContextWithClient()
	if err != nil {
		return nil, err
	}
	response, err := instagramClient.Account.Media.GetMediaByUserID(
		ctx,
		params,
	)
	if err != nil {
		return nil, err
	}
	return response, nil
}
