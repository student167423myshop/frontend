package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"time"
)

var client http.Client

func init() {
	jar, err := cookiejar.New(nil)
	if err != nil {
		log.Fatalf("Got error while creating cookie jar %s", err.Error())
	}
	client = http.Client{
		Jar: jar,
	}
}

func getSessionId(r *http.Request) string {
	fmt.Printf(" -- Sesje przed: \n.")
	for _, c := range r.Cookies() {
		fmt.Println(c)
	}

	cookie, err := r.Cookie("sessionId")
	var sessionId string
	if err != nil {
		fmt.Printf(" -- Tworzenie nowej sesji.")
		sessionId = getNewSessionId()
		cookie := &http.Cookie{
			Name:     "sessionId",
			Value:    sessionId,
			Expires:  time.Now().Add(time.Minute * 60),
			HttpOnly: true,
			Secure:   true,
			SameSite: http.SameSiteStrictMode,
		}
		r.AddCookie(cookie)
	} else {
		sessionId = cookie.Value
	}

	fmt.Printf(" -- Sesje po: \n.")
	for _, c := range r.Cookies() {
		fmt.Println(c)
	}

	return sessionId
}

func getNewSessionId() string {
	b := make([]byte, 32)
	io.ReadFull(rand.Reader, b)
	sessionIdB := base64.URLEncoding.EncodeToString(b)
	sessionId := url.QueryEscape(sessionIdB)
	fmt.Printf(" -- sessionId: %s\n", sessionId)
	return sessionId
}
