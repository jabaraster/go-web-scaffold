package handler

import (
	"../../model"
	"../app_session"
	"net/http"
	"strings"
	"time"
)

func AddAppUserHandler(w http.ResponseWriter, r *http.Request) {
	userId := strings.TrimSpace(r.FormValue("userId"))
	password := strings.TrimSpace(r.FormValue("password"))
	passwordConfirmation := strings.TrimSpace(r.FormValue("passwordConfirmation"))
	if password != passwordConfirmation {
		writeErrorJsonResponse("パスワードが一致しません.", w)
		return
	}

	newUser, invalidValue := model.AddAppUser(userId, password)
	if invalidValue != nil {
		writeErrorJsonResponse(invalidValue.GetDescription(), w)
		return
	}
	ret := map[string]interface{}{
		"error":     "",
		"newEntity": newUser,
	}
	writeJsonResponse(ret, w)
}

func AllAppUsersHandler(w http.ResponseWriter, r *http.Request) {
	writeJsonResponse(model.GetAllAppUsers(), w)
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	app_session.UnsetLoginUser(w, r)
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func AuthenticationHandler(w http.ResponseWriter, r *http.Request) {
	userId := r.FormValue("userId")
	password := r.FormValue("password")

	var ret map[string]interface{}
	if b := model.Authenticate(userId, password); !b {
		ret = map[string]interface{}{
			"error": "authentication was failed.",
		}
	} else {
		ret = map[string]interface{}{
			"error": "",
		}
		lu := app_session.LoginUser{
			UserId:        userId,
			LoginDatetime: time.Now(),
		}
		app_session.SetLoginUser(lu, w, r)
	}
	writeJsonResponse(ret, w)
}
