package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_SetCookie(t *testing.T) {
	recorder := httptest.NewRecorder()
	http.SetCookie(recorder, &http.Cookie{Name: "test", Value: "expected"})
	cookie := recorder.Result().Cookies()[0].Value
	require.Equal(t, "expected", cookie)
}

func Test_CookieJar(t *testing.T) {
	cookie := &http.Cookie{
		Name:   "token",
		Value:  "some_token",
		MaxAge: 300,
	}

	cookie2 := &http.Cookie{
		Name:   "clicked",
		Value:  "true",
		MaxAge: 300,
	}

	req, err := http.NewRequest("GET", "http://google.com", nil)
	if err != nil {
		t.Errorf("Got error")
	}
	req.AddCookie(cookie)
	req.AddCookie(cookie2)
	for _, c := range req.Cookies() {
		fmt.Println(c)
	}
	resp, err := client.Do(req)
	if err != nil {
		t.Errorf("Got error")
	}

	cookieStr, _ := req.Cookie("token")
	require.Equal(t, "some_token", cookieStr.Value)

	defer resp.Body.Close()
	fmt.Printf("StatusCode: %d\n", resp.StatusCode)
}
