package notifications

import (
	"encoding/base64"
	"encoding/json"
)

type (
	// TMALinkingScheme defines the JSON representation of the linking scheme for the Shoppigram TMA
	// Reference:
	// https://www.notion.so/WIP-RFC-TWA-page-linking-and-data-transfer-format-d8ba392b9b19475b80be8aeed415ea30?pvs=4
	TMALinkingScheme struct {
		PageName string `json:"page_name"`
		// We have to use any in here, because of Go constraints
		// The actual type is any "simple" JSON value, like string, number, boolean, or null
		PageData map[string]any `json:"page_data"`
	}
)

// ToBase64String converts the TMALinkingScheme to a base64 string
// That can be decoded and parsed as JSON
func (t TMALinkingScheme) ToBase64String() (string, error) {
	data, err := json.Marshal(t)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(data), nil
}
