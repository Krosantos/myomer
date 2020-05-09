package elements

type onAttack func(*Unit, *Tile)

var onAttackRegistry map[string]onAttack = map[string]onAttack{}
