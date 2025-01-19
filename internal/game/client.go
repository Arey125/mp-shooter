package game

import (
	"fmt"

	"github.com/gorilla/websocket"
)

type Client struct {
	hub  *Hub
	conn *websocket.Conn
	send chan Message
}

func (c *Client) readPump() {
    fmt.Println("client readPump")
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()
	for {
		command := make(map[string]any)
		err := c.conn.ReadJSON(&command)
        fmt.Println("command: ", command)
		if err != nil {
            fmt.Printf("error: %v\n", err)
			break
		}
        fmt.Println(command)
		c.hub.commands <- Command{client: c, message: command}
	}
}

func (c *Client) writePump() {
    fmt.Println("client writePump")
	for {
		message, ok := <-c.send
		if !ok {
			c.conn.WriteMessage(websocket.CloseMessage, []byte{})
			return
		}
        err := c.conn.WriteJSON(message)
        if err != nil {
            fmt.Println(err)
            return
        }
	}
}
