package game

// tile -- A hexagonal map tile
type tile struct {
	x       int
	y       int
	z       int
	unit    *unit
	corpse  *unit
	terrain terrain
	board   *board
}
type terrain string

const (
	grass terrain = "grass"
	void  terrain = "void"
	water terrain = "water"
)

// board -- The gameboard map, which holds tiles
type board struct {
	tiles map[int]map[int]*tile
}

// get -- get a tile by coordinates
func (b board) get(x int, y int) *tile {
	row, exists := b.tiles[x]
	if exists == false {
		return nil
	}
	t, exists := row[y]
	if exists == false {
		return nil
	}
	return t
}

// set -- set a tile by coordinates
func (b board) set(t *tile) {
	x := t.x
	y := t.y
	if b.tiles == nil {
		b.tiles = make(map[int]map[int]*tile)
	}
	_, exists := b.tiles[x]
	if exists == false {
		b.tiles[x] = make(map[int]*tile)
	}
	b.tiles[x][y] = t
}

// These coordinates are currently magic numbers, apologies
func (b board) getLeftBase() *tile {
	return b.get(-7, -3)
}
func (b board) getRightBase() *tile {
	return b.get(7, 4)
}

// neighbours -- Get all neighbouring tiles
func (t tile) neighbours() []*tile {
	result := []*tile{}
	up := t.board.get(t.x, t.y+1)
	down := t.board.get(t.x, t.y-1)
	upRight := t.board.get(t.x+1, t.y+1)
	upLeft := t.board.get(t.x-1, t.y)
	downRight := t.board.get(t.x+1, t.y)
	downLeft := t.board.get(t.x-1, t.y-1)
	if up != nil {
		result = append(result, up)
	}
	if upRight != nil {
		result = append(result, upRight)
	}
	if downRight != nil {
		result = append(result, downRight)
	}
	if down != nil {
		result = append(result, down)
	}
	if downLeft != nil {
		result = append(result, downLeft)
	}
	if upLeft != nil {
		result = append(result, upLeft)
	}
	return result
}
