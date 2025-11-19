package utils

import (
	"fmt"
	"reflect"
	"time"

	access_token_handler "github.com/qcserestipy/instagram-api-go-client/pkg/access"
	"github.com/qcserestipy/instagram-api-go-client/pkg/sdk/v24.0/page/client/access_token"
)

// ParseAPIError extracts detailed error information from API error responses
// Works with any API error response that follows the standard generated pattern
func ParseAPIError(err error, context string) error {
	if err == nil {
		return nil
	}

	// Try to extract payload from error using GetPayload() method
	errValue := reflect.ValueOf(err)
	getPayloadMethod := errValue.MethodByName("GetPayload")

	if !getPayloadMethod.IsValid() {
		return fmt.Errorf("%s: %w", context, err)
	}

	// Call GetPayload()
	results := getPayloadMethod.Call(nil)
	if len(results) == 0 || results[0].IsNil() {
		return fmt.Errorf("%s: %w", context, err)
	}

	payload := results[0].Interface()
	payloadValue := reflect.ValueOf(payload)

	// Handle pointer to struct
	if payloadValue.Kind() == reflect.Ptr {
		payloadValue = payloadValue.Elem()
	}

	if payloadValue.Kind() != reflect.Struct {
		return fmt.Errorf("%s: %w", context, err)
	}

	// Look for Error field
	errorField := payloadValue.FieldByName("Error")
	if !errorField.IsValid() || errorField.IsNil() {
		return fmt.Errorf("%s: %w", context, err)
	}

	// Dereference Error field if it's a pointer
	if errorField.Kind() == reflect.Ptr {
		errorField = errorField.Elem()
	}

	// Extract Code, Type, Message fields
	codeField := errorField.FieldByName("Code")
	typeField := errorField.FieldByName("Type")
	messageField := errorField.FieldByName("Message")

	if !codeField.IsValid() || !typeField.IsValid() || !messageField.IsValid() {
		return fmt.Errorf("%s: %w", context, err)
	}

	return fmt.Errorf("%s - API error (code: %d, type: %s): %s",
		context,
		codeField.Int(),
		typeField.String(),
		messageField.String())
}

// ParseTimestamp tries multiple common layouts and returns the parsed time or an error.
// Accepts timestamps like RFC3339 (with colon offsets) and offsets without colon (+0000).
// Parsed times are converted to the local timezone to match `time.Unix` formatting.
func ParseTimestamp(ts string) (time.Time, error) {
	layouts := []string{
		time.RFC3339,
		"2006-01-02T15:04:05-0700",
		"2006-01-02T15:04:05Z0700",
		"2006-01-02T15:04:05",
	}
	var lastErr error
	for _, l := range layouts {
		if t, err := time.Parse(l, ts); err == nil {
			return t.UTC(), nil
		} else {
			lastErr = err
		}
	}
	return time.Time{}, fmt.Errorf("could not parse timestamp %s: %v", ts, lastErr)
}

func RefreshAccessToken(pageID string) error {
	if pageID == "" {
		return fmt.Errorf("page ID cannot be empty")
	}

	response, err := access_token_handler.GetPageAccessToken(&access_token.GetPageAccessTokenParams{
		PageID: pageID,
		Fields: "access_token",
	})
	if err != nil {
		return ParseAPIError(err, "failed to get page access token")
	}

	fmt.Printf("\nSuccess!\n")
	fmt.Printf("New Access Token: %s\n", response.Payload.AccessToken)

	return nil
}
