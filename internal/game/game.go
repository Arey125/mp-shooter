package game

type Player struct {
    X float64 `json:"x"`
    Y float64 `json:"y"`
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
    g.players[c] = &Player{}
}

func (g *Game) removePlayer(c *Client) {
    delete(g.players, c)
}

func (g *Game) updateState(c Command) map[string]any {
    commandType := c.message["type"].(string)
    if commandType == "move" {
        dx := c.message["dx"].(float64)
        dy := c.message["dy"].(float64)

        player := g.players[c.client]
        player.X += dx
        player.Y += dy
    } 

    state := make(map[string]any)
    for client, player := range g.players {
        state[client.conn.RemoteAddr().String()] = player
    }
    return state
}
