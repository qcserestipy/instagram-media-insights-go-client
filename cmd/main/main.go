package main

import (
	"fmt"
	"reflect"

	"github.com/qcserestipy/instagram-api-go-client/pkg/account"
	"github.com/qcserestipy/instagram-api-go-client/pkg/media"
	accinsights "github.com/qcserestipy/instagram-api-go-client/pkg/sdk/v24.0/account/client/insights"
	mediaModel "github.com/qcserestipy/instagram-api-go-client/pkg/sdk/v24.0/account/client/media"
	"github.com/qcserestipy/instagram-api-go-client/pkg/sdk/v24.0/media/client/insights"
	"github.com/sirupsen/logrus"
)

func main() {

	// Now you can use the client to make requests
	params := insights.NewGetInsightsByMediaIDParams()
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
			params := insights.NewGetInsightsByMediaIDParams()
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

	mediaResp, err := account.GetMediaByUserID(&mediaModel.GetMediaByUserIDParams{
		InstagramAccountID: "17841464714098258",
	})
	if err != nil {
		logrus.Fatalf("fatal error: %v", err)
	}
	fmt.Printf("Success! Media Response: %+v\n", mediaResp)
	if mediaResp.Payload != nil && mediaResp.Payload.Data != nil {
		for _, mediaItem := range mediaResp.Payload.Data {
			fmt.Printf("Media ID: %s\n", mediaItem.ID)

			params := insights.NewGetInsightsByMediaIDParams()
			params.InstagramMediaID = mediaItem.ID
			params.Metric = "views,likes,comments"
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
		}
	}
}
