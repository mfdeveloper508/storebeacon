package api

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

// UnmarshalPayload converts an io into a struct instance.
func UnmarshalPayload(in io.Reader, model interface{}) error {
	content, err := ioutil.ReadAll(in)

	if err != nil {
		return err
	}

	err = json.Unmarshal(content, &model)

	return err
}

// MarshalPayload writes an interface to an io.Writer
func MarshalPayload(out io.Writer, model interface{}) error {
	output, err := json.Marshal(model)

	if err != nil {
		return err
	}

	_, err = out.Write(output)

	return err
}

// ServeJSON takes an interface and writes the JSON response
func ServeJSON(w http.ResponseWriter, model interface{}) error {
	w.Header().Set("Content-Type", "application/json")

	err := MarshalPayload(w, model)

	if err != nil {
		http.Error(w, err.Error(), 500)
	}
	return err
}

// ServeJSONWithStatusCode takes an interface and status code
// This is just an extension of the ServeJSON method
func ServeJSONWithStatusCode(w http.ResponseWriter, model interface{}, statusCode int) error {
	w.WriteHeader(statusCode)

	// 500 errors are already handled by ServeJSON
	// so no further action is needed
	err := ServeJSON(w, model)
	return err
}

// ServeJSONErrors takes an error response and parses errors into string messages
func ServeJSONErrors(w http.ResponseWriter, response ErrorResponse) error {
	// Convert errors into just strings
	parsedResponse := struct {
		Success bool     `json:"success"`
		Errors  []string `json:"errors"`
	}{
		Success: response.Success,
	}

	for _, err := range response.Errors {
		parsedResponse.Errors = append(parsedResponse.Errors, err.Error())
	}

	return ServeJSON(w, parsedResponse)
}

type GenericResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type ErrorResponse struct {
	Success bool    `json:"success"`
	Errors  []error `json:"errors"`
}
