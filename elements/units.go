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
	OnMove     map[string]func(*Unit, *Tile)
	OnAttack   map[string]func(*Unit, *Tile)
	OnDie      map[string]func(*Unit, *Unit)
	OnStrike   map[string]func(*Unit, *Unit, int) int
	OnStruck   map[string]func(*Unit, *Unit, int) int
	OnKill     map[string]func(*Unit, *Unit)
	OnRoundEnd map[string]func(*Unit)
	Conditions map[string]Condition
}

// Move -- Move a unit to a specified tile
func (unit Unit) Move(tile *Tile) {
	for _, fn := range unit.OnMove {
		fn(&unit, tile)
	}
}

// Attack -- Attack a specified tile
func (unit Unit) Attack(tile *Tile) {
	for _, fn := range unit.OnAttack {
		fn(&unit, tile)
	}
}

// Die -- Die with dignity and side effects
func (unit Unit) Die(killer *Unit) {
	for _, fn := range unit.OnDie {
		fn(&unit, killer)
	}
}

// Strike -- Calculate pre-mitigation damage, and go through side effects of an attack
func (unit Unit) Strike(target *Unit) int {
	initialDamage := unit.Strength
	for _, fn := range unit.OnStrike {
		initialDamage = fn(&unit, target, initialDamage)
	}
	return initialDamage
}

// Struck -- Calculate post-mitigation damage, go through side effects
func (unit Unit) Struck(attacker *Unit, damage int) {
	damageReceived := damage
	for _, fn := range unit.OnStruck {
		damageReceived = fn(&unit, attacker, damageReceived)
	}
	unit.TakeDamage(attacker, damageReceived)
}

// TakeDamage -- Actually receive damage, possibly die
func (unit Unit) TakeDamage(attacker *Unit, damage int) {
	unit.Damage += damage
	if unit.Damage >= unit.Health {
		unit.Tile.Corpse = &unit
		unit.Tile.Unit = nil
		unit.Die(attacker)
		attacker.Kill(&unit)
	}
}

// Kill -- Run when a unit kills another
func (unit Unit) Kill(victim *Unit) {
	for _, fn := range unit.OnKill {
		fn(&unit, victim)
	}
}

// EndRound -- Run on end of round for per-turn effects
func (unit Unit) EndRound() {
	for _, fn := range unit.OnRoundEnd {
		fn(&unit)
	}
	for _, condition := range unit.Conditions {
		condition.onRoundEnd(&unit)
	}
}

// AddCondition -- Add a condition by name
func (unit Unit) AddCondition(condition Condition) {
	unit.Conditions[condition.name()] = condition
}

// RemoveCondition -- Remove a condition by name
func (unit Unit) RemoveCondition(condition Condition) {
	delete(unit.Conditions, condition.name())
}

// MoveIsValid -- Determine whether a unit can legally move to a tile
func (unit Unit) MoveIsValid(tile *Tile) bool {
	return true
}

// AttackIsValid -- Determine whether a unit can legally attack a tile
func (unit Unit) AttackIsValid(tile *Tile) bool {
	return true
}
