package restapi

import (
	"net/http"
	"time"

	chi "github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/loqutus/O-1/pkg/types"
	"github.com/sirupsen/logrus"
)

func Start() {
	logrus.Println("restapi: starting")
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/*", GetFile)
	r.Post("/*", PostFileHandler)
	r.Delete("/*", DeleteFileHandler)

	rProbe := chi.NewRouter()
	rProbe.Get("/probe/ready", ReadyProbeHandler)

	go http.ListenAndServe(":"+types.Server.ListenPortProbe, rProbe)
	for !types.Server.Ready {
		time.Sleep(time.Second)
	}
	http.ListenAndServe(":"+types.Server.ListenPort, r)
}
