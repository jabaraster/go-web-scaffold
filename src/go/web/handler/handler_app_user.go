package handler

import (
	"../../model"
    "../app_session"
	"net/http"
    "time"
)

func AllAppUsersHandler(w http.ResponseWriter, r *http.Request) {
	writeJsonResponse(model.GetAllAppUsers(), w)
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
    http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func AuthenticationHandler(w http.ResponseWriter, r *http.Request) {
    userId := r.FormValue("userId")
    password := r.FormValue("password")

    var ret map[string]interface{}
    if b := model.Authenticate(userId, password); !b {
        ret = map[string]interface{} {
            "error": "authentication was failed.",
        }
    } else {
        ret = map[string]interface{} {
            "error": "",
        }
        lu := app_session.LoginUser {
            UserId: userId,
            LoginDatetime: time.Now(),
        }
        app_session.SetLoginUser(lu, w, r)
    }
    writeJsonResponse(ret, w)
}
