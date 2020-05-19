package elements

// Game -- An individual game in memory, with a board, units, and players
type Game struct {
	Board Board
	units []Unit
}

// GetUnitsForPlayer -- Get all units for a given player
func (g *Game) GetUnitsForPlayer(p int) []Unit {
	result := []Unit{}
	for _, u := range g.units {
		if u.Team == p {
			result = append(result, u)
		}
	}
	return result
}

// GetCorpses -- Get all dead units
func (g *Game) GetCorpses() []Unit {
	result := []Unit{}
	for _, u := range g.units {
		if u.IsDead == true {
			result = append(result, u)
		}
	}
	return result
}

// EndRound -- End the round, triggering abilities, then checking for victory
func (g *Game) EndRound() {
	for _, c := range g.GetCorpses() {
		c.EndTurn()
	}

	leftBase := g.Board.getLeftBase()
	rightBase := g.Board.getRightBase()

	if leftBase.Unit != nil && leftBase.Unit.Team != 0 {
		print("Victory kinda")
	}
	if rightBase.Unit != nil && rightBase.Unit.Team != 1 {
		print("Victory kinda")
	}
}
