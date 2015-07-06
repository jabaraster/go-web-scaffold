package middleware

import (
    "testing"
    "fmt"
    "github.com/jabaraster/go-web-scaffold/src/go/web/app_session"
)

func Test_(t *testing.T) {
    m := map[string]interface{} {
        "a": "aaaaaaa",
        "b": "bbbbbbbb",
    }
    fmt.Println(app_session.LoginUser{})
    t.Error(m)
    delete(m, "a")
    t.Error(m)

    t.Fail()
}