package middleware

import (
    "github.com/zenazn/goji/web"
    "net/http"
    "github.com/jabaraster/go-web-scaffold/src/go/web/app_session"
)

func Authenticator(c *web.C, h http.Handler) http.Handler {
    fn := func(w http.ResponseWriter, r *http.Request) {
        if !app_session.IsLogin(r) {
            http.Redirect(w, r, "/login", http.StatusSeeOther)
            return;
        }
        h.ServeHTTP(w, r)
    }
    return http.HandlerFunc(fn)
}
