package units

type unit interface {
	move()
	attack()
	onDie()
	onStruck()
	onStrike()
	onKill()
	onRoundEnd()
}

type moveType string

const (
	basic      moveType = "basic"
	flying     moveType = "flying"
	teleport   moveType = "teleport"
	climb      moveType = "climb"
	infiltrate moveType = "infiltrate"
)

type baseUnit struct {
	name     string
	cost     int
	strength int
	health   int
	damage   int
	atkRange int
	speed    int
	moxie    int
	moveType moveType
}
