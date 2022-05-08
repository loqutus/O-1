package server

import (
	"net/http"

	chi "github.com/go-chi/chi/v5"
	"github.com/loqutus/O-1/pkg/types"
)

// startProbe starts the probe.
func startProbe() {
	rProbe := chi.NewRouter()
	rProbe.Get("/probe/ready", ReadyProbeHandler)
	http.ListenAndServe(":"+types.Server.ListenPortProbe, rProbe)
}
