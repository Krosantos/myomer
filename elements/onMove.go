package elements

type onMove func(*Unit, *Tile)

var onMoveRegistry map[string]onMove

func init() {
	onMoveRegistry = map[string]onMove{
		"grassy": func(u *Unit, t *Tile) {
			u.Tile.Terrain = grass
			t.Terrain = grass
		},
	}
}
