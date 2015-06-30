package main

import (
	"./env"
	"./model"
	"./web/handler"
	"github.com/zenazn/goji"
	"net/http"
)

const (
	ASSET_ROOT = "./assets"
)

func main() {
	env.Dump()

	model.SelectDb()

	goji.Get("/", handler.IndexHtmlHandler)

	css := handler.ContentType{Value: "text/css", AssetsRoot: ASSET_ROOT}
	goji.Get("/css/*", css.StaticHandler)

	js := handler.ContentType{Value: "text/javascript", AssetsRoot: ASSET_ROOT}
	goji.Get("/js/*", js.StaticHandler)

	goji.Get("/*", http.FileServer(http.Dir(ASSET_ROOT)))
	goji.Serve()
}
