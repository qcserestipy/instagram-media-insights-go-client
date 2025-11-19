package account

import (
	"context"

	"github.com/qcserestipy/instagram-api-go-client/pkg/sdk/v24.0/account/client/insights"
)

func (s *Service) GetInsightsByAccountID(
	ctx context.Context,
	params *insights.GetInsightsByAccountIDParams,
) (*insights.GetInsightsByAccountIDOK, error) {
	response, err := s.client.Account.Insights.GetInsightsByAccountID(ctx, &insights.GetInsightsByAccountIDParams{
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
