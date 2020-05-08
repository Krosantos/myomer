package elements

type moveType string

const (
	basic      moveType = "basic"
	flying     moveType = "flying"
	teleport   moveType = "teleport"
	climb      moveType = "climb"
	infiltrate moveType = "infiltrate"
)

type attackType string

const (
	adjacent attackType = "adjacent"
	spear    attackType = "spear"
	ranged   attackType = "ranged"
	cone     attackType = "cone"
	pbaoe    attackType = "pbaoe"
)
