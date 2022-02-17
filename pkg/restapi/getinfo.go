package server

import (
	"encoding/json"
	"net/http"

	"github.com/loqutus/O-1/pkg/types"
)

func GetInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(types.Node)
	return
}
