# Instagram Media Insights Go Client

A Go client library for accessing Instagram Media and Account Insights via the Facebook Graph API v24.0. This client is generated from OpenAPI/Swagger specifications and provides type-safe access to Instagram metrics.

## Features

- 🚀 Type-safe API clients generated from OpenAPI/Swagger specifications
- 📊 **Media Insights**: Fetch insights for individual Instagram posts (reach, likes, comments, etc.)
- 👥 **Account Insights**: Fetch account-level metrics (engagement, demographics, etc.)
- 🔑 Built-in authentication via access token
- 🎯 Support for different metric periods and breakdowns
- 🛠️ Unified singleton client pattern with thread-safe initialization
- 📦 Shared configuration for both API clients

## Installation

```bash
go get github.com/qcserestipy/instagram-media-insights-go-client
```

## Prerequisites

- Go 1.25.0 or higher
- Instagram Business Account
- Facebook Graph API access token with appropriate permissions

## Configuration

Set your Facebook Graph API access token as an environment variable:

```bash
export ACCESS_TOKEN="your-access-token-here"
```

## Usage

### Media Insights Example

Get insights for a specific Instagram post:

```go
package main

import (
    "fmt"
    "github.com/qcserestipy/instagram-media-insights-go-client/pkg/media"
    "github.com/qcserestipy/instagram-media-insights-go-client/pkg/sdk/v24.0/client/insights"
    "github.com/sirupsen/logrus"
)

func main() {
    // Create parameters for media insights
    params := insights.NewGetInsightsByMediaIDParams()
    params.InstagramMediaID = "18112405726596121"
    params.Metric = "reach,likes,comments"
    params.Period = "day"
    
    // Fetch media insights
    resp, err := media.GetInsightsByMediaID(params)
    if err != nil {
        logrus.Fatalf("error: %v", err)
    }
    
    // Process the response
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
```

### Account Insights Example

Get account-level insights with demographic breakdowns:

```go
package main

import (
    "fmt"
    "github.com/qcserestipy/instagram-media-insights-go-client/pkg/account"
    accinsights "github.com/qcserestipy/instagram-media-insights-go-client/pkg/sdk-account/v24.0/client/insights"
    "github.com/sirupsen/logrus"
)

func main() {
    // Create parameters for account insights
    params := accinsights.NewGetInsightsByAccountIDParams()
    params.InstagramAccountID = "17841464714098258"
    params.Metric = "engaged_audience_demographics"
    params.Period = "lifetime"
    
    metricType := "total_value"
    params.MetricType = &metricType
    
    timeframe := "this_month"
    params.Timeframe = &timeframe
    
    breakdown := "country"
    params.Breakdown = &breakdown
    
    // Fetch account insights
    resp, err := account.GetInsightsByAccountID(params)
    if err != nil {
        logrus.Fatalf("error: %v", err)
    }
    
    // Process demographic breakdowns
    if resp.Payload != nil && resp.Payload.Data != nil {
        for _, data := range resp.Payload.Data {
            fmt.Printf("Metric: %s\n", data.Name)
            if data.TotalValue != nil && data.TotalValue.Breakdowns != nil {
                for _, breakdown := range data.TotalValue.Breakdowns {
                    for _, result := range breakdown.Results {
                        fmt.Printf("  %s: %d\n", result.DimensionValues[0], result.Value)
                    }
                }
            }
        }
    }
}
```

### Available Metrics

#### Media Insights Metrics
- `reach` - Total number of unique accounts that have seen the media
- `likes` - Number of likes
- `comments` - Number of comments
- `saved` - Number of times the media was saved
- `engagement` - Total engagement (likes + comments + saves)
- `impressions` - Total number of times the media was seen

#### Account Insights Metrics

**Interaction Metrics** (period: `day`):
- `accounts_engaged` - Number of accounts that interacted with content
- `reach` - Number of unique accounts that saw content
- `likes` - Total likes on posts, reels, and videos
- `comments` - Total comments
- `shares` - Total shares
- `saved` - Total saves
- `total_interactions` - Sum of all interactions
- `views` - Total views of content

**Demographic Metrics** (period: `lifetime`):
- `engaged_audience_demographics` - Demographics of engaged users
- `follower_demographics` - Demographics of followers

Available breakdowns: `age`, `city`, `country`, `gender`, `media_product_type`, `follow_type`

### Period Options

- `day` - Daily metrics (for interaction metrics)
- `lifetime` - Lifetime metrics (for demographic metrics)

### Metric Types

- `total_value` - Simple total with optional breakdowns
- `time_series` - Aggregated by time period

