package media

import (
	"context"

	"github.com/qcserestipy/instagram-api-go-client/pkg/sdk/v24.0/media/client/media"
)

func (s *Service) GetMediaByID(ctx context.Context, params *media.GetMediaByIDParams) (*media.GetMediaByIDOK, error) {
	response, err := s.client.Media.Media.GetMediaByID(
		ctx,
		params,
	)
	if err != nil {
		return nil, err
	}
	return response, nil
}
