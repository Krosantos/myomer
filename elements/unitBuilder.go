package elements

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

// BuildUnit -- Given a JSON template, build a unit
func BuildUnit(s string, team int) Unit {
	var t unitTemplate
	err := json.Unmarshal([]byte(s), &t)
	if err != nil {
		panic(err)
	}
	unit := Unit{
		Name:            t.Name,
		Team:            team,
		Cost:            t.Cost,
		Color:           t.Color,
		Strength:        t.Strength,
		Health:          t.Health,
		Speed:           t.Speed,
		Moxie:           t.Moxie,
		AttackRange:     t.AttackRange,
		MoveType:        t.MoveType,
		OnAttack:        t.OnAttack,
		OnDie:           t.OnDie,
		OnKill:          t.OnKill,
		OnMove:          t.OnMove,
		OnStrike:        t.OnStrike,
		OnStruck:        t.OnStruck,
		OnTurnEnd:       t.OnTurnEnd,
		ActiveAbilities: t.ActiveAbilities,
		Conditions:      make(map[string]Condition),
	}

	return unit
}
