package handler

import (
    "net/http"

    "../webutil"
)

func JsonHandler(w http.ResponseWriter, r *http.Request) {
    o := [](map[string]interface{}) {
        {
            "code": "001",
            "name" : "jabara",
        },
        {
            "code": "002",
            "name" : "jabaraster",
        },
    }
    webutil.WriteJsonResponse(w, o)
}
