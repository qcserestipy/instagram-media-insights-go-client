package account

import (
	"context"

	"github.com/qcserestipy/instagram-api-go-client/pkg/sdk/v24.0/account/client/stories"
)

func (s *Service) GetStoriesByUserID(ctx context.Context, accountId string) (*stories.GetStoriesByUserIDOK, error) {
	response, err := s.client.Account.Stories.GetStoriesByUserID(ctx, &stories.GetStoriesByUserIDParams{
		InstagramAccountID: accountId,
	})
	if err != nil {
		return nil, err
	}
	return response, nil
}
