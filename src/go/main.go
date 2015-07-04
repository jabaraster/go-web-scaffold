package main

import (
	"./env"
	"./web/handler"
	"github.com/zenazn/goji"
	"net/http"
)

const (
	ASSET_ROOT = "./assets"
)

func main() {
	env.Dump()

	goji.Get("/model/order/", handler.AllOrdersHandler)
	goji.Get("/model/app-user/", handler.AllAppUsersHandler)

	goji.Get("/css/*", handler.GetAssetsHandlerWithContentType("text/css", ASSET_ROOT))
	goji.Get("/js/*", handler.GetAssetsHandlerWithContentType("text/javascript", ASSET_ROOT))

    htmlHandler := handler.GetHtmlHandler(ASSET_ROOT+"/html", ASSET_ROOT)
	goji.Get("/page/:page/", htmlHandler)
    goji.Get("/login", handler.GetHtmlPathHandler(ASSET_ROOT+"/html/login.html", ASSET_ROOT))
	goji.Get("/", htmlHandler)

	goji.Get("/*", http.FileServer(http.Dir(ASSET_ROOT)))

	goji.Serve()
}
