package main

import (
	"context"
	"fmt"
	"net/url"
	"os"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/qcserestipy/insta-swagger/pkg/sdk/v24.0/client"
	"github.com/qcserestipy/insta-swagger/pkg/sdk/v24.0/client/insights"
	"github.com/sirupsen/logrus"
)

func main() {
	// read access token from environment variable or configuration
	accessToken := os.Getenv("ACCESS_TOKEN")
	if accessToken == "" {
		logrus.Fatal("ACCESS_TOKEN environment variable is not set")
	}

	// Create the client configuration
	cfg := client.Config{
		URL: &url.URL{
			Scheme: "https",
			Host:   "graph.facebook.com",
			Path:   "/v24.0",
		},
		// Pass the access token as a query parameter
		AuthInfo: httptransport.APIKeyAuth("access_token", "query", accessToken),
	}

	// Create the client
	apiClient := client.New(cfg)

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
