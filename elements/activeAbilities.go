package elements

// ActiveAbility -- An activated ability, to be used by a unit
type ActiveAbility interface {
	id() string
	targetIsValid(*Unit, *Tile) bool
	activate(*Unit, *Tile)
}

var abilityRegistry = map[string]func(...interface{}) ActiveAbility{
	poisonCloudID: buildPoisonCloud,
}

const poisonCloudID = "AB_PoisonCloud"

// poisonCloud -- yeet me mama like a wagon wheel
type poisonCloud struct {
	strength int
}

func buildPoisonCloud(args ...interface{}) ActiveAbility {
	lemon := args[0]
	strength, ok := lemon.(int)
	if ok == false {
		println("FAILED FAILED FAILED")
	}
	return poisonCloud{strength}
}
func (f poisonCloud) targetIsValid(unit *Unit, tile *Tile) bool {
	return true
}
func (f poisonCloud) activate(unit *Unit, tile *Tile) {
	if tile.Unit != nil {
		tile.Unit.Damage -= f.strength
		poison := Poison{3, 1, 0}
		tile.Unit.AddCondition(poison)
	}
}
func (f poisonCloud) id() string {
	return poisonCloudID
}
