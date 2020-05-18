package elements

type onAttack func(*Unit, *Tile)

var onAttackRegistry map[string]onAttack

func init() {
	onAttackRegistry = map[string]onAttack{
		"furyCutter": func(u *Unit, t *Tile) {
			u.Strength++
		},
	}
}
