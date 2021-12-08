package restapi

import (
	"net/http"

	chi "github.com/go-chi/chi/v5"
	"github.com/loqutus/O-1/pkg/types"
)

func Start() {

	r := chi.NewRouter()

	r.Get("/", GetFile)
	r.Post("/", PostFileHandler)

	http.ListenAndServe(":"+types.Server.ListenPort, r)
}
