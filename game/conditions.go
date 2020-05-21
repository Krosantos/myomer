package game

// Condition -- A temporary unit status condition; buffs and debuffs
type Condition interface {
	onAdd(*unit)
	onRemove(*unit)
	id() string
	duration() int
	onTurnEnd(*unit)
}

// Poison -- Take damage each turn
type Poison struct {
	length       int
	strength     int
	turnsElapsed int
}

const poisonID string = "Poison"

func (p Poison) onAdd(u *unit)    {}
func (p Poison) onRemove(u *unit) {}
func (p Poison) id() string       { return poisonID }
func (p Poison) duration() int    { return p.length - p.turnsElapsed }
func (p Poison) onTurnEnd(u *unit) {
	u.takeDamage(nil, p.strength)
	p.turnsElapsed++
	if p.turnsElapsed >= p.length {
		p.onRemove(u)
	}
}
