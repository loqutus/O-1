package httpserver

import (
	"net/http"

	chi "github.com/go-chi/chi/v5"
)

func Start() {
	r := chi.NewRouter()

	r.Get(":id", GetFile)
	r.Get("", ListFiles)
	r.Post("", AddFile)
	r.Delete(":id", DeleteFile)

	http.ListenAndServe(":6969", r)
}
