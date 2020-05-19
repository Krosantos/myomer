package elements

// ActiveAbility -- An activated ability, to be used by a unit
type ActiveAbility interface {
	id() string
	targetIsValid(*Unit, *Tile) bool
	activate(*Unit, *Tile)
}

var abilityRegistry map[string]ActiveAbility

func init() {
	abilityRegistry = map[string]ActiveAbility{

		poisonCloudID: poisonCloud{},
	}
}

const poisonCloudID = "AB_PoisonCloud"

// poisonCloud -- yeet me mama like a wagon wheel
type poisonCloud struct{}

func (f poisonCloud) targetIsValid(unit *Unit, tile *Tile) bool {
	return true
}
func (f poisonCloud) activate(unit *Unit, tile *Tile) {
	if tile.Unit != nil {
		poison := Poison{3, 1, 0}
		tile.Unit.AddCondition(poison)
	}
}
func (f poisonCloud) id() string {
	return poisonCloudID
}
