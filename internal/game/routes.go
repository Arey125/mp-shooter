package game

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

func (g *Game) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /{$}", g.index)
	mux.HandleFunc("/ws", g.ws)
}

func (g *Game) index(w http.ResponseWriter, r *http.Request) {
	indexPage().Render(context.Background(), w)
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func (g *Game) ws(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(conn.RemoteAddr(), "connected")
	client := &Client{hub: &g.hub, conn: conn, send: make(chan Message, 256)}
    client.hub.register <- client

    go client.readPump()
    go client.writePump()
}
