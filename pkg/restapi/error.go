package restapi

import (
	"log"
	"net/http"
)

func Error(err error, w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(err.Error()))
	log.Println(err.Error())
}
