package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSetCookie(t *testing.T) {
	recorder := httptest.NewRecorder()
	http.SetCookie(recorder, &http.Cookie{Name: "test", Value: "expected"})
	cookie := recorder.Result().Cookies()[0].Value
	require.Equal(t, cookie, "expected")
}
