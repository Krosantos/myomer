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
		b.tiles = map[int]map[int]*tile{}
	}
	_, exists := b.tiles[x]
	if exists == false {
		b.tiles[x] = map[int]*tile{}
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

// getDefaultBoard -- Get a default flat board
func getDefaultBoard() board {
	result := board{
		tiles: map[int]map[int]*tile{},
	}
	for _, coord := range defaultBoardCoords {
		t := tile{
			x:       coord.X,
			y:       coord.Y,
			z:       0,
			unit:    nil,
			corpse:  nil,
			terrain: grass,
			board:   &result,
		}
		result.set(&t)
	}
	return result
}

// Coord -- X and Y coordinates to refer to a tile.
type Coord struct {
	X int `json:"x"`
	Y int `json:"y"`
}

// defaultBoardCoords -- The coordinates that make up a ~13x7 game board
var defaultBoardCoords []Coord = []Coord{
	{-7, -3},
	{-6, 0},
	{-6, -1},
	{-6, -2},
	{-6, -3},
	{-6, -4},
	{-6, -5},
	{-5, 1},
	{-5, 0},
	{-5, -1},
	{-5, -2},
	{-5, -3},
	{-5, -4},
	{-5, -5},
	{-4, 1},
	{-4, 0},
	{-4, -1},
	{-4, -2},
	{-4, -3},
	{-4, -4},
	{-3, 2},
	{-3, 1},
	{-3, 0},
	{-3, -1},
	{-3, -2},
	{-3, -3},
	{-3, -4},
	{-2, 2},
	{-2, 1},
	{-2, 0},
	{-2, -1},
	{-2, -2},
	{-2, -3},
	{-1, 3},
	{-1, 2},
	{-1, 1},
	{-1, 0},
	{-1, -1},
	{-1, -2},
	{-1, -3},
	{0, 3},
	{0, 2},
	{0, 1},
	{0, 0},
	{0, -1},
	{0, -2},
	{1, 4},
	{1, 3},
	{1, 2},
	{1, 1},
	{1, 0},
	{1, -1},
	{1, -2},
	{2, 4},
	{2, 3},
	{2, 2},
	{2, 1},
	{2, 0},
	{2, -1},
	{3, 5},
	{3, 4},
	{3, 3},
	{3, 2},
	{3, 1},
	{3, 0},
	{3, -1},
	{4, 5},
	{4, 4},
	{4, 3},
	{4, 2},
	{4, 1},
	{4, 0},
	{5, 6},
	{5, 5},
	{5, 4},
	{5, 3},
	{5, 2},
	{5, 1},
	{5, 0},
	{6, 6},
	{6, 5},
	{6, 4},
	{6, 3},
	{6, 2},
	{6, 1},
	{7, 4},
}
