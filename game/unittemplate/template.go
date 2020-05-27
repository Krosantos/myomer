package unit

// Template -- A template used to instantiate a unit
type Template struct {
	TemplateID      string   `json:"templateId"`
	Name            string   `json:"name"`
	Cost            int      `json:"cost"`
	Color           string   `json:"color"`
	Strength        int      `json:"strength"`
	Health          int      `json:"health"`
	Speed           int      `json:"speed"`
	Moxie           int      `json:"moxie"`
	AttackRange     int      `json:"attackRange"`
	AttackType      string   `json:"attackType"`
	MoveType        string   `json:"moveType"`
	OnAttack        []string `json:"onAttack"`
	OnDie           []string `json:"onDie"`
	OnKill          []string `json:"onKill"`
	OnMove          []string `json:"onMove"`
	OnStrike        []string `json:"onStrike"`
	OnStruck        []string `json:"onStruck"`
	OnTurnEnd       []string `json:"onTurnEnd"`
	ActiveAbilities []string `json:"activeAbilities"`
}
