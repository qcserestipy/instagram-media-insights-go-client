# Instagram Media Insights Go Client

A Go client library for accessing Instagram Media Insights via the Facebook Graph API v24.0. This client is generated from OpenAPI/Swagger specifications and provides type-safe access to Instagram media metrics.

## Features

- 🚀 Type-safe API client generated from OpenAPI/Swagger specifications
- 🔑 Built-in authentication via access token
- 📊 Fetch insights for Instagram media (reach, likes, comments, etc.)
- 🎯 Support for different metric periods (day, week, etc.)
- 🛠️ Singleton client pattern with thread-safe initialization

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

### Basic Example

```go
package main

import (
    "fmt"
    "github.com/qcserestipy/instagram-media-insights-go-client/pkg/api"
    "github.com/qcserestipy/instagram-media-insights-go-client/pkg/sdk/v24.0/client/insights"
    "github.com/sirupsen/logrus"
)

func main() {
    // Create parameters for the insights request
    params := insights.NewGetInsightsByMediaIDParams()
    params.InstagramMediaID = "18112405726596121"
    params.Metric = "reach,likes,comments"
    params.Period = "day"
    
    // Fetch insights
    resp, err := api.GetInsightsByMediaID(params)
    if err != nil {
        logrus.Fatalf("fatal error: %v", err)
    }
    
    // Process the response
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
```

### Available Metrics

Common metrics you can request:
- `reach` - Total number of unique accounts that have seen the media
- `likes` - Number of likes
- `comments` - Number of comments
- `saved` - Number of times the media was saved
- `engagement` - Total engagement (likes + comments + saves)
- `impressions` - Total number of times the media was seen

### Period Options

- `day` - Daily metrics
- `week` - Weekly metrics
- `days_28` - 28-day metrics
- `lifetime` - Lifetime metrics

## Project Structure

```
.
├── cmd/
│   └── main/
│       └── main.go              # Example application
├── pkg/
│   ├── api/
│   │   ├── client.go            # Client initialization and configuration
│   │   └── insights_handler.go  # High-level API functions
│   └── sdk/
│       └── v24.0/
│           ├── client/          # Generated API client
│           └── models/          # Generated data models
├── api/                         # Git submodule with Swagger specs
├── go.mod
├── go.sum
└── README.md
```

## Building

```bash
# Build the example application
go build -o ./bin/main cmd/main/main.go

# Run the example
./bin/main
```

## API Reference

### `api.GetClient()`

Returns a singleton instance of the Instagram Media Insights API client.

**Returns:** `(*InstagramMediaInsightsAPI, error)`

### `api.GetInsightsByMediaID(params *insights.GetInsightsByMediaIDParams)`

Fetches insights for a specific Instagram media by its ID.

**Parameters:**
- `params.InstagramMediaID` (string, required) - The Instagram media ID
- `params.Metric` (string, required) - Comma-separated list of metrics
- `params.Period` (string, optional) - Time period for metrics

**Returns:** `(*insights.GetInsightsByMediaIDOK, error)`

## Development

The API client is generated from OpenAPI/Swagger specifications maintained in a separate repository and included as a git submodule.

### Updating the API Specification

```bash
# Update the submodule to the latest version
git submodule update --remote api/v24.0

# Regenerate the client if needed
# (Add your code generation commands here)
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