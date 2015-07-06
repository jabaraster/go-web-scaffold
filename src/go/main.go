package main

import (
	"./env"
	"./web/handler"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
	"net/http"
    "./web/middleware"
)

const (
	ASSET_ROOT = "./assets"
)

func main() {
	env.Dump()

	// modelMuxやhtmlMuxには認証ミドルウェアを仕込む必要がある.

	resoureMux := web.New()
    resoureMux.Use(middleware.JsonAuthenticator)
	resoureMux.Get("/resource/order/", handler.AllOrdersHandler)
	resoureMux.Get("/resource/app-user/", handler.AllAppUsersHandler)

	staticMux := web.New()
	staticMux.Get("/css/*", handler.GetAssetsHandlerWithContentType("text/css", ASSET_ROOT))
	staticMux.Get("/js/*", handler.GetAssetsHandlerWithContentType("text/javascript", ASSET_ROOT))
	staticMux.Get("/*", http.FileServer(http.Dir(ASSET_ROOT)))

	htmlMux := web.New()
    htmlMux.Use(middleware.PageAuthenticator)
	htmlMux.Get("/page/:page/", handler.GetHtmlHandler(ASSET_ROOT+"/html", ASSET_ROOT))
	htmlMux.Get("/", handler.GetHtmlPathHandler(ASSET_ROOT+"/html/index.html", ASSET_ROOT))

	publicMux := web.New()
    publicMux.Get("/login", handler.GetHtmlPathHandler(ASSET_ROOT+"/html/login.html", ASSET_ROOT))
	publicMux.Post("/resource/authenticator", handler.AuthenticationHandler)
    publicMux.Get("/logout", handler.LogoutHandler)

	// 各MuxをURLに割り当てる
	// MuxでもURLが登場するので、冗長と言えば冗長.
	// 指定の順番は要注意.
	// 先に割り当てた方が先にマッチするので
	// 競合する指定は優先させたいMuxを先に記述する必要がある.
	defaultMux := goji.DefaultMux
	defaultMux.Handle("/", htmlMux)
	defaultMux.Handle("/resource/authenticator", publicMux)
	defaultMux.Handle("/resource/*", resoureMux)
	defaultMux.Handle("/css/*", staticMux)
	defaultMux.Handle("/js/*", staticMux)
	defaultMux.Handle("/page/*", htmlMux)
    defaultMux.Handle("/login", publicMux)
    defaultMux.Handle("/logout", publicMux)
	defaultMux.Handle("/*", staticMux)

	goji.Serve()
}
