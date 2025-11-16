package media

import (
	"github.com/qcserestipy/instagram-api-go-client/pkg/client"
	"github.com/qcserestipy/instagram-api-go-client/pkg/sdk/v24.0/media/client/media"
)

func GetMediaByID(params *media.GetMediaByIDParams) (*media.GetMediaByIDOK, error) {
	ctx, instagramClient, err := client.ContextWithClient()
	if err != nil {
		return nil, err
	}
	response, err := instagramClient.Media.Media.GetMediaByID(
		ctx,
		params,
	)
	if err != nil {
		return nil, err
	}
	return response, nil
}
