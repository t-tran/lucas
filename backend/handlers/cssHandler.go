package handlers

import (
    "os"
    "io/ioutil"
    "net/http"
)

func CssHandler(w http.ResponseWriter, r *http.Request) {
    if _, err := os.Stat("static/index.css"); err == nil {
        data, err := ioutil.ReadFile("static/index.css")
        if (err == nil) {
            w.Write([]byte(data))
        }
    }
}
