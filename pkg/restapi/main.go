package restapi

import (
	"net/http"

	chi "github.com/go-chi/chi/v5"
)

var api RestAPI

func Start() {

	r := chi.NewRouter()

	r.Get("", GetFile)
	r.Post("", PostFileHandler)
	r.Delete("", DeleteFile)

	http.ListenAndServe(":6969", r)
}
