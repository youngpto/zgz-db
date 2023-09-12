package enum

type HeroProperty int

const (
	BaseHp HeroProperty = iota + 1
	BaseSan
	BaseStrength
	BaseAgility
	BaseKnowledge
	BaseWillPower
)
