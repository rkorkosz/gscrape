package app

import (
	"io"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func NewRouter(logOut io.Writer) http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/", index)
	c := handlers.CompressHandler(r)
	cors := handlers.CORS(handlers.AllowedOrigins([]string{"*"}))
	return handlers.LoggingHandler(logOut, cors(c))
}
