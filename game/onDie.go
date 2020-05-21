package game

type onDie func(*unit, *unit)

var onDieRegistry map[string]onDie

func init() {
	onDieRegistry = map[string]onDie{
		"aftermath": func(victim *unit, killer *unit) {
			killer.takeDamage(victim, 4)
		},
	}
}
