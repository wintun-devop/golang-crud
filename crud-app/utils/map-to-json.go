package utils

import (
	"encoding/json"
	"fmt"
)

// Returns the JSON string or an error.
func MapToJSON(data map[string]interface{}) (string, error) {
	bytes, err := json.Marshal(data)
	if err != nil {
		return "", fmt.Errorf("failed to convert map to JSON: %w", err)
	}
	return string(bytes), nil
}

// MapToPrettyJSON converts a map into a pretty-formatted JSON string.
func MapToPrettyJSON(data map[string]interface{}) (string, error) {
	bytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return "", fmt.Errorf("failed to convert map to pretty JSON: %w", err)
	}
	return string(bytes), nil
}
