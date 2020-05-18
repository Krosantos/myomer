package elements

type onStruck func(*Unit, *Unit, int) int

var onStruckRegistry map[string]onStruck

func init() {
	onStruckRegistry = map[string]onStruck{
		"gorgon": func(victim *Unit, attacker *Unit, damage int) int {
			_, isPoisoned := victim.Conditions[poisonID]
			if isPoisoned {
				return 0
			}
			return damage
		},
	}
}
