package service

import (
	"bytes"
	logger "log"
	"net/http"
)

var log = logger.Logger{}

type HttpService struct {
}

func (service HttpService) buildRequest(url string, method string, data string) *http.Request {
	req, _ := http.NewRequest(method, url, bytes.NewBuffer([]byte(data)))
	return req
}

func (service HttpService) Do(req *http.Request) (*http.Response, error) {
	client := &http.Client{}

	return client.Do(req)
}

func (service HttpService) Close() {

}
