package game

type onAttack func(*unit, *tile)

var onAttackRegistry map[string]onAttack

func init() {
	onAttackRegistry = map[string]onAttack{
		"furyCutter": func(u *unit, t *tile) {
			u.strength++
		},
	}
}
