package main

import (
	"./env"
	"./web/handler"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
	"net/http"
    "github.com/jabaraster/go-web-scaffold/src/go/web/middleware"
)

const (
	ASSET_ROOT = "./assets"
)

func main() {
	env.Dump()

	// modelMuxやhtmlMuxには認証ミドルウェアを仕込む必要がある.

	modelMux := web.New()
    modelMux.Use(middleware.Authenticator)
	modelMux.Get("/model/order/", handler.AllOrdersHandler)
	modelMux.Get("/model/app-user/", handler.AllAppUsersHandler)

	staticMux := web.New()
	staticMux.Get("/css/*", handler.GetAssetsHandlerWithContentType("text/css", ASSET_ROOT))
	staticMux.Get("/js/*", handler.GetAssetsHandlerWithContentType("text/javascript", ASSET_ROOT))
	staticMux.Get("/*", http.FileServer(http.Dir(ASSET_ROOT)))

	htmlMux := web.New()
    htmlMux.Use(middleware.Authenticator)
	htmlMux.Get("/page/:page/", handler.GetHtmlHandler(ASSET_ROOT+"/html", ASSET_ROOT))
	htmlMux.Get("/", handler.GetHtmlPathHandler(ASSET_ROOT+"/html/index.html", ASSET_ROOT))

	publicMux := web.New()
    publicMux.Get("/login", handler.GetHtmlPathHandler(ASSET_ROOT+"/html/login.html", ASSET_ROOT))
	publicMux.Post("/model/authenticator", handler.AuthenticationHandler)
    publicMux.Get("/logout", handler.LogoutHandler)

	// 各MuxをURLに割り当てる
	// MuxでもURLが登場するので、冗長と言えば冗長.
	// 指定の順番は要注意.
	// 先に割り当てた方が先にマッチするので
	// 競合する指定は優先させたいMuxを先に記述する必要がある.
	defaultMux := goji.DefaultMux
	defaultMux.Handle("/", htmlMux)
	defaultMux.Handle("/model/authenticator", publicMux)
	defaultMux.Handle("/model/*", modelMux)
	defaultMux.Handle("/css/*", staticMux)
	defaultMux.Handle("/js/*", staticMux)
	defaultMux.Handle("/page/*", htmlMux)
    defaultMux.Handle("/login", publicMux)
    defaultMux.Handle("/logout", publicMux)
	defaultMux.Handle("/*", staticMux)

	goji.Serve()
}
