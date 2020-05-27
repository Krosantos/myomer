package game

import (
	"encoding/json"

	"github.com/krosantos/myomer/v2/game/unittemplate"
)

/*
Apologies for the jargon. Instructions are sent out from the game to communicate what has happened.
Commands are sent to the game, and parsed to cause action.
*/

type instructionsEnum struct {
	MOVE            string
	ATTACK          string
	ABILITY         string
	ADDCONDITION    string
	REMOVECONDITION string
	DAMAGE          string
	HEAL            string
	DIE             string
	ENDTURN         string
	ENDGAME         string
	ADDUNIT         string
	REMOVECORPSE    string
	REANIMATE       string
	GAMEEND         string
}

// Instruction -- A command emitted by the game, to be sent to players
type Instruction interface {
	ToString() string
}

var action instructionsEnum = instructionsEnum{
	MOVE:            "MOVE",
	ATTACK:          "ATTACK",
	ABILITY:         "ABILITY",
	ADDCONDITION:    "ADDCONDITION",
	REMOVECONDITION: "REMOVECONDITION",
	DAMAGE:          "DAMAGE",
	HEAL:            "HEAL",
	DIE:             "DIE",
	ENDTURN:         "ENDTURN",
	ENDGAME:         "ENDGAME",
	ADDUNIT:         "ADDUNIT",
	REMOVECORPSE:    "REMOVECORPSE",
	REANIMATE:       "REANIMATE",
}

type moveInstruction struct {
	Action string `json:"action"`
	Unit   string `json:"unit"`
	Tile   Coord  `json:"tile"`
}

func (i moveInstruction) ToString() string {
	raw, _ := json.Marshal(i)
	return string(raw)
}

type abilityInstruction struct {
	Action    string  `json:"action"`
	AbilityID string  `json:"abilityId"`
	Unit      string  `json:"unit"`
	Targets   []Coord `json:"targets"`
}

func (i abilityInstruction) ToString() string {
	raw, _ := json.Marshal(i)
	return string(raw)
}

type attackInstruction struct {
	Action string `json:"action"`
	Unit   string `json:"unit"`
	Target Coord  `json:"target"`
}

func (i attackInstruction) ToString() string {
	raw, _ := json.Marshal(i)
	return string(raw)
}

type addConditionInstruction struct {
	Action   string `json:"action"`
	Unit     string `json:"unit"`
	StatusID string `json:"statusId"`
	Duration int    `json:"duration"`
}

func (i addConditionInstruction) ToString() string {
	raw, _ := json.Marshal(i)
	return string(raw)
}

type removeConditionInstruction struct {
	Action   string `json:"action"`
	Unit     string `json:"unit"`
	StatusID string `json:"statusId"`
}

func (i removeConditionInstruction) ToString() string {
	raw, _ := json.Marshal(i)
	return string(raw)
}

type damageInstruction struct {
	Action string `json:"action"`
	Unit   string `json:"unit"`
	Amount int    `json:"amount"`
}

func (i damageInstruction) ToString() string {
	raw, _ := json.Marshal(i)
	return string(raw)
}

type healInstruction struct {
	Action string `json:"action"`
	Unit   string `json:"unit"`
	Amount int    `json:"amount"`
}

func (i healInstruction) ToString() string {
	raw, _ := json.Marshal(i)
	return string(raw)
}

type dieInstruction struct {
	Action string `json:"action"`
	Unit   string `json:"unit"`
}

func (i dieInstruction) ToString() string {
	raw, _ := json.Marshal(i)
	return string(raw)
}

type endTurnInstruction struct {
	Action   string `json:"action"`
	NextUnit string `json:"nextUnit"`
}

func (i endTurnInstruction) ToString() string {
	raw, _ := json.Marshal(i)
	return string(raw)
}

type endGameInstruction struct {
	Action string `json:"action"`
	Winner int    `json:"winner"`
}

func (i endGameInstruction) ToString() string {
	raw, _ := json.Marshal(i)
	return string(raw)
}

type addUnitInstruction struct {
	Action   string                `json:"action"`
	Template unittemplate.Template `json:"template"`
}

func (i addUnitInstruction) ToString() string {
	raw, _ := json.Marshal(i)
	return string(raw)
}

type removeCorpseInstruction struct {
	Action string `json:"action"`
	Tile   Coord  `json:"tile"`
}

func (i removeCorpseInstruction) ToString() string {
	raw, _ := json.Marshal(i)
	return string(raw)
}

type reanimateInstruction struct {
	Action string `json:"action"`
	Unit   string `json:"unit"`
}

func (i reanimateInstruction) ToString() string {
	raw, _ := json.Marshal(i)
	return string(raw)
}
