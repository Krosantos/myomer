package game

import (
	"encoding/json"
)

type unitTemplate struct {
	Name            string     `json:"name"`
	Cost            int        `json:"cost"`
	Color           color      `json:"color"`
	Strength        int        `json:"strength"`
	Health          int        `json:"health"`
	Speed           int        `json:"speed"`
	Moxie           int        `json:"moxie"`
	AttackRange     int        `json:"attackRange"`
	AttackType      targetType `json:"attackType"`
	MoveType        moveType   `json:"moveType"`
	OnAttack        []string   `json:"onAttack"`
	OnDie           []string   `json:"onDie"`
	OnKill          []string   `json:"onKill"`
	OnMove          []string   `json:"onMove"`
	OnStrike        []string   `json:"onStrike"`
	OnStruck        []string   `json:"onStruck"`
	OnTurnEnd       []string   `json:"onTurnEnd"`
	ActiveAbilities []string   `json:"activeAbilities"`
}

// buildUnit -- Given a JSON template, build a unit
func buildUnit(s string, team int) unit {
	var t unitTemplate
	err := json.Unmarshal([]byte(s), &t)
	if err != nil {
		panic(err)
	}
	result := unit{
		name:            t.Name,
		team:            team,
		cost:            t.Cost,
		color:           t.Color,
		strength:        t.Strength,
		health:          t.Health,
		speed:           t.Speed,
		moxie:           t.Moxie,
		attackRange:     t.AttackRange,
		moveType:        t.MoveType,
		onAttack:        t.OnAttack,
		onDie:           t.OnDie,
		onKill:          t.OnKill,
		onMove:          t.OnMove,
		onStrike:        t.OnStrike,
		onStruck:        t.OnStruck,
		onTurnEnd:       t.OnTurnEnd,
		activeAbilities: t.ActiveAbilities,
		conditions:      make(map[string]Condition),
	}

	return result
}
