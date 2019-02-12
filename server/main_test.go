package main

import (
	"myapp/handlers"
	"myapp/handlers/handlerstest"
	"net/http"
	"net/url"
	"testing"
)

func TestPingAPI(t *testing.T) {
	req.URL = &url.URL{Path: "/api/ping"}
	req.Method = "GET"
	req.Header = make(http.Header)
	req.Header.Set("Content-Type", "application/json")
	handlers.PingAPI(res, req)
	data, err := handlerstest.GetJsonData(res.Body)
	if err != nil {
		t.Fatalf("Getting json data error %v", err)
	}
	if data["message"] != "pong" {
		t.Fatalf(`Wrong message; expected "pong", got %q`, data["Message"])
	}
}
