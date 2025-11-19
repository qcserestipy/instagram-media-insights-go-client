package instagram

import (
	"fmt"

	"github.com/qcserestipy/instagram-api-go-client/pkg/account"
	"github.com/qcserestipy/instagram-api-go-client/pkg/sdk/v24.0/account/client/insights"
	"github.com/qcserestipy/instagram-api-go-client/pkg/sdk/v24.0/account/models"
	"github.com/qcserestipy/instagram-api-go-client/pkg/utils"
)

func GetDemographics(accountID string, kind string) (*Demographics, error) {

	if kind != "follower_demographics" && kind != "engaged_audience_demographics" {
		return nil, fmt.Errorf("invalid kind: %s", kind)
	}

	breakdowns := []string{"country", "age", "gender", "city"}
	responseBreakdowns := []*models.AccountBreakdown{}
	metric_type := "total_value"
	timeframe := "this_month"
	for _, breakdown := range breakdowns {
		insightsResponse, err := account.GetInsightsByAccountID(&insights.GetInsightsByAccountIDParams{
			InstagramAccountID: accountID,
			Metric:             kind,
			Period:             "lifetime",
			Breakdown:          &breakdown,
			Timeframe:          &timeframe,
			MetricType:         &metric_type,
		})
		if err != nil {
			return nil, fmt.Errorf("failed to %v", utils.ParseAPIError(err, "get "+kind))
		}
		responseBreakdowns = append(responseBreakdowns, insightsResponse.Payload.Data[0].TotalValue.Breakdowns[0])
	}
	demographics := &Demographics{
		AgeRanges: make(map[string]int64),
		Genders:   make(map[string]int64),
		Countries: make(map[string]int64),
		Cities:    make(map[string]int64),
	}

	for _, breakdown := range responseBreakdowns {
		if len(breakdown.DimensionKeys) == 0 {
			continue
		}

		switch breakdown.DimensionKeys[0] {
		case "age":
			for _, result := range breakdown.Results {
				if len(result.DimensionValues) > 0 {
					ageRange := result.DimensionValues[0]
					demographics.AgeRanges[ageRange] = result.Value
				}
			}
		case "gender":
			for _, result := range breakdown.Results {
				if len(result.DimensionValues) > 0 {
					gender := result.DimensionValues[0]
					demographics.Genders[gender] = result.Value
				}
			}
		case "country":
			for _, result := range breakdown.Results {
				if len(result.DimensionValues) > 0 {
					country := result.DimensionValues[0]
					demographics.Countries[country] = result.Value
				}
			}
		case "city":
			for _, result := range breakdown.Results {
				if len(result.DimensionValues) > 0 {
					city := result.DimensionValues[0]
					demographics.Cities[city] = result.Value
				}
			}
		}
	}

	return demographics, nil
}

func GetAccountDemographics(accountID string) (*AccountDemographics, error) {
	follower, err := GetDemographics(accountID, "follower_demographics")
	if err != nil {
		return nil, err
	}
	engaged, err := GetDemographics(accountID, "engaged_audience_demographics")
	if err != nil {
		return nil, err
	}
	return &AccountDemographics{
		Follower: follower,
		Engaged:  engaged,
	}, nil
}
