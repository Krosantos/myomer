package elements

type moveType string

const (
	basic      moveType = "basic"
	flying     moveType = "flying"
	teleport   moveType = "teleport"
	climb      moveType = "climb"
	infiltrate moveType = "infiltrate"
)

type targetType string

const (
	adjacent targetType = "adjacent"
	spear    targetType = "spear"
	ranged   targetType = "ranged"
	cone     targetType = "cone"
	pbaoe    targetType = "pbaoe"
)
