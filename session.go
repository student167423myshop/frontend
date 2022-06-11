package main

import (
	"net/http"
	"net/http/cookiejar"
	"time"

	uuid "github.com/satori/go.uuid"
)

var client http.Client

func init() {
	jar, err := cookiejar.New(nil)
	if err != nil {
		panic(err.Error())
	}
	client = http.Client{
		Jar: jar,
	}
}

// INNER FUNCTIONS
func getSessionId(w http.ResponseWriter, r *http.Request) string {
	var sessionId string
	cookie, err := r.Cookie("sessionId")
	if err != nil {
		sessionId = uuid.NewV4().String()
		cookie := &http.Cookie{
			Name:     "sessionId",
			Value:    sessionId,
			Expires:  time.Now().Add(time.Minute * 60),
			HttpOnly: true,
			Secure:   true,
			SameSite: http.SameSiteStrictMode,
		}
		r.AddCookie(cookie)
		http.SetCookie(w, cookie)
	} else {
		sessionId = cookie.Value
	}
	return sessionId
}
