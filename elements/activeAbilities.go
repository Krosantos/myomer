package elements

// ActiveAbility -- An activated ability, to be used by a unit
type ActiveAbility interface {
	id() string
	targetIsValid(*Unit, *Tile) bool
	activate(*Unit, *Tile)
}

var abilityRegistry = map[string]func(...interface{}) ActiveAbility{
	"0010123": buildFireball,
}

// fireball -- yeet me mama like a wagon wheel
type fireball struct {
	horse int
}

func buildFireball(args ...interface{}) ActiveAbility {
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
