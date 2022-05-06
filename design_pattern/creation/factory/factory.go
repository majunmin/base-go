// Created By: junmin.ma
// Description: <description>
// Date: 2022-04-22 23:47
package factory

import (
	"net/http"
	"net/http/httptest"
)

type Doer interface {
	Do(reqeust *http.Request) (*http.Response, error)
}

func NewHttpClient() Doer {
	return &http.Client{}
}

type mockHttpClient struct {
}

func (m *mockHttpClient) Do(reqeust *http.Request) (*http.Response, error) {
	req := httptest.NewRecorder()
	return req.Result(), nil
}

func NewMockHttpClient() Doer {
	return &mockHttpClient{}
}
