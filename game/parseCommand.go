package game

// ParseCommand -- Parse incoming commands, apply them to the game state, return a list of instructions to send out to players
func (g Game) ParseCommand(cmd Command) {
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
		//TODO: Implement any of this
	case attackCommand:
		u, ok := g.units[c.Unit]
		if !ok {
			break
		}
		t := g.board.get(c.Target.X, c.Target.Y)
		if u.attackIsValid(t) {
			u.attack(t)
		}
		break
	case endTurnCommand:
	case forfeitCommand:
		egi := endGameInstruction{
			Action: action.ENDGAME,
			Winner: 0,
		}
		g.Out <- egi
		break
	default:
	}
}
