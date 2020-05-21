package game

// ActiveAbility -- An activated ability, to be used by a unit
type ActiveAbility interface {
	id() string
	targetIsValid(*unit, *tile) bool
	activate(*unit, *tile)
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

func (f poisonCloud) targetIsValid(u *unit, t *tile) bool {
	return true
}
func (f poisonCloud) activate(u *unit, t *tile) {
	if t.unit != nil {
		poison := Poison{3, 1, 0}
		t.unit.addCondition(poison)
	}
}
func (f poisonCloud) id() string {
	return poisonCloudID
}
