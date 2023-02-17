package controller

import (
	"embed"
	"fmt"
	"html/template"
	"net/http"
	"os"

	"github.com/artamananda/artanymous/app/repository"
)

type API struct {
	messageRepo repository.MessageRepo
	embed       embed.FS
	mux         *http.ServeMux
}

func (api *API) BaseViewPath() *template.Template {
	var tmpl = template.Must(template.ParseFS(api.embed, "app/view/*"))
	return tmpl
}

func NewAPI(messageRepo repository.MessageRepo, embed embed.FS) API {
	mux := http.NewServeMux()
	api := API{
		messageRepo,
		embed,
		mux,
	}

	mux.HandleFunc("/", api.IndexPage)
	mux.HandleFunc("/icon", api.Favicon)

	mux.Handle("/api/message/add", http.HandlerFunc(api.AddMessage))
	mux.Handle("/api/message/read", http.HandlerFunc(api.ReadMessage))
	mux.Handle("/api/message/delete", http.HandlerFunc(api.DeleteMessage))
	mux.Handle("/api/message/reset", http.HandlerFunc(api.ResetMessage))

	return api
}

func (api *API) Handler() *http.ServeMux {
	return api.mux
}

func (api *API) Start() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "9990"

	}

	fmt.Println("starting web server at http://localhost:9990")
	http.ListenAndServe(":"+port, api.Handler())
}
