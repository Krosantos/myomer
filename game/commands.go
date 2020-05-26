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

// Command -- an enum of commands
var Command commandsEnum = commandsEnum{
	MOVE:    "MOVE",
	ABILITY: "ABILITY",
	ATTACK:  "ATTACK",
	ENDTURN: "ENDTURN",
	FORFEIT: "FORFEIT",
}

type moveCommand struct {
	Command string `json:"command"`
	Unit    string `json:"unit"`
	Tile    Coord  `json:"tile"`
}

type abilityCommand struct {
	Command string  `json:"command"`
	Unit    string  `json:"unit"`
	Targets []Coord `json:"targets"`
}

type attackCommand struct {
	Command string `json:"command"`
	Unit    string `json:"unit"`
	Target  Coord  `json:"target"`
}

type endTurnCommand struct {
	Command string `json:"command"`
}

type forfeitCommand struct {
	Command   string `json:"command"`
	Automatic bool   `json:"automatic"`
}

// FormatCommand -- Given a string in, attempt to marshal a command
func FormatCommand(s string) (interface{}, error) {
	raw := make(map[string]interface{})
	err := json.Unmarshal([]byte(s), &raw)
	if err != nil {
		return nil, err
	}
	cmd, ok := raw["command"]
	if ok != true {
		return nil, errors.New("no command specified")
	}
	switch cmd {
	case Command.MOVE:
		res := moveCommand{}
		err = json.Unmarshal([]byte(s), &res)
		return res, err
	case Command.ABILITY:
		res := abilityCommand{}
		err = json.Unmarshal([]byte(s), &res)
		return res, err
	case Command.ATTACK:
		res := attackCommand{}
		err = json.Unmarshal([]byte(s), &res)
		return res, err
	case Command.ENDTURN:
		res := endTurnCommand{}
		err = json.Unmarshal([]byte(s), &res)
		return res, err
	case Command.FORFEIT:
		res := forfeitCommand{}
		err = json.Unmarshal([]byte(s), &res)
		return res, err
	default:
		return nil, errors.New("unrecognized command")
	}
}
