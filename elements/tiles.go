package elements

// Tile -- A map tile
type Tile struct {
	X       int
	Y       int
	Z       int
	Unit    *Unit
	Corpse  *Unit
	Terrain terrain
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
	return t.tiles[x][y]
}
