package instagram

import (
	"context"
	"math"

	"github.com/qcserestipy/instagram-api-go-client/pkg/account"
	"github.com/qcserestipy/instagram-api-go-client/pkg/media"
	mediaInsightsModel "github.com/qcserestipy/instagram-api-go-client/pkg/sdk/v24.0/media/client/insights"
	mediaApiModel "github.com/qcserestipy/instagram-api-go-client/pkg/sdk/v24.0/media/client/media"
	"github.com/qcserestipy/instagram-api-go-client/pkg/utils"
	"github.com/sirupsen/logrus"
)

// GetReels fetches all reels (VIDEO media) for a given Instagram account
// and returns them with their metrics (views, likes, comments, caption).
// since and until are optional Unix timestamp filters - pass nil to omit them.
func GetStories(ctx context.Context, accountsvc *account.Service, mediasvc *media.Service, accountID string) ([]Story, error) {
	var stories []Story

	// Get all media for the account
	accountMediaResponse, err := accountsvc.GetStoriesByUserID(ctx, accountID)
	if err != nil {
		return nil, err
	}

	// Collect media IDs
	storyIDs := []string{}
	for _, data := range accountMediaResponse.Payload.Data {
		storyIDs = append(storyIDs, data.ID)
	}

	// Fetch details for each media item
	for _, storyID := range storyIDs {
		fields := "media_product_type,timestamp"
		mediaApiResponse, err := mediasvc.GetMediaByID(
			ctx,
			&mediaApiModel.GetMediaByIDParams{
				InstagramMediaID: storyID,
				Fields:           &fields,
			},
		)
		if err != nil {
			logrus.Warnf("Could not fetch media %s: %v", storyID, err)
			continue
		}

		// Only process VIDEO media (STORY)
		if mediaApiResponse.Payload.MediaProductType != "STORY" {
			continue
		}

		// Fetch insights for view count using "views" metric
		var follows, navigation, profile_activity, profile_visits, reach, replies, shares, total_interactions, views int64
		mediaInsightsApiResponse, err := mediasvc.GetInsightsByMediaID(
			ctx,
			&mediaInsightsModel.GetInsightsByMediaIDParams{
				InstagramMediaID: storyID,
				Metric:           "follows,navigation,profile_activity,profile_visits,reach,replies,shares,total_interactions,views",
			},
		)
		if err != nil {
			logrus.Warnf("Could not fetch media insights %s: %v", storyID, err)
		} else if mediaInsightsApiResponse.Payload != nil &&
			mediaInsightsApiResponse.Payload.Data != nil &&
			len(mediaInsightsApiResponse.Payload.Data) > 0 &&
			len(mediaInsightsApiResponse.Payload.Data[0].Values) > 0 {
			for _, data := range mediaInsightsApiResponse.Payload.Data {
				if len(data.Values) > 0 {
					switch data.Name {
					case "follows":
						views = data.Values[0].Value
					case "navigation":
						navigation = data.Values[0].Value
					case "profile_activity":
						profile_activity = data.Values[0].Value
					case "profile_visits":
						profile_visits = data.Values[0].Value
					case "reach":
						reach = data.Values[0].Value
					case "replies":
						replies = data.Values[0].Value
					case "shares":
						shares = data.Values[0].Value
					case "total_interactions":
						total_interactions = data.Values[0].Value
					case "views":
						views = data.Values[0].Value

					}
				}
			}
		}

		engagementViews := 0.0
		if views > 0 {
			engagementViews = float64(total_interactions) / float64(views) * 100
		}

		parsedTime, err := utils.ParseTimestamp(mediaApiResponse.Payload.Timestamp)
		if err != nil {
			logrus.Warnf("Could not parse timestamp %s: %v", mediaApiResponse.Payload.Timestamp, err)
		}
		stories = append(stories, Story{
			ID:                mediaApiResponse.Payload.ID,
			Views:             views,
			Reach:             reach,
			Replies:           replies,
			Shares:            shares,
			Navigation:        navigation,
			ProfileActivity:   profile_activity,
			ProfileVisits:     profile_visits,
			Follows:           follows,
			TotalInteractions: total_interactions,
			DateTime:          parsedTime,
			EngagementViews:   math.Round(engagementViews*100) / 100,
		})
	}

	return stories, nil
}
