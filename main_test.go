package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRouter(t *testing.T) {
	// Arrange
	r := getRouter()
	mockServer := httptest.NewServer(r)

	// Act
	resp, _ := http.Get(mockServer.URL + "/")

	// Assert
	require.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestRouterString(t *testing.T) {
	// Arrange
	expected := "home"
	r := getRouter()
	mockServer := httptest.NewServer(r)
	resp, _ := http.Get(mockServer.URL + "/")

	// Act
	b, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	respString := string(b)

	// Assert
	if !strings.Contains(respString, expected) {
		t.Errorf("Response should contain %s, got %s", expected, respString)
	}
}

func Test_RouterForNonExistentRoute(t *testing.T) {
	// Arrange
	r := getRouter()
	mockServer := httptest.NewServer(r)

	// Act
	resp, _ := http.Post(mockServer.URL+"/", "", nil)

	// Assert
	require.Equal(t, http.StatusMethodNotAllowed, resp.StatusCode)
}

func Test_StaticFileServerRouter(t *testing.T) {
	// Arrange
	expectedContentType := "text/html; charset=utf-8"
	r := getRouter()
	mockServer := httptest.NewServer(r)

	// Act
	resp, _ := http.Get(mockServer.URL + "/static/")

	// Arrange
	require.Equal(t, http.StatusOK, resp.StatusCode)
	contentType := resp.Header.Get("Content-Type")
	require.Equal(t, expectedContentType, contentType)
}
