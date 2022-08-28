package service

import "net/http"

type ThirdPartyService interface {
	// Init(path string) string
	buildRequest(url, method, data string) *http.Request
	Do(r *http.Request) (*http.Response, error)
	Close()
}
