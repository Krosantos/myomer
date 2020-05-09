package elements

import (
	"encoding/json"
)

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
	OnAttack        []string          `json:"onAttack"`
	OnDie           []string          `json:"onDie"`
	OnKill          []string          `json:"onKill"`
	OnMove          []string          `json:"onMove"`
	OnRoundEnd      []string          `json:"onRoundEnd"`
	OnStrike        []string          `json:"onStrike"`
	OnStruck        []string          `json:"onStruck"`
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
		OnAttack:        make(map[string]onAttack),
		OnDie:           make(map[string]onDie),
		OnKill:          make(map[string]onKill),
		OnMove:          make(map[string]onMove),
		OnRoundEnd:      make(map[string]onRoundEnd),
		OnStrike:        make(map[string]onStrike),
		OnStruck:        make(map[string]onStruck),
		ActiveAbilities: make(map[string]ActiveAbility),
		Conditions:      make(map[string]Condition),
	}

	for _, key := range t.OnAttack {
		fn, ok := onAttackRegistry[key]
		if ok == true {
			unit.OnAttack[key] = fn
		}
	}

	for _, key := range t.OnDie {
		fn, ok := onDieRegistry[key]
		if ok == true {
			unit.OnDie[key] = fn
		}
	}

	for _, key := range t.OnKill {
		fn, ok := onKillRegistry[key]
		if ok == true {
			unit.OnKill[key] = fn
		}
	}

	for _, key := range t.OnMove {
		fn, ok := onMoveRegistry[key]
		if ok == true {
			unit.OnMove[key] = fn
		}
	}

	for _, key := range t.OnRoundEnd {
		fn, ok := onRoundEndRegistry[key]
		if ok == true {
			unit.OnRoundEnd[key] = fn
		}
	}

	for _, key := range t.OnStrike {
		fn, ok := onStrikeRegistry[key]
		if ok == true {
			unit.OnStrike[key] = fn
		}
	}

	for _, key := range t.OnStruck {
		fn, ok := onStruckRegistry[key]
		if ok == true {
			unit.OnStruck[key] = fn
		}
	}

	return unit
}
