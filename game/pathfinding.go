package game

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
func (t *tile) heuristic(d *tile) int {
	Δx := t.x - d.x
	Δy := t.y - d.y

	if shareSign(Δx, Δy) {
		return abs(Δx + Δy)
	}
	if abs(Δx) > abs(Δy) {
		return abs(Δx)
	}
	return abs(Δy)
}

// getPassable --Determine whether it's possiblt to move into a tile, based on terrain, occupants, and unit movetype
func getPassable(u *unit, from *tile, to *tile) bool {
	hasEnemy := to.unit != nil && to.unit.team != u.team
	isTall := abs(from.z-to.z) > 1
	isImpassable := to.terrain == water || to.terrain == void

	canPassEnemy := u.moveType == flying || u.moveType == teleport || u.moveType == infiltrate
	canPassTall := u.moveType == flying || u.moveType == teleport || u.moveType == climb
	canPassImpass := u.moveType == flying || u.moveType == teleport

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
func getCanEndOn(u *unit, t *tile) bool {
	isEmpty := t.unit == nil
	if t.terrain == void || t.terrain == water {
		return u.moveType == flying
	}
	return isEmpty
}

// getMovableTiles -- Get all tiles a unit can move to
func getMovableTiles(u *unit) map[*tile]bool {
	t := u.tile
	openList := map[*tile]bool{t: true}
	closedList := map[*tile]bool{}
	costDict := map[*tile]int{t: 0}
	result := map[*tile]bool{}

	for len(openList) > 0 {
		for ot := range openList {
			for _, neighbour := range ot.neighbours() {
				costToMove := costDict[ot] + 1
				hasMovesRemaining := costToMove <= u.speed
				isOnNoLists := closedList[neighbour] != true && openList[neighbour] != true
				// If it's not in the open or closed list, and you can move to it, add it to the open list, and the cost dictionary
				if getPassable(u, ot, neighbour) && isOnNoLists && hasMovesRemaining {
					openList[neighbour] = true
					costDict[neighbour] = costToMove
					if getCanEndOn(u, neighbour) {
						result[neighbour] = true
					}
					// If I'm already on the list, and this is a more efficient route, replace the cost
				} else if openList[neighbour] == true && costToMove < costDict[neighbour] {
					costDict[neighbour] = costToMove
				}
			}
			// We've looked at all the neighbours, consign this tile to the closed list.
			closedList[ot] = true
			delete(openList, ot)
		}
	}
	return result
}

// getMoveIsValid -- Return whether or not a given unit can move to a given tile.
func getMoveIsValid(u *unit, destination *tile) bool {
	t := u.tile
	openList := map[*tile]bool{t: true}
	closedList := map[*tile]bool{}
	costDict := map[*tile]int{t: 0}

	for len(openList) > 0 {
		for ot := range openList {
			for _, neighbour := range ot.neighbours() {
				costToMove := costDict[ot] + 1
				hasMovesRemaining := costToMove <= u.speed
				isOnNoLists := closedList[neighbour] != true && openList[neighbour] != true
				// If it's not in the open or closed list, and you can move to it, add it to the open list, and the cost dictionary
				if getPassable(u, ot, neighbour) && isOnNoLists && hasMovesRemaining {
					openList[neighbour] = true
					costDict[neighbour] = costToMove
					if getCanEndOn(u, neighbour) && neighbour == destination {
						return true
					}
					// If I'm already on the list, and this is a more efficient route, replace the cost
				} else if openList[neighbour] == true && costToMove < costDict[neighbour] {
					costDict[neighbour] = costToMove
				}
			}
			// We've looked at all the neighbours, consign this tile to the closed list.
			closedList[ot] = true
			delete(openList, ot)
		}
	}
	return false
}