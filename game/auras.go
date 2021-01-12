package game

// Aura -- An area of effect, centered around a unit, which moves with them
type Aura interface {
	id() string
	source() *unit
	remove()
	move(*tile)
	onEnter(*unit, *tile)
	onLeave(*unit, *tile)
}

// PoisonCloud -- Poison units who enter
type PoisonCloud struct {
	unit *unit
}

func (p PoisonCloud) id() string { return "PoisonCloud" }
func source(p PoisonCloud) *unit { return p.unit }
func remove()                    { println("butt") }
