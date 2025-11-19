package instagram

import (
	"math"

	"github.com/qcserestipy/instagram-api-go-client/pkg/account"
	"github.com/qcserestipy/instagram-api-go-client/pkg/media"
	accountMediaApiModel "github.com/qcserestipy/instagram-api-go-client/pkg/sdk/v24.0/account/client/media"
	mediaInsightsModel "github.com/qcserestipy/instagram-api-go-client/pkg/sdk/v24.0/media/client/insights"
	mediaApiModel "github.com/qcserestipy/instagram-api-go-client/pkg/sdk/v24.0/media/client/media"
	"github.com/qcserestipy/instagram-api-go-client/pkg/utils"
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
		fields := "media_product_type,like_count,comments_count,caption,timestamp"
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
		if mediaApiResponse.Payload.MediaProductType != "REELS" {
			continue
		}

		// Fetch insights for view count using "views" metric
		views := int64(0)
		reach := int64(0)
		shares := int64(0)
		saved := int64(0)
		totalInteractions := int64(0)
		mediaInsightsApiResponse, err := media.GetInsightsByMediaID(
			&mediaInsightsModel.GetInsightsByMediaIDParams{
				InstagramMediaID: reelID,
				Metric:           "views,reach,shares,saved,total_interactions",
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
						views = data.Values[0].Value
					case "reach":
						reach = data.Values[0].Value
					case "shares":
						shares = data.Values[0].Value
					case "saved":
						saved = data.Values[0].Value
					case "total_interactions":
						totalInteractions = data.Values[0].Value
					}
				}
			}
		}

		engagementViews := 0.0
		if views > 0 {
			engagementViews = float64(totalInteractions) / float64(views) * 100
		}

		parsedTime, err := utils.ParseTimestamp(mediaApiResponse.Payload.Timestamp)
		if err != nil {
			logrus.Warnf("Could not parse timestamp for media %s: %v", reelID, err)
			continue
		}
		reels = append(reels, Reel{
			ID:                mediaApiResponse.Payload.ID,
			Views:             views,
			Reach:             reach,
			Shares:            shares,
			Saves:             saved,
			Likes:             mediaApiResponse.Payload.LikeCount,
			Comments:          mediaApiResponse.Payload.CommentsCount,
			Caption:           mediaApiResponse.Payload.Caption,
			DateTime:          parsedTime,
			TotalInteractions: totalInteractions,
			EngagementViews:   math.Round(engagementViews*100) / 100,
		})
	}

	return reels, nil
}
