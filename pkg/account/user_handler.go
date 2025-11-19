package account

import (
	"context"

	"github.com/qcserestipy/instagram-api-go-client/pkg/sdk/v24.0/account/client/user"
)

func (s *Service) GetUserByID(ctx context.Context, params *user.GetInstagramUserByIDParams) (*user.GetInstagramUserByIDOK, error) {
	response, err := s.client.Account.User.GetInstagramUserByID(
		ctx,
		params,
	)
	if err != nil {
		return nil, err
	}
	return response, nil
}
