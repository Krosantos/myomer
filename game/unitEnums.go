package game

type moveType string

const (
	basic      moveType = "basic"
	flying     moveType = "flying"
	teleport   moveType = "teleport"
	climb      moveType = "climb"
	infiltrate moveType = "infiltrate"
)

type color string

const (
	white color = "white"
	blue  color = "blue"
	black color = "black"
	red   color = "red"
	green color = "green"
)

type targetType string

const (
	adjacent targetType = "adjacent"
	spear    targetType = "spear"
	ranged   targetType = "ranged"
	cone     targetType = "cone"
	pbaoe    targetType = "pbaoe"
)
