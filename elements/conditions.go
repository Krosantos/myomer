package elements

// Condition -- A temporary unit status condition; buffs and debuffs
type Condition interface {
	onAdd(*Unit)
	onRemove(*Unit)
	id() string
	duration() int
	onTurnEnd(*Unit)
}

// Poison -- Take damage each turn
type Poison struct {
	length       int
	strength     int
	turnsElapsed int
}

const poisonID string = "Poison"

func (p Poison) onAdd(unit *Unit)    {}
func (p Poison) onRemove(unit *Unit) {}
func (p Poison) id() string          { return poisonID }
func (p Poison) duration() int       { return p.length - p.turnsElapsed }
func (p Poison) onTurnEnd(unit *Unit) {
	unit.TakeDamage(nil, p.strength)
	p.turnsElapsed++
	if p.turnsElapsed >= p.length {
		p.onRemove(unit)
	}
}
