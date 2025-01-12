package game

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

func (g Game) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /{$}", index)
	mux.HandleFunc("/ws", g.ws)
}

func index(w http.ResponseWriter, r *http.Request) {
	indexPage().Render(context.Background(), w)
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func (g Game) ws(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        fmt.Println(err)
        return
    }
    defer conn.Close()

    fmt.Println(conn.RemoteAddr(), "connected")
    conn.WriteJSON(map[string]any{"test": "test"})
    message := make(map[string]any)
    conn.ReadJSON(&message)
    fmt.Println("client answered", message)
}
