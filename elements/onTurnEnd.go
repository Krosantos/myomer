package elements

type onTurnEnd func(*Unit)

var onTurnEndRegistry map[string]onTurnEnd

func init() {
	onTurnEndRegistry = map[string]onTurnEnd{
		"regenerate": func(u *Unit) {
			u.HealDamage(1)
		},
	}

}
