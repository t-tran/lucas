package handlers

import (
    "os"
    "io/ioutil"
    "net/http"
)

func JsHandler(w http.ResponseWriter, r *http.Request) {
    if _, err := os.Stat("static/index.js"); err == nil {
        data, err := ioutil.ReadFile("static/index.js")
        if (err == nil) {
            w.Write([]byte(data))
        }
    }
}
