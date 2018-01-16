package handlers

import (
	"net/http"
	"github.com/t-tran/lucas/frontend"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(frontend.HOME_PAGE))
}
