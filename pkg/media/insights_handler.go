package media

import (
	"context"

	"github.com/qcserestipy/instagram-api-go-client/pkg/sdk/v24.0/media/client/insights"
)

func (s *Service) GetInsightsByMediaID(ctx context.Context, params *insights.GetInsightsByMediaIDParams) (*insights.GetInsightsByMediaIDOK, error) {
	response, err := s.client.Media.Insights.GetInsightsByMediaID(ctx, &insights.GetInsightsByMediaIDParams{
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
