package game

type onStrike func(*unit, *unit, int) int

var onStrikeRegistry map[string]onStrike

func init() {
	onStrikeRegistry = map[string]onStrike{
		"viper": func(attacker *unit, victim *unit, damage int) int {
			_, isPoisoned := victim.conditions[poisonID]
			if isPoisoned {
				return damage * 2
			}
			return damage
		},
	}
}
