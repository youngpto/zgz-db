package enum

type HeroPassive int

const (
	HeroPassive1 HeroPassive = iota + 1 // 手牌上限+3。
	HeroPassive2                        // 开局受到1点伤害、1点恐惧，额外抽3牌。
	HeroPassive3                        // 开局时，将1张<逢凶化吉>加入你的手牌。
	HeroPassive4                        // 开局受到1点伤害、1点恐惧，装备槽+1。
	HeroPassive5                        // 你的首个回合开始时，从卡组检索1牌。
	HeroPassive6                        // 每局首次被击败时，保留1点生命/理智。
)
