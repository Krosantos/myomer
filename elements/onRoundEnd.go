package elements

type onRoundEnd func(*Unit)

var onRoundEndRegistry = map[string]onRoundEnd{
	"regenerate": func(u *Unit) {
		u.Damage--
	},
}
