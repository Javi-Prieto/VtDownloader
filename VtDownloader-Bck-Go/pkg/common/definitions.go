package common

import "net/http"

// Definition for the error message
type CommonErrorMessage struct {
	ErrorMessage string `json:"errorMessage"`
}

// Definition of route for the handlers
type Route struct {
	Path    string
	Handler http.HandlerFunc
	Method  string
}
