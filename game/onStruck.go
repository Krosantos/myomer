package game

type onStruck func(*unit, *unit, int) int

var onStruckRegistry map[string]onStruck

func init() {
	onStruckRegistry = map[string]onStruck{
		"gorgon": func(victim *unit, attacker *unit, damage int) int {
			_, isPoisoned := victim.conditions[poisonID]
			if isPoisoned {
				return 0
			}
			return damage
		},
	}
}
