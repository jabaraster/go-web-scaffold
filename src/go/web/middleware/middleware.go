package middleware

import (
	"../app_session"
	"github.com/zenazn/goji/web"
	"net/http"
)

func PageAuthenticator(c *web.C, h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if !app_session.IsLogin(r) {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

func JsonAuthenticator(c *web.C, h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if !app_session.IsLogin(r) {
			http.Error(w, "modele not found.", http.StatusNotFound)
			return
		}
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
