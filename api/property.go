package api

import (
	"github.com/youngpto/zgz-db/dto"
	"github.com/youngpto/zgz-db/service"
)

// FindUserHeroTotalProperty 获取玩家六维属性基础值及偏移量
func FindUserHeroTotalProperty(userId int64, heroId int64) *dto.UserHeroProperty {
	result := new(dto.UserHeroProperty)
	hero, ok := service.FindPropertyOffsetByUserAndHero(userId, heroId)
	if ok {
		result.BaseHp += int32(hero.LifeOffset)
		result.BaseSan += int32(hero.ReasonOffset)
		result.BaseStrength += int32(hero.PowerOffset)
		result.BaseAgility += int32(hero.AgileOffset)
		result.BaseKnowledge += int32(hero.KnowledgeOffset)
		result.BaseWillPower += int32(hero.WillOffset)
	}
	property, ok := service.FindHeroBaseProperty(heroId)
	if ok {
		result.BaseHp += int32(property.Life)
		result.BaseSan += int32(property.Reason)
		result.BaseStrength += int32(property.Power)
		result.BaseAgility += int32(property.Agile)
		result.BaseKnowledge += int32(property.Knowledge)
		result.BaseWillPower += int32(property.Will)
	}
	return result
}
