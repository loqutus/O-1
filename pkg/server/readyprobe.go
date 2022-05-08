package server

import (
	"net/http"

	"github.com/loqutus/O-1/pkg/types"
)

// ReadyProbeHandler handles the ready probe requests.
func ReadyProbeHandler(w http.ResponseWriter, r *http.Request) {
	if types.Server.Ready {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusServiceUnavailable)
	}
}
