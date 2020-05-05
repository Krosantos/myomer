package elements

// Tile -- A map tile
type Tile struct {
	X      int
	Y      int
	Z      int
	Unit   *Unit
	Corpse *Unit
}
