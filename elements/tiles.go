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
	Tiles map[int]map[int]*Tile
}

// Get -- Get a tile by coordinates
func (t TileMap) Get(x int, y int) *Tile {
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
func (t TileMap) Set(tile *Tile) {
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
	up := t.Map.Get(t.X, t.Y+1)
	down := t.Map.Get(t.X, t.Y-1)
	upRight := t.Map.Get(t.X+1, t.Y+1)
	upLeft := t.Map.Get(t.X-1, t.Y)
	downRight := t.Map.Get(t.X+1, t.Y)
	downLeft := t.Map.Get(t.X-1, t.Y-1)
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
