package restapi

import (
	"net/http"

	"github.com/loqutus/O-1/pkg/types"
)

func ReadyProbeHandler(w http.ResponseWriter, r *http.Request) {
	if types.Server.Ready {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusServiceUnavailable)
	}
	w.WriteHeader(http.StatusOK)
}
