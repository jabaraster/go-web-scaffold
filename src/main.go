package main

import (
    "regexp"
    "net/http"

    "github.com/zenazn/goji"
    "github.com/zenazn/goji/web/middleware"

    "./env"
    "./handler"
    "./assets"
)

func main() {
    env.Dump()

    // htmlページ
    goji.Get("", http.RedirectHandler("/", http.StatusSeeOther))
    goji.Get("/", assets.BasicLayoutHtmlHandler("html/index.html"))

    // 静的リソース
    goji.Get("/css/*", assets.ContentTypeHandler("text/css"))
    goji.Get("/js/*", assets.ContentTypeHandler("text/javascript"))
    goji.Get(regexp.MustCompile("/img/.*\\.jpg"), assets.ContentTypeHandler("image/jpeg"))
    goji.Get(regexp.MustCompile("/img/.*\\.png"), assets.ContentTypeHandler("image/png"))

    // API
    goji.Get("/api/json/index", handler.JsonHandler)

    // ミドルウェア
    goji.Use(middleware.Recoverer)
    goji.Use(middleware.NoCache)
    // ログイン済みチェックは、ミドルウェアで実装可能？

    goji.Serve()
}
