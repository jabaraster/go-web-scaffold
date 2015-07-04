package main

import (
	"./env"
	"./web/handler"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
	"net/http"
)

const (
	ASSET_ROOT = "./assets"
)

func main() {
	env.Dump()

	// modelMuxやhtmlMuxには認証ミドルウェアを仕込む必要がある.

	modelMux := web.New()
	modelMux.Get("/model/order/", handler.AllOrdersHandler)
	modelMux.Get("/model/app-user/", handler.AllAppUsersHandler)

	staticMux := web.New()
	staticMux.Get("/css/*", handler.GetAssetsHandlerWithContentType("text/css", ASSET_ROOT))
	staticMux.Get("/js/*", handler.GetAssetsHandlerWithContentType("text/javascript", ASSET_ROOT))
	staticMux.Get("/*", http.FileServer(http.Dir(ASSET_ROOT)))

	htmlMux := web.New()
	htmlHandler := handler.GetHtmlHandler(ASSET_ROOT+"/html", ASSET_ROOT)
	htmlMux.Get("/page/:page/", htmlHandler)
	htmlMux.Get("/", handler.GetHtmlPathHandler(ASSET_ROOT+"/html/index.html", ASSET_ROOT))

	publicPageMux := web.New()
	publicPageMux.Get("/login", handler.GetHtmlPathHandler(ASSET_ROOT+"/html/login.html", ASSET_ROOT))

	// 各MuxをURLに割り当てる
	// MuxでもURLが登場するので、冗長と言えば冗長.
	// 指定の順番は要注意.
	// 先に割り当てた方が先にマッチするので
	// 競合する指定は優先させたいMuxを先に記述する必要がある.
	defaultMux := goji.DefaultMux
	defaultMux.Handle("/", htmlMux)
	defaultMux.Handle("/model/*", modelMux)
	defaultMux.Handle("/css/*", staticMux)
	defaultMux.Handle("/js/*", staticMux)
	defaultMux.Handle("/page/*", htmlMux)
	defaultMux.Handle("/login", publicPageMux)
	defaultMux.Handle("/*", staticMux)

	goji.Serve()
}
