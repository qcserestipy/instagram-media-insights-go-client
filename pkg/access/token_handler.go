package access

import (
	"context"

	"github.com/qcserestipy/instagram-api-go-client/pkg/sdk/v24.0/page/client/access_token"
)

func (s *Service) GetPageAccessToken(ctx context.Context, params *access_token.GetPageAccessTokenParams) (*access_token.GetPageAccessTokenOK, error) {
	response, err := s.client.Page.AccessToken.GetPageAccessToken(ctx, params)
	if err != nil {
		return nil, err
	}
	return response, nil
}
