package instagram

import (
	"context"
	"fmt"
	"time"

	"github.com/qcserestipy/instagram-api-go-client/pkg/account"
	"github.com/qcserestipy/instagram-api-go-client/pkg/sdk/v24.0/account/client/insights"
	"github.com/qcserestipy/instagram-api-go-client/pkg/sdk/v24.0/account/client/user"
	"github.com/qcserestipy/instagram-api-go-client/pkg/utils"
)

func GetFollowers(ctx context.Context, svc *account.Service, accountID string) (int64, error) {
	fields := "followers_count"
	insightsResponse, err := svc.GetUserByID(ctx, &user.GetInstagramUserByIDParams{
		InstagramAccountID: accountID,
		Fields:             &fields,
	})
	if err != nil {
		return -1, fmt.Errorf("failed to %v", utils.ParseAPIError(err, "get "+"follower count"))
	}
	return insightsResponse.Payload.FollowersCount, nil
}

// GetFollowerDynamics returns the net follower change (follows - unfollows)
// for the provided account over the given time range string.
// Supported ranges: last_30_days, last_21_days, last_14_days,last_7_days, yesterday, today
func GetFollowerDynamics(ctx context.Context, svc *account.Service, accountID string, rangeStr string) (*FollowerDynamics, error) {
	metrics := "follows_and_unfollows"
	breakdown := "follow_type"
	metricType := "total_value"
	var since int64
	var until int64

	now := time.Now().UTC()
	switch rangeStr {
	case "last_14_days":
		since = now.Add(-14 * 24 * time.Hour).Unix()
		until = now.Unix()
	case "last_21_days":
		since = now.Add(-21 * 24 * time.Hour).Unix()
		until = now.Unix()
	case "last_30_days":
		since = now.Add(-30 * 24 * time.Hour).Unix()
		until = now.Unix()
	case "last_7_days":
		since = now.Add(-7 * 24 * time.Hour).Unix()
		until = now.Unix()
	case "yesterday":
		startToday := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
		startYesterday := startToday.AddDate(0, 0, -1)
		since = startYesterday.Unix()
		until = startToday.Unix() - 1
	case "today":
		startToday := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
		since = startToday.Unix()
		until = now.Unix()
	default:
		return nil, fmt.Errorf("unsupported range: %s", rangeStr)
	}

	insightsResponse, err := svc.GetInsightsByAccountID(ctx, &insights.GetInsightsByAccountIDParams{
		InstagramAccountID: accountID,
		Metric:             metrics,
		Period:             "day",
		Breakdown:          &breakdown,
		MetricType:         &metricType,
		Since:              &since,
		Until:              &until,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to %v", utils.ParseAPIError(err, "get follower dynamics"))
	}

	if insightsResponse == nil || insightsResponse.Payload == nil || len(insightsResponse.Payload.Data) == 0 {
		return nil, fmt.Errorf("no data returned in payload for follower dynamics")
	}

	data := insightsResponse.Payload.Data[0]
	// If TotalValue or Breakdowns nil -> no results
	if data.TotalValue == nil {
		return nil, fmt.Errorf("no total value returned in payload for follower dynamics")
	}
	breakdowns := data.TotalValue.Breakdowns
	if len(breakdowns) == 0 {
		return nil, fmt.Errorf("no breakdowns returned in payload for follower dynamics")
	}

	var followsTotal int64
	var unfollowsTotal int64

	for _, bd := range breakdowns {
		for _, res := range bd.Results {
			if len(res.DimensionValues) > 0 {
				switch res.DimensionValues[0] {
				case "FOLLOWER":
					followsTotal += res.Value
				case "NON_FOLLOWER":
					unfollowsTotal += res.Value
				}
			}
		}
	}
	net := followsTotal - unfollowsTotal
	parsedSince, err := utils.ParseTimestamp(time.Unix(since, 0).Format(time.RFC3339))
	if err != nil {
		return nil, fmt.Errorf("could not parse since timestamp: %v", err)
	}
	parsedUntil, err := utils.ParseTimestamp(time.Unix(until, 0).Format(time.RFC3339))
	if err != nil {
		return nil, fmt.Errorf("could not parse until timestamp: %v", err)
	}
	return &FollowerDynamics{
		NewFollowers: followsTotal,
		Unfollowers:  unfollowsTotal,
		NetFollowers: net,
		Since:        parsedSince,
		Until:        parsedUntil,
	}, nil
}
