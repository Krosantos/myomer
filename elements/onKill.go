package elements

type onKill func(*Unit, *Unit)

var onKillRegistry map[string]onKill

func init() {
	onKillRegistry = map[string]onKill{
		"highlander": func(killer *Unit, victim *Unit) {
			killer.Strength += victim.Strength
		},
	}
}
