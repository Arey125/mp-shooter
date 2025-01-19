package game

import "fmt"

type Message map[string]any

type Command struct {
    client *Client
    message Message
}

type Hub struct {
	clients    map[*Client]bool
	commands  chan Command
	register   chan *Client
	unregister chan *Client
}

func newHub() Hub {
	return Hub{
		clients:    make(map[*Client]bool),
		commands:  make(chan Command, 256),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
}

func (h *Hub) run(g *Game) {
	for {
		select {
		case client := <-h.register:
            fmt.Println(client.conn.RemoteAddr(), "registered")
			h.clients[client] = true
            g.addPlayer(client)
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
                g.removePlayer(client)
				close(client.send)
			}
		case command := <-h.commands:
            state := g.updateState(command)
			for client := range h.clients {
				select {
                case client.send <- Message(state):
                    break
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}
