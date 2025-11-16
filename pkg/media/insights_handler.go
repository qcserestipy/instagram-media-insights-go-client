package media

import (
	"github.com/qcserestipy/instagram-api-go-client/pkg/client"
	"github.com/qcserestipy/instagram-api-go-client/pkg/sdk/v24.0/client/insights"
)

func GetInsightsByMediaID(params *insights.GetInsightsByMediaIDParams) (*insights.GetInsightsByMediaIDOK, error) {
	ctx, instagramClient, err := client.ContextWithClient()
	if err != nil {
		return nil, err
	}
	response, err := instagramClient.Media.Insights.GetInsightsByMediaID(ctx, &insights.GetInsightsByMediaIDParams{
		InstagramMediaID: params.InstagramMediaID,
		Metric:           params.Metric,
		Period:           params.Period,
		Breakdown:        params.Breakdown,
		MetricType:       params.MetricType,
	})
	if err != nil {
		return nil, err
	}
	return response, nil
}
