package app

import (
	"log"
	"net/http"
	"net/url"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func index(w http.ResponseWriter, r *http.Request) {
	site := r.URL.Query().Get("site")
	parsed, err := url.Parse(site)
	if err != nil {
		panic(err)
	}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	out := Parse(parsed)
	for link := range out {
		err = conn.WriteJSON(link)
		if err != nil {
			log.Print(err)
		}
	}
	err = conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(1000, "woops"))
	if err != nil {
		log.Println(err)
	}
}
