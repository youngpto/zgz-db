package enum

type HeroResource int64

const (
	Taoist    HeroResource = iota // 道士
	Gambler                       // 赌徒
	Doctor                        // 医生
	College                       // 大学生
	Police                        // 警察
	Pupil                         // 小学生
	Mercenary                     // 佣兵
)
