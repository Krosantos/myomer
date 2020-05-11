package elements

// abs -- go's default abs returns floats, and this is easier
func abs(i int) int {
	if i < 0 {
		return -1 * i
	}
	return i
}

// shareSign -- See if tow ints share a sign
func shareSign(a int, b int) {
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
	if abs(Δx) > abs(Δy) return abs(Δx)
	return abs(Δy)
}

// Path unwrapper

// A*, yeet yeet yeet

func (t *Tile) aStar(d *Tile) {

}
