package restapi

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

func Error(err error, w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(err.Error()))
	logrus.Warningln(err.Error())
}
