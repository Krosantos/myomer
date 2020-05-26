package game

/*
Apologies for the jargon. Instructions are sent out from the game to communicate what has happened.
Commands are sent to the game, and parsed to cause action.
*/

type actionsEnum struct {
	MOVE            string
	ANIMATE         string
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

var action actionsEnum = actionsEnum{
	MOVE:            "MOVE",
	ANIMATE:         "ANIMATE",
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
	Tile   coord  `json:"tile"`
}

type animateInstruction struct {
	Action  string  `json:"action"`
	Unit    string  `json:"unit"`
	Targets []coord `json:"targets"`
}

type addConditionInstruction struct {
	Action   string `json:"action"`
	Unit     string `json:"unit"`
	StatusID string `json:"statusId"`
	Duration int    `json:"duration"`
}

type removeConditionInstruction struct {
	Action   string `json:"action"`
	Unit     string `json:"unit"`
	StatusID string `json:"statusId"`
}

type damageInstruction struct {
	Action string `json:"action"`
	Unit   string `json:"unit"`
	Amount int    `json:"amount"`
}

type healInstruction struct {
	Action string `json:"action"`
	Unit   string `json:"unit"`
	Amount int    `json:"amount"`
}

type dieInstruction struct {
	Action string `json:"action"`
	Unit   string `json:"unit"`
}

type endTurnInstruction struct {
	Action   string `json:"action"`
	NextUnit string `json:"nextUnit"`
}

type endGameInstruction struct {
	Action string `json:"action"`
	Winner int    `json:"winner"`
}

// TODO: -- flesh this out, along with general "them's the units"
type addUnitInstruction struct {
	Action string `json:"action"`
}

type removeCorpseInstruction struct {
	Action string `json:"action"`
	Tile   coord  `json:"tile"`
}

type reanimateInstruction struct {
	Action string `json:"action"`
	Unit   string `json:"unit"`
}
