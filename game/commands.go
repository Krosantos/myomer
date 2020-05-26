package game

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

var command commandsEnum = commandsEnum{
	MOVE:    "MOVE",
	ABILITY: "ABILITY",
	ATTACK:  "ATTACK",
	ENDTURN: "ENDTURN",
	FORFEIT: "FORFEIT",
}

type moveCommand struct {
	Command string `json:"command"`
	Unit    string `json:"unit"`
	Tile    coord  `json:"tile"`
}

type abilityCommand struct {
	Command string  `json:"command"`
	Unit    string  `json:"unit"`
	Targets []coord `json:"targets"`
}

type attackCommand struct {
	Command string `json:"command"`
	Unit    string `json:"unit"`
	Target  coord  `json:"target"`
}

type endTurnCommand struct {
	Command string `json:"command"`
}

type forfeitCommand struct {
	Command string `json:"command"`
}
