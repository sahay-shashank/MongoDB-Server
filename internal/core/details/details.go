package details

import (
	"encoding/json"
	"fmt"
	"os"
)

type APIDetails struct {
	Error             bool
	StatusCode        int         `json:"statusCode"`
	Message           string      `json:"message"`
	AdditionalDetails interface{} `json:"additionalDetails,omitempty"`
}

type HTTPJSONDetails struct {
	HTTPStatusCode int        `json:"HTTPStatusCode"`
	Message        string     `json:"message"`
	Details        APIDetails `json:"details"`
}

func (a *APIDetails) ToJSON() ([]byte, error) {
	data, err := json.Marshal(a)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (a *APIDetails) LogToStderr() {
	data, err := a.ToJSON()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to serialize APIError: %v\n", err)
		return
	}
	fmt.Fprintln(os.Stderr, string(data))
}

func (a *APIDetails) LogToStdout() {
	data, err := a.ToJSON()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to serialize APIError: %v\n", err)
		return
	}
	fmt.Fprintln(os.Stdout, string(data))
}
