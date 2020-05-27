package game

// ParseCommand -- Parse incoming commands, apply them to the game state, return a list of instructions to send out to players
func (g Game) ParseCommand(cmd interface{}) {
	switch c := cmd.(type) {
	case moveCommand:
		u, ok := g.units[c.Unit]
		if !ok {
			break
		}
		t := g.board.get(c.Tile.X, c.Tile.Y)
		if u.moveIsValid(t) {
			u.move(t)
		}
		break
	case abilityCommand:
	case attackCommand:
	case endTurnCommand:
	case forfeitCommand:
	}
}
