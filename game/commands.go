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
	Team   int
	Command
}

type abilityCommand struct {
	Action    string  `json:"action"`
	Unit      string  `json:"unit"`
	AbilityID string  `json:"abilityId"`
	Targets   []Coord `json:"targets"`
	Team      int
	Command
}

type attackCommand struct {
	Action string `json:"action"`
	Unit   string `json:"unit"`
	Target Coord  `json:"target"`
	Team   int
	Command
}

type endTurnCommand struct {
	Action string `json:"action"`
	Team   int
	Command
}

type forfeitCommand struct {
	Action    string `json:"action"`
	Automatic bool   `json:"automatic"`
	Team      int
	Command
}

// FormatCommand -- Given a string in, attempt to marshal a command
func FormatCommand(s string, t int) (Command, error) {
	raw := make(map[string]Command)
	err := json.Unmarshal([]byte(s), &raw)
	if err != nil {
		return nil, err
	}
	cmd, ok := raw["action"]
	if ok != true {
		return nil, errors.New("no command specified")
	}
	switch cmd {
	case cmdAction.MOVE:
		res := moveCommand{}
		err = json.Unmarshal([]byte(s), &res)
		res.Team = t
		return res, err
	case cmdAction.ABILITY:
		res := abilityCommand{}
		err = json.Unmarshal([]byte(s), &res)
		res.Team = t
		return res, err
	case cmdAction.ATTACK:
		res := attackCommand{}
		err = json.Unmarshal([]byte(s), &res)
		res.Team = t
		return res, err
	case cmdAction.ENDTURN:
		res := endTurnCommand{}
		err = json.Unmarshal([]byte(s), &res)
		res.Team = t
		return res, err
	case cmdAction.FORFEIT:
		println("I am definitely telling you it is forfeit")
		res := forfeitCommand{}
		err = json.Unmarshal([]byte(s), &res)
		res.Team = t
		return res, err
	default:
		return nil, errors.New("unrecognized command")
	}
}
