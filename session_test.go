package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_SettingCookie(t *testing.T) {
	// Arrange
	r := getRouter()
	mockServer := httptest.NewServer(r)
	recorder := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", mockServer.URL+"/", nil)

	expectedCookieName := "test"
	expectedCookieValue := "expected"
	expectedCookie := &http.Cookie{
		Name:  expectedCookieName,
		Value: expectedCookieValue,
	}

	// Act
	req.AddCookie(expectedCookie)
	http.SetCookie(recorder, expectedCookie)

	// Assert
	require.Equal(t, expectedCookieValue, recorder.Result().Cookies()[0].Value)
	require.Equal(t, expectedCookieValue, req.Cookies()[0].Value)
}

func Test_getSessionId(t *testing.T) {
	// Arrange
	r := getRouter()
	mockServer := httptest.NewServer(r)
	recorder := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", mockServer.URL+"/", nil)

	// Acte
	sessionIdFirst := getSessionId(recorder, req)
	sessionIdSecond := getSessionId(recorder, req)

	// Assert
	require.Equal(t, sessionIdFirst, sessionIdSecond)
}
