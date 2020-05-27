package game

import "encoding/json"

// Game -- An individual game in memory, with a board, units, and players
type Game struct {
	board      *board           // The game board
	units      map[string]unit  // All the units
	activeUnit *unit            // The unit whose turn it is
	In         chan Command     // The game takes
	Out        chan Instruction // The game broadcasts instructions here
	End        chan bool        // The game makes a single call here when it's over
}

// BuildGame -- Create a new game state with the provided armies
func BuildGame() *Game {
	b := getDefaultBoard()
	g := &Game{
		board:      &b,
		units:      make(map[string]unit),
		activeUnit: nil,
		In:         make(chan Command),
		Out:        make(chan Instruction),
		End:        make(chan bool),
	}
	return g
}

// PopulateArmy -- Given a JSON template and a board, fill it with an army
func (g Game) PopulateArmy(s string, team int) {
	var a army
	err := json.Unmarshal([]byte(s), &a)
	if err != nil {
		panic(err)
	}
	for pos, templateID := range a.Units {
		tile := getUnitTile(pos, team, g.board)
		unit := g.buildUnit(templateID, team, tile)
		g.units[unit.id] = unit
	}
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

	leftBase := g.board.getLeftBase()
	rightBase := g.board.getRightBase()

	if leftBase.unit != nil && leftBase.unit.team != 0 {
		print("Victory kinda")
	}
	if rightBase.unit != nil && rightBase.unit.team != 1 {
		print("Victory kinda")
	}
}
