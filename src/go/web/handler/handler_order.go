package handler

import (
	"../../model"
	"net/http"
)

func AllOrdersHandler(w http.ResponseWriter, r *http.Request) {
	writeJsonResponse(w, model.GetAllOrders())
}
