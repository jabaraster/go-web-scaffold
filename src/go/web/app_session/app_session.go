package app_session

import (
    "time"
    "net/http"
    "github.com/gorilla/sessions"
    "fmt"
    "encoding/json"
)


var store = sessions.NewCookieStore([]byte("secret-token"))

const (
    defaultSessionName = "default-session"

    sessionKey_loginUser = "sessionKey_loginUser"
)

type LoginUser struct {
    UserId string
    LoginDatetime time.Time
}

func SetLoginUser(loginUser LoginUser, w http.ResponseWriter, r *http.Request) {
    str, err := json.Marshal(loginUser)
    if err != nil {
        panic(err)
    }
    save(sessionKey_loginUser, str, w, r)
}

func IsLogin(r *http.Request) bool {
    session := mustGetSession(r)
    s, have := session.Values[sessionKey_loginUser]
    fmt.Println("★ ", s, have)
    return have
}

func get(key string, r *http.Request) interface{} {
    session := mustGetSession(r)
    return session.Values[key]
}

func save(key string, value interface{}, w http.ResponseWriter, r *http.Request) {
    session, err := store.Get(r, defaultSessionName)
    if err != nil {
        panic(err)
    }
    session.Values[key] = value
    if e := session.Save(r, w); e != nil {
        panic(e)
    }
}

func mustGetSession(r *http.Request) *sessions.Session {
    session, err := store.Get(r, defaultSessionName)
    if err != nil {
        panic(err)
    }
    return session
}
