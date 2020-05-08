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
