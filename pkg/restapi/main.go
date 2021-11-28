package httpserver

import (
	"net/http"

	"github.com/loqutus/O-1/pkg/restapi/controller"

	chi "github.com/go-chi/chi/v5"
)

func Start() {
	r := chi.NewRouter()

	c := controller.NewController()

	v1 := r.Group("/api/v1")
	{
		files := v1.Group("/accounts")
		{
			files.GET(":id", c.GetFile)
			files.GET("", c.ListFiles)
			files.POST("", c.AddFile)
			files.DELETE(":id", c.DeleteFile)
		}
		//...
	}

	http.ListenAndServe(":6969", r)
}
