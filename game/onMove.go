package game

type onMove func(*unit, *tile)

var onMoveRegistry map[string]onMove

func init() {
	onMoveRegistry = map[string]onMove{
		"grassy": func(u *unit, t *tile) {
			u.tile.terrain = plain
			t.terrain = plain
		},
	}
}
