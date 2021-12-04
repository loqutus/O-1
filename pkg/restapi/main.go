package restapi

import (
	"context"
	"net/http"

	chi "github.com/go-chi/chi/v5"
)

var api RestAPI

func Start(ctx *context.Context, cli *clientv3.Client, localDir string) {
	api.Ctx = ctx
	api.Cli = cli
	api.LocalDir = localDir

	r := chi.NewRouter()

	r.Get("", GetFile)
	r.Post("", PostFileHandler)
	r.Delete("", DeleteFile)

	http.ListenAndServe(":6969", r)
}
