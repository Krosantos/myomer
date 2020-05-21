package game

type onKill func(*unit, *unit)

var onKillRegistry map[string]onKill

func init() {
	onKillRegistry = map[string]onKill{
		"highlander": func(killer *unit, victim *unit) {
			killer.strength += victim.strength
		},
	}
}
