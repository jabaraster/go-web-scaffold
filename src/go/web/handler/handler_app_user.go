package handler

import (
	"../../model"
	"net/http"
)

func AllAppUsersHandler(w http.ResponseWriter, r *http.Request) {
	writeJsonResponse(w, model.GetAllAppUsers())
}
