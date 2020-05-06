package elements

// Condition -- A temporary unit status condition; buffs and debuffs
type Condition interface {
	onAdd()
	onRemove()
	name()
}
