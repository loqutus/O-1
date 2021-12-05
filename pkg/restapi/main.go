package restapi

import (
	"net/http"

	chi "github.com/go-chi/chi/v5"
)

func Start() {

	r := chi.NewRouter()

	r.Get("", GetFile)
	r.Post("", PostFileHandler)

	http.ListenAndServe(":6969", r)
}
