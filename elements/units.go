package elements

// Unit -- A combat unit/board piece
type Unit struct {
	Name       string
	Cost       int
	Strength   int
	Health     int
	Damage     int
	AtkRange   int
	Speed      int
	Moxie      int
	MoveType   moveType
	Tile       *Tile
	onMove     []func(Unit, *Tile)
	onAttack   []func(Unit, *Tile)
	onDie      []func(Unit, *Unit)
	onStrike   []func(Unit, *Unit, int) int
	onStruck   []func(Unit, *Unit, int) int
	onKill     []func(Unit, *Unit)
	onRoundEnd []func(Unit)
	Conditions []Condition
}

// Move -- Move a unit to a specified tile
func (unit Unit) Move(tile *Tile) {
	for _, fn := range unit.onMove {
		fn(unit, tile)
	}
}

// Attack -- Attack a specified tile
func (unit Unit) Attack(tile *Tile) {
	for _, fn := range unit.onAttack {
		fn(unit, tile)
	}
}

// Die -- Die with dignity and side effects
func (unit Unit) Die(killer *Unit) {
	for _, fn := range unit.onDie {
		fn(unit, killer)
	}
}

// Strike -- Calculate pre-mitigation damage, and go through side effects of an attack
func (unit Unit) Strike(target *Unit) int {
	initialDamage := unit.Strength
	for _, fn := range unit.onStrike {
		initialDamage = fn(unit, target, initialDamage)
	}
	return initialDamage
}

// Struck -- Calculate and receive post-mitigation damage, go through side effects, possibly die.
func (unit Unit) Struck(attacker *Unit, damage int) {
	damageReceived := damage
	for _, fn := range unit.onStruck {
		damageReceived = fn(unit, attacker, damageReceived)
	}
	unit.Damage += damageReceived
	if unit.Damage >= unit.Health {
		unit.Tile.Corpse = &unit
		unit.Tile.Unit = nil
		unit.Die(attacker)
		attacker.Kill(&unit)
	}
}

// Kill -- Run when a unit kills another
func (unit Unit) Kill(victim *Unit) {
	for _, fn := range unit.onKill {
		fn(unit, victim)
	}
}

// EndRound -- Run on end of round for per-turn effects
func (unit Unit) EndRound() {
	for _, fn := range unit.onRoundEnd {
		fn(unit)
	}
}
