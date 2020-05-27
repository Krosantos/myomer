package unittemplate

type moveTypeEnum struct {
	Basic      string
	Flying     string
	Teleport   string
	Climb      string
	Infiltrate string
}

// MoveType -- An enum of move types
var MoveType moveTypeEnum = moveTypeEnum{
	Basic:      "basic",
	Flying:     "flying",
	Teleport:   "teleport",
	Climb:      "climb",
	Infiltrate: "infiltrate",
}

type colorEnum struct {
	Black  string
	Blue   string
	Green  string
	Red    string
	White  string
	Yellow string
}

// Color -- An enum of unit colors
var Color colorEnum = colorEnum{
	Black:  "black",
	Blue:   "blue",
	Green:  "green",
	Red:    "red",
	White:  "white",
	Yellow: "yellow",
}

type targetTypeEnum struct {
	Adjacent string
	Spear    string
	Ranged   string
	Cone     string
	Pbaoe    string
}

// TargetType -- An enum of target types
var TargetType targetTypeEnum = targetTypeEnum{
	Adjacent: "adjacent",
	Spear:    "spear",
	Ranged:   "ranged",
	Cone:     "cone",
	Pbaoe:    "pbaoe",
}
