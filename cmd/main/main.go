package main

import (
	"fmt"
	"reflect"

	"github.com/qcserestipy/instagram-api-go-client/pkg/account"
	"github.com/qcserestipy/instagram-api-go-client/pkg/media"
	accinsights "github.com/qcserestipy/instagram-api-go-client/pkg/sdk/v24.0/account/client/insights"
	accountMediaModel "github.com/qcserestipy/instagram-api-go-client/pkg/sdk/v24.0/account/client/media"
	insightModel "github.com/qcserestipy/instagram-api-go-client/pkg/sdk/v24.0/media/client/insights"
	"github.com/sirupsen/logrus"
)

func main() {
	params := insightModel.NewGetInsightsByMediaIDParams()
	params.InstagramMediaID = "18112405726596121"
	params.Metric = "reach,likes,comments"
	params.Period = "day"
	resp, err := media.GetInsightsByMediaID(params)
	if err != nil {
		logrus.Fatalf("fatal error: %v", err)
	}

	fmt.Printf("Success! Response: %+v\n", resp)
	if resp.Payload != nil && resp.Payload.Data != nil {
		for _, data := range resp.Payload.Data {
			fmt.Printf("Metric: %s, Title: %s\n", data.Name, data.Title)
			if data.Values != nil {
				for _, val := range data.Values {
					fmt.Printf("  Value: %d\n", val.Value)
				}
			}
		}
	}

	accountParams := accinsights.NewGetInsightsByAccountIDParams()
	accountParams.InstagramAccountID = "17841464714098258"
	accountParams.Metric = "engaged_audience_demographics"
	accountParams.Period = "lifetime"
	metricType := "total_value"
	accountParams.MetricType = &metricType
	timeframe := "this_month"
	accountParams.Timeframe = &timeframe
	breakdown := "country"
	accountParams.Breakdown = &breakdown
	accountResp, err := account.GetInsightsByAccountID(accountParams)
	if err != nil {
		logrus.Fatalf("fatal error: %v", err)
	}

	fmt.Printf("Success! Account Response: %+v\n", accountResp)
	if accountResp.Payload != nil && accountResp.Payload.Data != nil {
		for _, data := range accountResp.Payload.Data {
			fmt.Printf("Metric: %s, Title: %s\n", data.Name, data.Title)
			if data.Values != nil {
				for _, val := range data.Values {
					fmt.Printf("  Value: %d\n", val.Value)
				}
			}
			if data.TotalValue.Breakdowns != nil {
				for _, val := range data.TotalValue.Breakdowns {
					value := reflect.ValueOf(val.Results)
					for i := 0; i < value.Len(); i++ {
						item := value.Index(i).Interface()
						fmt.Printf("    Breakdown Item: %+v\n", item)
					}
				}
			}
		}
	}

	storiesResp, err := account.GetStoriesByUserID("17841464714098258")
	if err != nil {
		logrus.Fatalf("fatal error: %v", err)
	}

	fmt.Printf("Success! Stories Response: %+v\n", storiesResp)
	if storiesResp.Payload != nil && storiesResp.Payload.Data != nil {
		for _, story := range storiesResp.Payload.Data {
			fmt.Printf("Story ID: %s\n", story.ID)
			// Now you can use the client to make requests
			params := insightModel.NewGetInsightsByMediaIDParams()
			params.InstagramMediaID = story.ID
			params.Metric = "reach"
			// params.Period = "day"
			resp, err := media.GetInsightsByMediaID(params)
			if err != nil {
				logrus.Fatalf("fatal error: %v", err)
			}

			fmt.Printf("Success! Response: %+v\n", resp)
			if resp.Payload != nil && resp.Payload.Data != nil {
				for _, data := range resp.Payload.Data {
					fmt.Printf("Metric: %s, Title: %s\n", data.Name, data.Title)
					if data.Values != nil {
						for _, val := range data.Values {
							fmt.Printf("  Value: %d\n", val.Value)
						}
					}
				}
			}

			// Note: Instagram Stories typically don't support comments through the Graph API
			// Comments are only available for feed posts, reels, and IGTV videos
			commentResp, err := media.GetCommentsByMediaID(story.ID)
			if err != nil {
				logrus.Warnf("Could not fetch comments for story %s: %v (Stories typically don't have comments)", story.ID, err)
			} else {
				fmt.Printf("Success! Comments Response: %+v\n", commentResp)
				if commentResp.Payload != nil && commentResp.Payload.Data != nil {
					for _, comment := range commentResp.Payload.Data {
						fmt.Printf("Comment ID: %s\n", comment.ID)
					}
				}
			}
		}
	}

	mediaResp, err := account.GetMediaByUserID(&accountMediaModel.GetMediaByUserIDParams{
		InstagramAccountID: "17841475608626706",
	})
	if err != nil {
		logrus.Fatalf("fatal error: %v", err)
	}
	for _, mediaItem := range mediaResp.Payload.Data {
		fmt.Printf("Media ID: %s\n", mediaItem.ID)
		resp, err := media.GetInsightsByMediaID(&insightModel.GetInsightsByMediaIDParams{
			InstagramMediaID: mediaItem.ID,
			Metric:           "reach,comments",
		})
		if err != nil {
			logrus.Warnf("not supported: %v", err)
			continue
		}
		if resp != nil && resp.Payload != nil && resp.Payload.Data != nil {
			for _, data := range resp.Payload.Data {
				fmt.Printf("Metric: %s, Title: %s\n", data.Name, data.Title)
				if data.Values != nil {
					for _, val := range data.Values {
						fmt.Printf("  Value: %d\n", val.Value)
					}
				}
			}
		}

		postResp, err := media.CreateCommentOnMedia(mediaItem.ID, "This is a remark from the Instagram API Go Client!")
		if err != nil {
			logrus.Warnf("failed to create comment on media %s: %v", mediaItem.ID, err)
		} else {
			fmt.Printf("Success! Comment created on media %s: %+v\n", mediaItem.ID, postResp)
		}
	}
}
