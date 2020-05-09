package elements

import "encoding/json"

type unitTemplate struct {
	Name            string            `json:"name"`
	Cost            int               `json:"cost"`
	Strength        int               `json:"strength"`
	Health          int               `json:"health"`
	Speed           int               `json:"speed"`
	Moxie           int               `json:"moxie"`
	AttackRange     int               `json:"attackRange"`
	AttackType      targetType        `json:"attackType"`
	MoveType        moveType          `json:"moveType"`
	OnMove          []string          `json:"onMove"`
	OnAttack        []string          `json:"onAttack"`
	OnDie           []string          `json:"onDie"`
	OnStrike        []string          `json:"onStrike"`
	OnStruck        []string          `json:"onStruck"`
	OnKill          []string          `json:"onKill"`
	OnRoundEnd      []string          `json:"onRoundEnd"`
	ActiveAbilities []abilityTemplate `json:"activeAbilities"`
}

type abilityTemplate struct {
	ID   string        `json:"id"`
	Args []interface{} `json:"args"`
}

// BuildUnit -- Given a JSON template, build a unit
func BuildUnit(s string) Unit {
	var t unitTemplate
	err := json.Unmarshal([]byte(s), &t)
	if err != nil {
		panic(err)
	}
	unit := Unit{
		Name:            t.Name,
		Cost:            t.Cost,
		Strength:        t.Strength,
		Health:          t.Health,
		Speed:           t.Speed,
		Moxie:           t.Moxie,
		AttackRange:     t.AttackRange,
		MoveType:        t.MoveType,
		OnMove:          make(map[string]func(*Unit, *Tile)),
		OnAttack:        make(map[string]func(*Unit, *Tile)),
		OnDie:           make(map[string]func(*Unit, *Unit)),
		OnStrike:        make(map[string]func(*Unit, *Unit, int) int),
		OnStruck:        make(map[string]func(*Unit, *Unit, int) int),
		OnKill:          make(map[string]func(*Unit, *Unit)),
		OnRoundEnd:      make(map[string]func(*Unit)),
		ActiveAbilities: make(map[string]ActiveAbility),
		Conditions:      make(map[string]Condition),
	}

	return unit
}
