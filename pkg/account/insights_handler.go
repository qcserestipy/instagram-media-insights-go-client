package account

import (
	"github.com/qcserestipy/instagram-api-go-client/pkg/client"
	"github.com/qcserestipy/instagram-api-go-client/pkg/sdk-account/v24.0/client/insights"
)

func GetInsightsByAccountID(params *insights.GetInsightsByAccountIDParams) (*insights.GetInsightsByAccountIDOK, error) {
	ctx, instagramClient, err := client.ContextWithClient()
	if err != nil {
		return nil, err
	}
	response, err := instagramClient.Account.Insights.GetInsightsByAccountID(ctx, &insights.GetInsightsByAccountIDParams{
		InstagramAccountID: params.InstagramAccountID,
		Breakdown:          params.Breakdown,
		MetricType:         params.MetricType,
		Period:             params.Period,
		Metric:             params.Metric,
		Since:              params.Since,
		Until:              params.Until,
		Timeframe:          params.Timeframe,
	})
	if err != nil {
		return nil, err
	}
	return response, nil
}
