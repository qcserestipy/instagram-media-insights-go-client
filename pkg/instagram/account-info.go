package instagram

import (
	"fmt"

	"github.com/qcserestipy/instagram-api-go-client/pkg/account"
	"github.com/qcserestipy/instagram-api-go-client/pkg/sdk/v24.0/account/client/user"
	"github.com/qcserestipy/instagram-api-go-client/pkg/utils"
)

func GetAccountInfo(accountID string) (*AccountInfo, error) {
	fields := "id,username,followers_count,follows_count,media_count,profile_picture_url,biography,website"
	insightsResponse, err := account.GetUserByID(&user.GetInstagramUserByIDParams{
		InstagramAccountID: accountID,
		Fields:             &fields,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to %v", utils.ParseAPIError(err, "get "+"account info"))
	}
	return &AccountInfo{
		ID:                insightsResponse.Payload.ID,
		Username:          insightsResponse.Payload.Username,
		FollowersCount:    insightsResponse.Payload.FollowersCount,
		FollowingCount:    insightsResponse.Payload.FollowsCount,
		MediaCount:        insightsResponse.Payload.MediaCount,
		ProfilePictureURL: insightsResponse.Payload.ProfilePictureURL,
		Bio:               insightsResponse.Payload.Biography,
		Website:           insightsResponse.Payload.Website,
	}, nil
}
