package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

// func ParseBody(r *http.Request, x interface{}) {
// 	body, err := io.ReadAll(r.Body)
// 	if err == nil {
// 		err := json.Unmarshal([]byte(body), x)
// 		if err != nil {
// 			return
// 		}
// 	}
// }

func ParseBody(r *http.Request, x interface{}) error {
	// Read the entire contents of the request body into the `body` variable.
	body, err := io.ReadAll(r.Body)
	if err != nil {
		// If there is an error reading the request body, return the error.
		return err
	}

	// Attempt to unmarshal the JSON data from the `body` variable into the `x` interface.
	err = json.Unmarshal([]byte(body), x)
	if err != nil {
		// If there is an error unmarshalling the JSON data, return the error.
		return err
	}

	return nil
}
