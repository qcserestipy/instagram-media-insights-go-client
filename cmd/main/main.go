package main

import (
	"fmt"
	"reflect"

	"github.com/qcserestipy/instagram-media-insights-go-client/pkg/media"
	"github.com/qcserestipy/instagram-media-insights-go-client/pkg/sdk/v24.0/client/insights"

	"github.com/qcserestipy/instagram-media-insights-go-client/pkg/account"
	accinsights "github.com/qcserestipy/instagram-media-insights-go-client/pkg/sdk-account/v24.0/client/insights"
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
}
