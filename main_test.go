package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestRouter(t *testing.T) {
	r := getRouter()
	mockServer := httptest.NewServer(r)
	resp, err := http.Get(mockServer.URL + "/")

	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Status should be ok, got %d", resp.StatusCode)
	}

	defer resp.Body.Close()
}

func TestRouterString(t *testing.T) {
	r := getRouter()
	mockServer := httptest.NewServer(r)
	resp, err := http.Get(mockServer.URL + "/")

	if err != nil {
		t.Fatal(err)
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	respString := string(b)
	expected := "Home"

	if !strings.Contains(respString, expected) {
		t.Errorf("Response should contain %s, got %s", expected, respString)
	}
}

func TestRouterForNonExistentRoute(t *testing.T) {
	r := getRouter()
	mockServer := httptest.NewServer(r)
	resp, err := http.Post(mockServer.URL+"/", "", nil)

	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("Status should be 405, got %d", resp.StatusCode)
	}

	defer resp.Body.Close()
}

func TestStaticFileServerRouter(t *testing.T) {
	r := getRouter()
	mockServer := httptest.NewServer(r)

	resp, err := http.Get(mockServer.URL + "/static/")
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Status should be 200, got %d", resp.StatusCode)
	}

	contentType := resp.Header.Get("Content-Type")
	expectedContentType := "text/html; charset=utf-8"

	if expectedContentType != contentType {
		t.Errorf("Wrong content type, expected %s, got %s",
			expectedContentType, contentType)
	}
}
