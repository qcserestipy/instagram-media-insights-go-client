package api

import (
	"github.com/qcserestipy/instagram-media-insights-go-client/pkg/sdk/v24.0/client/insights"
)

func GetInsightsByMediaID(params *insights.GetInsightsByMediaIDParams) (*insights.GetInsightsByMediaIDOK, error) {
	ctx, client, err := ContextWithClient()
	if err != nil {
		return nil, err
	}
	response, err := client.Insights.GetInsightsByMediaID(ctx, &insights.GetInsightsByMediaIDParams{
		InstagramMediaID: params.InstagramMediaID,
		Metric:           *&params.Metric,
		Period:           *&params.Period,
		Breakdown:        params.Breakdown,
		MetricType:       params.MetricType,
	})
	if err != nil {
		return nil, err
	}
	return response, nil
}
