package elements

// Ability -- An activated ability, to be used by a unit
type Ability interface {
	id() string
	targetIsValid(*Unit, *Tile) bool
	activate(*Unit, *Tile)
}

var abilityRegistry = make(map[string]func(...interface{}) Ability)

func init() {
	abilityRegistry["0010123"] = buildFireball
}

// fireball -- yeet me mama like a wagon wheel
type fireball struct {
	horse int
}

func buildFireball(args ...interface{}) Ability {
	lemon := args[0]
	strength, ok := lemon.(int)
	if ok == false {
		println("FAILED FAILED FAILED")
	}
	return fireball{strength}
}
func (f fireball) targetIsValid(unit *Unit, tile *Tile) bool {
	return true
}
func (f fireball) activate(unit *Unit, tile *Tile) {
	print("yeet")
	if tile.Unit != nil {
		tile.Unit.Damage -= 3
		poison := Poison{3, 1, 0}
		tile.Unit.AddCondition(poison)
	}
}
func (f fireball) id() string {
	return "0010123"
}
