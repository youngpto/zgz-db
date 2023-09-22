package dto

// HeroInfo 玩家英雄配置的各项数据
type HeroInfo struct {
	HeroId int
	GainExperienceRes
	UserHeroProperty
	Speciality []*UserHeroSpeciality
	Passive    []*UserHeroPassive
}
