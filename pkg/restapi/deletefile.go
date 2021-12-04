package restapi

import "net/http"

func DeleteFile(w http.ResponseWriter, r *http.Request) {
	fileName := r.URL.Path
	fileInfo, err := api.Cli.Get(ctx, fileName)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
}
