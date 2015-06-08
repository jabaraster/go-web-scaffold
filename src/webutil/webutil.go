package webutil

import (
    "net/http"
    "../assets"
    "encoding/json"
)

func WriteTemplateResponse(w http.ResponseWriter, r *http.Request, templatePath string,  data interface{}) {
    tmpl, err := assets.ParseAsset(templatePath)
    if err != nil {
        http.NotFound(w, r)
        return
    }
    err = tmpl.Execute(w, data)
    if err != nil {
        panic(err)
    }
}

func WriteJsonResponse(w http.ResponseWriter, data interface{}) error {
    b, e := json.Marshal(data)
    if e != nil {
        return e
    }
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    _, e2 := w.Write(b)
    return e2
}
