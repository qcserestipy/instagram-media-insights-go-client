package account

import (
	"context"

	"github.com/qcserestipy/instagram-api-go-client/pkg/sdk/v24.0/account/client/media"
)

func (s *Service) GetMediaByUserID(ctx context.Context, params *media.GetMediaByUserIDParams) (*media.GetMediaByUserIDOK, error) {
	response, err := s.client.Account.Media.GetMediaByUserID(
		ctx,
		params,
	)
	if err != nil {
		return nil, err
	}
	return response, nil
}
