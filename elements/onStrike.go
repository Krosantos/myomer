package elements

type onStrike func(*Unit, *Unit, int) int

var onStrikeRegistry map[string]onStrike

func init() {
	onStrikeRegistry = map[string]onStrike{
		"viper": func(attacker *Unit, victim *Unit, damage int) int {
			_, isPoisoned := victim.Conditions[poisonID]
			if isPoisoned {
				return damage * 2
			}
			return damage
		},
	}
}
