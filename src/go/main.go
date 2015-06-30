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

	goji.Get("/css/*", handler.GetAssetsHandlerWithContentType("text/css", ASSET_ROOT))
	goji.Get("/js/*", handler.GetAssetsHandlerWithContentType("text/javascript", ASSET_ROOT))

	goji.Get("/page/:page", handler.GetHtmlHandler(ASSET_ROOT, ASSET_ROOT))

	goji.Get("/", handler.GetHtmlHandler(ASSET_ROOT+"/page", ASSET_ROOT))

	goji.Get("/*", http.FileServer(http.Dir(ASSET_ROOT)))

	goji.Serve()
}
