package elements

// Tile -- A hexagonal map tile
type Tile struct {
	X       int
	Y       int
	Z       int
	Unit    *Unit
	Corpse  *Unit
	Terrain terrain
	Map     *TileMap
}
type terrain string

const (
	grass terrain = "grass"
	void  terrain = "void"
	water terrain = "water"
)

// TileMap -- The gameboard map, which holds tiles
type TileMap struct {
	tiles map[int]map[int]*Tile
}

// Get -- Get a tile by coordinates
func (t TileMap) Get(x int, y int) *Tile {
	row, exists := t.tiles[x]
	if exists == false {
		return nil
	}
	tile, exists := row[y]
	if exists == false {
		return nil
	}
	return tile
}

// neighbours -- Get all neighbouring tiles
func (t Tile) neighbours() []*Tile {
	result := []*Tile{}
	up := t.Map.Get(t.X, t.Y+1)
	down := t.Map.Get(t.X, t.Y-1)
	upRight := t.Map.Get(t.X+1, t.Y+1)
	upLeft := t.Map.Get(t.X-1, t.Y)
	downRight := t.Map.Get(t.X+1, t.Y)
	downLeft := t.Map.Get(t.X-1, t.Y-1)
	if up != nil {
		append(result, up)
	}
	if down != nil {
		append(result, down)
	}
	if upRight != nil {
		append(result, upRight)
	}
	if upLeft != nil {
		append(result, upLeft)
	}
	if downRight != nil {
		append(result, downRight)
	}
	if downLeft != nil {
		append(result, downLeft)
	}
	return result
}