## Project Structure

```
.
├── cmd/
│   └── main/
│       └── main.go                  # Example application
├── pkg/
│   ├── client/
│   │   └── client.go                # Unified client for both APIs
│   ├── config/
│   │   └── config.go                # Shared configuration
│   ├── media/
│   │   └── insights_handler.go      # Media insights functions
│   ├── account/
│   │   └── insights_handler.go      # Account insights functions
│   ├── sdk/
│   │   └── v24.0/
│   │       ├── client/              # Generated media API client
│   │       └── models/              # Generated media data models
│   └── sdk-account/
│       └── v24.0/
│           ├── client/              # Generated account API client
│           └── models/              # Generated account data models
├── api/                             # Git submodule with Swagger specs
│   └── v24.0/
│       ├── swagger.yaml             # Media insights OpenAPI spec
│       └── swagger-account.yaml     # Account insights OpenAPI spec
├── Makefile                         # Build automation
├── go.mod
├── go.sum
└── README.md
```

## Building

```bash
# Build the example application
go build -o ./bin/main cmd/main/main.go

# Or use make
make build

# Run the example
./bin/main
```

## Makefile Commands

```bash
# Generate media insights client
make gen-api-client

# Generate account insights client
make gen-api-client-account

# Generate both clients
make gen-all-clients

# Clean generated code
make cleanup

# Build the application
make build

# Run tests
make test

# Update git submodule with latest specs
make update-submodule

# Clean and regenerate everything
make all
```

## API Reference

### Unified Client

#### `client.Get()`

Returns a singleton instance of the unified Instagram client with both Media and Account APIs.

**Returns:** `(*client.InstagramClient, error)`

The `InstagramClient` struct provides:
- `Media` - Instagram Media Insights API client
- `Account` - Instagram Account Insights API client

### Media Insights

#### `media.GetInsightsByMediaID(params *insights.GetInsightsByMediaIDParams)`

Fetches insights for a specific Instagram media by its ID.

**Parameters:**
- `params.InstagramMediaID` (string, required) - The Instagram media ID
- `params.Metric` (string, required) - Comma-separated list of metrics
- `params.Period` (string, required) - Time period for metrics (`day`, `week`, `days_28`, `lifetime`)
- `params.Breakdown` (*string, optional) - Breakdown dimension
- `params.MetricType` (*string, optional) - Metric aggregation type

**Returns:** `(*insights.GetInsightsByMediaIDOK, error)`

### Account Insights

#### `account.GetInsightsByAccountID(params *insights.GetInsightsByAccountIDParams)`

Fetches insights for an Instagram account.

**Parameters:**
- `params.InstagramAccountID` (string, required) - The Instagram account ID
- `params.Metric` (string, required) - Comma-separated list of metrics
- `params.Period` (string, required) - Time period (`day` or `lifetime`)
- `params.MetricType` (*string, optional) - `total_value` or `time_series`
- `params.Breakdown` (*string, optional) - Breakdown dimension
- `params.Timeframe` (*string, optional) - Required for demographics: `this_month`, `this_week`
- `params.Since` (*int64, optional) - Unix timestamp for start of range
- `params.Until` (*int64, optional) - Unix timestamp for end of range

**Returns:** `(*insights.GetInsightsByAccountIDOK, error)`

## Development

The API clients are generated from OpenAPI/Swagger specifications maintained in a separate repository and included as a git submodule.

### Updating the API Specifications

```bash
# Update the submodule to the latest version
make update-submodule

# Or manually
cd api && git pull origin main && cd ..

# Regenerate both clients
make all
```

### Code Generation

The project uses [go-swagger](https://goswagger.io/) to generate API clients from OpenAPI specifications:

```bash
# Generate only media insights client
make gen-api-client

# Generate only account insights client  
make gen-api-client-account

# Generate both
make gen-all-clients
```

## Dependencies

- [go-openapi/runtime](https://github.com/go-openapi/runtime) - OpenAPI runtime for Go
- [go-openapi/strfmt](https://github.com/go-openapi/strfmt) - String formatting utilities
- [sirupsen/logrus](https://github.com/sirupsen/logrus) - Structured logging

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

[Add your license here]

## Related Projects

- [instagram-media-insights-swagger](https://github.com/qcserestipy/instagram-media-insights-swagger) - OpenAPI/Swagger specifications

## Support

For issues and questions:
- Open an issue on GitHub
- Check the [Facebook Graph API documentation](https://developers.facebook.com/docs/instagram-api)

---

Made with ❤️ for the Instagram developer community