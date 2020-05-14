package elements

// Tile -- A hexagonal map tile
type Tile struct {
	X       int
	Y       int
	Z       int
	Unit    *Unit
	Corpse  *Unit
	Terrain terrain
	Board   *Board
}
type terrain string

const (
	grass terrain = "grass"
	void  terrain = "void"
	water terrain = "water"
)

// Board -- The gameboard map, which holds tiles
type Board struct {
	Tiles map[int]map[int]*Tile
}

// Get -- Get a tile by coordinates
func (t Board) Get(x int, y int) *Tile {
	row, exists := t.Tiles[x]
	if exists == false {
		return nil
	}
	tile, exists := row[y]
	if exists == false {
		return nil
	}
	return tile
}

// Set -- Set a tile by coordinates
func (t Board) Set(tile *Tile) {
	x := tile.X
	y := tile.Y
	if t.Tiles == nil {
		t.Tiles = make(map[int]map[int]*Tile)
	}
	_, exists := t.Tiles[x]
	if exists == false {
		t.Tiles[x] = make(map[int]*Tile)
	}
	t.Tiles[x][y] = tile
}

// Neighbours -- Get all neighbouring tiles
func (t Tile) Neighbours() []*Tile {
	result := []*Tile{}
	up := t.Board.Get(t.X, t.Y+1)
	down := t.Board.Get(t.X, t.Y-1)
	upRight := t.Board.Get(t.X+1, t.Y+1)
	upLeft := t.Board.Get(t.X-1, t.Y)
	downRight := t.Board.Get(t.X+1, t.Y)
	downLeft := t.Board.Get(t.X-1, t.Y-1)
	if up != nil {
		result = append(result, up)
	}
	if down != nil {
		result = append(result, down)
	}
	if upRight != nil {
		result = append(result, upRight)
	}
	if upLeft != nil {
		result = append(result, upLeft)
	}
	if downRight != nil {
		result = append(result, downRight)
	}
	if downLeft != nil {
		result = append(result, downLeft)
	}
	return result
}
