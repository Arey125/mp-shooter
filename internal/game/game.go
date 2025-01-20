package game

type Position struct {
    X float64 `json:"x"`
    Y float64 `json:"y"`
}

type Player struct {
    position Position
    keys map[string]bool
}

type Game struct {
    hub Hub
    players map[*Client]*Player
}

func NewGame() *Game {
    g := &Game{
        hub: newHub(),
        players: make(map[*Client]*Player),
    }
    go g.hub.run(g)
    return g
}

func (g *Game) addPlayer(c *Client) {
    g.players[c] = &Player{keys: make(map[string]bool)}
}

func (g *Game) removePlayer(c *Client) {
    delete(g.players, c)
}

func (g *Game) updateState(c Command) {
    commandType := c.message["type"].(string)
    if commandType == "keyDown" {
        g.players[c.client].keys[c.message["key"].(string)] = true
        return;
    } 
    if commandType == "keyUp" {
        g.players[c.client].keys[c.message["key"].(string)] = false
        return;
    }
}

func (g *Game) getState() map[string]any {
    const SPEED = 5
    for i := range g.players {
        keys := g.players[i].keys
        if keys["W"] {
            g.players[i].position.Y += SPEED
        }
        if keys["A"] {
            g.players[i].position.X -= SPEED
        }
        if keys["S"] {
            g.players[i].position.Y -= SPEED
        }
        if keys["D"] {
            g.players[i].position.X += SPEED
        }
    }

    state := make(map[string]any)
    for client, player := range g.players {
        state[client.conn.RemoteAddr().String()] = player.position
    }
    return state
}
