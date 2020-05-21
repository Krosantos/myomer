package game

// unit -- A combat unit/board piece
type unit struct {
	name            string
	team            int
	cost            int
	color           color
	strength        int
	health          int
	damage          int
	speed           int
	moxie           int
	attackRange     int
	attackType      targetType
	moveType        moveType
	tile            *tile
	isDead          bool
	onAttack        []string
	onDie           []string
	onKill          []string
	onMove          []string
	onStrike        []string
	onStruck        []string
	onTurnEnd       []string
	activeAbilities []string
	conditions      map[string]Condition
}

// move -- move a unit to a specified tile
func (u unit) move(t *tile) {
	for _, ability := range u.onMove {
		fn := onMoveRegistry[ability]
		fn(&u, t)
	}
}

// attack -- attack a specified tile
func (u unit) attack(t *tile) {
	for _, ability := range u.onAttack {
		fn := onAttackRegistry[ability]
		fn(&u, t)
	}
}

// die -- die with dignity and side effects
func (u unit) die(killer *unit) {
	u.tile.unit = nil
	u.tile.corpse = &u
	u.isDead = true

	for _, ability := range u.onDie {
		fn := onDieRegistry[ability]
		fn(&u, killer)
	}
}

// strike -- Calculate pre-mitigation damage, and go through side effects of an attack
func (u unit) strike(target *unit) int {
	initialDamage := u.strength
	for _, ability := range u.onStrike {
		fn := onStrikeRegistry[ability]
		initialDamage = fn(&u, target, initialDamage)
	}
	return initialDamage
}

// struck -- Calculate post-mitigation damage, go through side effects
func (u unit) struck(attacker *unit, damage int) {
	damageReceived := damage
	for _, ability := range u.onStruck {
		fn := onStruckRegistry[ability]
		damageReceived = fn(&u, attacker, damageReceived)
	}
	u.takeDamage(attacker, damageReceived)
}

// healDamage -- Recover health
func (u unit) healDamage(damage int) {
	u.damage -= damage
	if u.damage < 0 {
		u.damage = 0
	}
}

// takeDamage -- Actually receive damage, possibly die
func (u unit) takeDamage(attacker *unit, damage int) {
	u.damage += damage
	if u.damage >= u.health {
		u.tile.corpse = &u
		u.tile.unit = nil
		u.die(attacker)
		attacker.kill(&u)
	}
}

// kill -- Run when a unit kills another
func (u unit) kill(victim *unit) {
	for _, ability := range u.onKill {
		fn := onKillRegistry[ability]
		fn(&u, victim)
	}
}

// endTurn -- Run on end of round for per-turn effects
func (u unit) endTurn() {
	for _, ability := range u.onTurnEnd {
		fn := onTurnEndRegistry[ability]
		fn(&u)
	}
	for _, condition := range u.conditions {
		condition.onTurnEnd(&u)
	}
}

// addCondition -- Add a condition by id
func (u unit) addCondition(condition Condition) {
	u.conditions[condition.id()] = condition
}

// removeCondition -- Remove a condition by id
func (u unit) removeCondition(condition Condition) {
	delete(u.conditions, condition.id())
}

// moveIsValid -- Determine whether a unit can legally move to a tile
func (u unit) moveIsValid(t *tile) bool {
	return getMoveIsValid(&u, t)
}

// attackIsValid -- Determine whether a unit can legally attack a tile
func (u unit) attackIsValid(t *tile) bool {
	return true
}
