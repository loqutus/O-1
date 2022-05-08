package restapi

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

// ErrorHandler writes error content to the response and logs the error.
func Error(err error, w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(err.Error()))
	logrus.Warningln(err.Error())
}
