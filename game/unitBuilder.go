package game

import (
	"github.com/google/uuid"
)

type army struct {
	Units map[int]unitTemplate `json:"units"`
}

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

// Currently, there are 19 eligible tiles a unit can start on on either side. This function maps those to a tile coordinate, based on team.
func getUnitTile(pos int, team int, b *board) *tile {
	coord := positionToTile[team][pos]
	return b.get(coord.X, coord.Y)
}

// buildUnit -- Given a unit template, team, and tile, build a unit
func buildUnit(t unitTemplate, team int, tile *tile) unit {
	result := unit{
		id:              uuid.New().String(),
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
		tile:            tile,
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
	tile.unit = &result

	return result
}

var positionToTile map[int]map[int]Coord = map[int]map[int]Coord{
	0: {
		0:  {-6, 0},
		1:  {-6, -1},
		2:  {-6, -2},
		3:  {-6, -3},
		4:  {-6, -4},
		5:  {-6, -5},
		6:  {-5, 1},
		7:  {-5, 0},
		8:  {-5, -1},
		9:  {-5, -2},
		10: {-5, -3},
		11: {-5, -4},
		12: {-5, -5},
		13: {-4, 1},
		14: {-4, 0},
		15: {-4, -1},
		16: {-4, -2},
		17: {-4, -3},
		18: {-4, -4},
	},
	1: {
		0:  {6, 6},
		1:  {6, 5},
		2:  {6, 4},
		3:  {6, 3},
		4:  {6, 2},
		5:  {6, 1},
		6:  {5, 6},
		7:  {5, 5},
		8:  {5, 4},
		9:  {5, 3},
		10: {5, 2},
		11: {5, 1},
		12: {5, 0},
		13: {4, 5},
		14: {4, 4},
		15: {4, 3},
		16: {4, 2},
		17: {4, 1},
		18: {4, 0},
	},
}
