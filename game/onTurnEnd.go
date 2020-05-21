package game

type onTurnEnd func(*unit)

var onTurnEndRegistry map[string]onTurnEnd

func init() {
	onTurnEndRegistry = map[string]onTurnEnd{
		"regenerate": func(u *unit) {
			u.healDamage(1)
		},
	}

}
