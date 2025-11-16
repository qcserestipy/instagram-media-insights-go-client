package main

import (
	"context"
	"fmt"

	"github.com/qcserestipy/insta-swagger/pkg/sdk/v24.0/client/insights"
	"github.com/sirupsen/logrus"
)

func main() {

	// Now you can use the client to make requests
	params := insights.NewGetInsightsByMediaIDParams()
	params.InstagramMediaID = "18112405726596121"
	params.Metric = "reach,likes,comments"
	params.Period = "day"

	ctx := context.Background()
	resp, err := apiClient.Insights.GetInsightsByMediaID(ctx, params)
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
