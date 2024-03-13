package main

import (
	"fmt"
	"net/http"

	"github.com/justinas/nosurf"
)

// WriteToConsole writes a message to the console and then calls the next http.Handler.
//
// next http.Handler.
// http.Handler
func WriteToConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hit the page")
		next.ServeHTTP(w, r)
	})
}

// NoSurf adds CSRF protection to the given http.Handler using the justinas/nosurf
// middleware.
//
// The base cookie is set with HttpOnly, a path of "/", a Secure flag if the app
// is in production, and a SameSite value of http.SameSiteLaxMode.
//
// next http.Handler.
// http.Handler
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,                 // The cookie is only accessible by the HTTP protocol.
		Path:     "/",                  // The cookie is only sent in requests that are targeted at the website root.
		Secure:   app.InProduction,     // The cookie is only sent when using HTTPS.
		SameSite: http.SameSiteLaxMode, // The cookie is only sent in a first-party context.
	})

	return csrfHandler
}

// SessionLoad loads and saves the session data for each request.
// This middleware should be loaded after the NoSurf middleware.
//
// next http.Handler.
// http.Handler
func SessionLoad(next http.Handler) http.Handler {
	return app.Session.LoadAndSave(next)
}
