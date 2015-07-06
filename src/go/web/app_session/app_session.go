package app_session

import (
    "time"
    "net/http"
    "github.com/gorilla/sessions"
    "encoding/json"
    "fmt"
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
    d, err := json.Marshal(loginUser)
    if err != nil {
        panic(err)
    }
    save(sessionKey_loginUser, string(d), w, r)
}

func IsLogin(r *http.Request) bool {
    session := mustGetSession(r)
    fmt.Println("IsLogin -> ", session)
    _, found := session.Values[sessionKey_loginUser]
    return found
}

func UnsetLoginUser(w http.ResponseWriter, r *http.Request) {
    remove(sessionKey_loginUser, w, r)
}

func save(key string, value interface{}, w http.ResponseWriter, r *http.Request) {
    session := mustGetSession(r)
    fmt.Println("save before -> ", session)
    session.Values[key] = value
    fmt.Println("save after -> ", session)
    if e := session.Save(r, w); e != nil {
        panic(e)
    }
}

func get(key string, r *http.Request) interface{} {
    session := mustGetSession(r)
    return session.Values[key]
}

func remove(key string, w http.ResponseWriter, r *http.Request) {
    session := mustGetSession(r)
    fmt.Println("remove before -> ", session)
    delete(session.Values, key)
    fmt.Println("remove after -> ", session)
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
