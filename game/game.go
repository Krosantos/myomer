package game

// Game -- An individual game in memory, with a board, units, and players
type Game struct {
	Board board
	units []unit
}

// getUnitsForPlayer -- Get all units for a given player
func (g *Game) getUnitsForPlayer(p int) []unit {
	result := []unit{}
	for _, u := range g.units {
		if u.team == p {
			result = append(result, u)
		}
	}
	return result
}

// getCorpses -- Get all dead units
func (g *Game) getCorpses() []unit {
	result := []unit{}
	for _, u := range g.units {
		if u.isDead == true {
			result = append(result, u)
		}
	}
	return result
}

// endRound -- End the round, triggering abilities, then checking for victory
func (g *Game) endRound() {
	for _, c := range g.getCorpses() {
		c.endTurn()
	}

	leftBase := g.Board.getLeftBase()
	rightBase := g.Board.getRightBase()

	if leftBase.unit != nil && leftBase.unit.team != 0 {
		print("Victory kinda")
	}
	if rightBase.unit != nil && rightBase.unit.team != 1 {
		print("Victory kinda")
	}
}
