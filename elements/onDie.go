package elements

type onDie func(*Unit, *Unit)

var onDieRegistry map[string]onDie

func init() {
	onDieRegistry = map[string]onDie{
		"aftermath": func(victim *Unit, killer *Unit) {
			killer.TakeDamage(victim, 4)
		},
	}
}
