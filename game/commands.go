package game

import (
	"encoding/json"
	"errors"
)

/*
Apologies for the jargon. Instructions are sent out from the game to communicate what has happened.
Commands are sent to the game, and parsed to cause action.
*/

type commandsEnum struct {
	MOVE    string
	ABILITY string
	ATTACK  string
	ENDTURN string
	FORFEIT string
}

// Command -- the lazy interface for sending moves to the game
type Command interface{}

// cmdAction -- an enum of commands
var cmdAction commandsEnum = commandsEnum{
	MOVE:    "MOVE",
	ABILITY: "ABILITY",
	ATTACK:  "ATTACK",
	ENDTURN: "ENDTURN",
	FORFEIT: "FORFEIT",
}

type moveCommand struct {
	Action string `json:"action"`
	Unit   string `json:"unit"`
	Tile   Coord  `json:"tile"`
	Command
}

type abilityCommand struct {
	Action    string  `json:"action"`
	Unit      string  `json:"unit"`
	AbilityID string  `json:"abilityId"`
	Targets   []Coord `json:"targets"`
	Command
}

type attackCommand struct {
	Action string `json:"action"`
	Unit   string `json:"unit"`
	Target Coord  `json:"target"`
	Command
}

type endTurnCommand struct {
	Action string `json:"action"`
	Command
}

type forfeitCommand struct {
	Action    string `json:"action"`
	Automatic bool   `json:"automatic"`
	Command
}

// FormatCommand -- Given a string in, attempt to marshal a command
func FormatCommand(s string) (Command, error) {
	raw := make(map[string]Command)
	err := json.Unmarshal([]byte(s), &raw)
	if err != nil {
		return nil, err
	}
	cmd, ok := raw["command"]
	if ok != true {
		return nil, errors.New("no command specified")
	}
	switch cmd {
	case cmdAction.MOVE:
		res := moveCommand{}
		err = json.Unmarshal([]byte(s), &res)
		return res, err
	case cmdAction.ABILITY:
		res := abilityCommand{}
		err = json.Unmarshal([]byte(s), &res)
		return res, err
	case cmdAction.ATTACK:
		res := attackCommand{}
		err = json.Unmarshal([]byte(s), &res)
		return res, err
	case cmdAction.ENDTURN:
		res := endTurnCommand{}
		err = json.Unmarshal([]byte(s), &res)
		return res, err
	case cmdAction.FORFEIT:
		res := forfeitCommand{}
		err = json.Unmarshal([]byte(s), &res)
		return res, err
	default:
		return nil, errors.New("unrecognized command")
	}
}
