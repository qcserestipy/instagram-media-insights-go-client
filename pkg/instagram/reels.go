package instagram

import (
	"math"

	"github.com/qcserestipy/instagram-api-go-client/pkg/account"
	"github.com/qcserestipy/instagram-api-go-client/pkg/media"
	accountMediaApiModel "github.com/qcserestipy/instagram-api-go-client/pkg/sdk/v24.0/account/client/media"
	mediaInsightsModel "github.com/qcserestipy/instagram-api-go-client/pkg/sdk/v24.0/media/client/insights"
	mediaApiModel "github.com/qcserestipy/instagram-api-go-client/pkg/sdk/v24.0/media/client/media"
	"github.com/sirupsen/logrus"
)

// GetReels fetches all reels (VIDEO media) for a given Instagram account
// and returns them with their metrics (views, likes, comments, caption).
// since and until are optional Unix timestamp filters - pass nil to omit them.
func GetReels(accountID string, since *int64, until *int64) ([]Reel, error) {
	var reels []Reel

	// Get all media for the account
	accountMediaResponse, err := account.GetMediaByUserID(&accountMediaApiModel.GetMediaByUserIDParams{
		InstagramAccountID: accountID,
		Since:              since,
		Until:              until,
	})
	if err != nil {
		return nil, err
	}

	// Collect media IDs
	reelIDs := []string{}
	for _, data := range accountMediaResponse.Payload.Data {
		reelIDs = append(reelIDs, data.ID)
	}

	// Fetch details for each media item
	for _, reelID := range reelIDs {
		fields := "media_type,like_count,comments_count,caption,timestamp"
		mediaApiResponse, err := media.GetMediaByID(
			&mediaApiModel.GetMediaByIDParams{
				InstagramMediaID: reelID,
				Fields:           &fields,
			},
		)
		if err != nil {
			logrus.Warnf("Could not fetch media %s: %v", reelID, err)
			continue
		}

		// Only process VIDEO media (reels)
		if mediaApiResponse.Payload.MediaType != "VIDEO" {
			continue
		}

		// Fetch insights for view count using "views" metric
		views := 0
		reach := 0
		shares := 0
		saved := 0
		mediaInsightsApiResponse, err := media.GetInsightsByMediaID(
			&mediaInsightsModel.GetInsightsByMediaIDParams{
				InstagramMediaID: reelID,
				Metric:           "views,reach,shares,saved",
			},
		)
		if err != nil {
			logrus.Warnf("Could not fetch media insights %s: %v", reelID, err)
		} else if mediaInsightsApiResponse.Payload != nil &&
			mediaInsightsApiResponse.Payload.Data != nil &&
			len(mediaInsightsApiResponse.Payload.Data) > 0 &&
			len(mediaInsightsApiResponse.Payload.Data[0].Values) > 0 {
			for _, data := range mediaInsightsApiResponse.Payload.Data {
				if len(data.Values) > 0 {
					switch data.Name {
					case "views":
						views = int(data.Values[0].Value)
					case "reach":
						reach = int(data.Values[0].Value)
					case "shares":
						shares = int(data.Values[0].Value)
					case "saved":
						saved = int(data.Values[0].Value)
					}
				}
			}
		}

		engagementViews := 0.0
		if views > 0 {
			engagementViews = float64(int(mediaApiResponse.Payload.LikeCount)+int(mediaApiResponse.Payload.CommentsCount)+shares+saved) / float64(views) * 100
		}

		reels = append(reels, Reel{
			ID:              mediaApiResponse.Payload.ID,
			Views:           views,
			Reach:           reach,
			Shares:          shares,
			Saves:           saved,
			Likes:           int(mediaApiResponse.Payload.LikeCount),
			Comments:        int(mediaApiResponse.Payload.CommentsCount),
			Caption:         mediaApiResponse.Payload.Caption,
			DateTime:        mediaApiResponse.Payload.Timestamp,
			EngagementViews: math.Round(engagementViews*100) / 100,
		})
	}

	return reels, nil
}
