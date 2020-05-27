package game

// unit -- A combat unit/board piece
type unit struct {
	id              string
	name            string
	templateID      string
	team            int
	cost            int
	color           string
	strength        int
	health          int
	damage          int
	speed           int
	moxie           int
	attackRange     int
	attackType      string
	moveType        string
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
	game            *Game
}

// move -- move a unit to a specified tile
func (u unit) move(t *tile) {
	mi := moveInstruction{
		Action: action.MOVE,
		Unit:   u.id,
		Tile:   Coord{t.x, t.y},
	}
	u.game.Out <- mi
	for _, ability := range u.onMove {
		fn := onMoveRegistry[ability]

		ai := abilityInstruction{
			Action:    action.ABILITY,
			AbilityID: ability,
			Unit:      u.id,
			Targets:   []Coord{{t.x, t.y}},
		}
		u.game.Out <- ai
		fn(&u, t)
	}
}

// attack -- attack a specified tile
func (u unit) attack(t *tile) {
	ati := attackInstruction{
		Action: action.ATTACK,
		Unit:   u.id,
		Target: Coord{t.x, t.y},
	}
	u.game.Out <- ati
	for _, ability := range u.onAttack {
		fn := onAttackRegistry[ability]
		ani := abilityInstruction{
			Action:    action.ABILITY,
			AbilityID: ability,
			Unit:      u.id,
			Targets:   []Coord{{t.x, t.y}},
		}
		u.game.Out <- ani
		fn(&u, t)
	}
	for _, target := range u.getTargets(t) {
		damage := u.strike(target)
		target.struck(&u, damage)
	}
}

// die -- die with dignity and side effects
func (u unit) die(killer *unit) {
	u.tile.unit = nil
	u.tile.corpse = &u
	u.isDead = true

	di := dieInstruction{
		Action: action.DIE,
		Unit:   u.id,
	}
	u.game.Out <- di

	for _, ability := range u.onDie {
		fn := onDieRegistry[ability]
		ai := abilityInstruction{
			Action:    action.ABILITY,
			AbilityID: ability,
			Unit:      u.id,
			Targets:   []Coord{{u.tile.x, u.tile.y}},
		}
		u.game.Out <- ai
		fn(&u, killer)
	}
}

// strike -- Calculate pre-mitigation damage, and go through side effects of an attack
func (u unit) strike(target *unit) int {
	initialDamage := u.strength
	for _, ability := range u.onStrike {
		fn := onStrikeRegistry[ability]
		ai := abilityInstruction{
			Action:    action.ABILITY,
			AbilityID: ability,
			Unit:      u.id,
			Targets:   []Coord{{target.tile.x, target.tile.y}},
		}
		u.game.Out <- ai
		initialDamage = fn(&u, target, initialDamage)
	}
	return initialDamage
}

// struck -- Calculate post-mitigation damage, go through side effects
func (u unit) struck(attacker *unit, damage int) {
	damageReceived := damage
	for _, ability := range u.onStruck {
		fn := onStruckRegistry[ability]
		ai := abilityInstruction{
			Action:    action.ABILITY,
			AbilityID: ability,
			Unit:      u.id,
			Targets:   []Coord{{attacker.tile.x, attacker.tile.y}},
		}
		u.game.Out <- ai
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
	i := healInstruction{
		Action: action.HEAL,
		Unit:   u.id,
		Amount: damage,
	}
	u.game.Out <- i
}

// takeDamage -- Actually receive damage, possibly die
func (u unit) takeDamage(attacker *unit, damage int) {
	u.damage += damage
	if u.damage >= u.health {
		u.die(attacker)
		attacker.kill(&u)
	}
	i := damageInstruction{
		Action: action.DAMAGE,
		Unit:   u.id,
		Amount: damage,
	}
	u.game.Out <- i
}

// kill -- Run when a unit kills another
func (u unit) kill(victim *unit) {
	for _, ability := range u.onKill {
		fn := onKillRegistry[ability]
		ai := abilityInstruction{
			Action:    action.ABILITY,
			AbilityID: ability,
			Unit:      u.id,
			Targets:   []Coord{{victim.tile.x, victim.tile.y}},
		}
		u.game.Out <- ai
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
	i := addConditionInstruction{
		Action:   action.ADDCONDITION,
		Unit:     u.id,
		StatusID: condition.id(),
		Duration: condition.duration(),
	}
	u.game.Out <- i
}

// removeCondition -- Remove a condition by id
func (u unit) removeCondition(condition Condition) {
	delete(u.conditions, condition.id())
	i := removeConditionInstruction{
		Action:   action.REMOVECONDITION,
		Unit:     u.id,
		StatusID: condition.id(),
	}
	u.game.Out <- i
}

// moveIsValid -- Determine whether a unit can legally move to a tile
func (u unit) moveIsValid(t *tile) bool {
	return getMoveIsValid(&u, t)
}

// attackIsValid -- Determine whether a unit can legally attack a tile
func (u unit) attackIsValid(t *tile) bool {
	// TODO: Actually check
	return true
}

//  getTargets -- Determine who will be struck by an ability or attack
func (u unit) getTargets(t *tile) []*unit {
	// TODO: Actually look at attack type and attack range
	if t.unit != nil {
		return []*unit{t.unit}
	}
	return []*unit{}
}
