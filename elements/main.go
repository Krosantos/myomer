package elements

type moveType string

const (
	basic      moveType = "basic"
	flying     moveType = "flying"
	teleport   moveType = "teleport"
	climb      moveType = "climb"
	infiltrate moveType = "infiltrate"
)
