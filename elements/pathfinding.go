package elements

// abs -- go's default abs returns floats, and this is easier
func abs(i int) int {
	if i < 0 {
		return -1 * i
	}
	return i
}

// shareSign -- See if two ints share a sign
func shareSign(a int, b int) bool {
	if a < 0 {
		return b < 0
	}
	return b >= 0
}

// Manhattan distance heuristic
func (t *Tile) heuristic(d *Tile) int {
	Δx := t.X - d.X
	Δy := t.Y - d.Y

	if shareSign(Δx, Δy) {
		return abs(Δx + Δy)
	}
	if abs(Δx) > abs(Δy) {
		return abs(Δx)
	}
	return abs(Δy)
}

// getPassable --Determine whether it's possiblt to move into a tile, based on terrain, occupants, and unit movetype
func getPassable(u *Unit, from *Tile, to *Tile) bool {
	hasEnemy := to.Unit != nil && to.Unit.Team != u.Team
	isTall := abs(from.Z-to.Z) > 1
	isImpassable := to.Terrain == water || to.Terrain == void

	canPassEnemy := u.MoveType == flying || u.MoveType == teleport || u.MoveType == infiltrate
	canPassTall := u.MoveType == flying || u.MoveType == teleport || u.MoveType == climb
	canPassImpass := u.MoveType == flying || u.MoveType == teleport

	if hasEnemy && !canPassEnemy {
		return false
	}
	if isTall && !canPassTall {
		return false
	}
	if isImpassable && !canPassImpass {
		return false
	}

	return true
}

// getCanEndOn -- Determine whether a given unit can legally end its turn on a given tile
func getCanEndOn(u *Unit, t *Tile) bool {
	isEmpty := t.Unit == nil
	if t.Terrain == void || t.Terrain == water {
		return u.MoveType == flying
	}
	return isEmpty
}

// Path unwrapper

// aStar -- A*, yeet yeet yeet
func (t *Tile) aStar(u *Unit, d *Tile) {

}

// getMovableTiles -- Get all tiles a unit can move to
func getMovableTiles(u *Unit) []*Tile {
	t := u.Tile
	openList := map[*Tile]bool{t: true}
	closedList := map[*Tile]bool{}
	costDict := map[*Tile]int{t: 0}
	result := []*Tile{}

	for len(openList) > 0 {
		for tile := range openList {
			for _, neighbour := range tile.neighbours() {
				costToMove := costDict[tile] + 1
				hasMovesRemaining := costToMove <= u.Speed
				isOnNoLists := closedList[neighbour] != true && openList[neighbour] != true
				// If it's not in the open or closed list, and you can move to it, add it to the open list, and the cost dictionary
				if getPassable(u, tile, neighbour) && isOnNoLists && hasMovesRemaining {
					openList[neighbour] = true
					costDict[neighbour] = costToMove
					if getCanEndOn(u, neighbour) {
						result = append(result, neighbour)
					}
					// If I'm already on the list, and this is a more efficient route, replace the cost
				} else if openList[neighbour] == true && costToMove < costDict[neighbour] {
					costDict[neighbour] = costToMove
				}
			}
			// We've looked at all the neighbours, consign this tile to the closed list.
			closedList[tile] = true
			delete(openList, tile)
		}
	}
	return result
}
