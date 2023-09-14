package api

import (
	"errors"
	"github.com/youngpto/zgz-db/dto"
	"github.com/youngpto/zgz-db/enum"
	"github.com/youngpto/zgz-db/service"
)

// UpdatePropertyOffset 修改玩家属性偏移值并返回目前存储的所有六维数值
func UpdatePropertyOffset(userId, resourceHeroId int64, property enum.HeroProperty, isIncr bool) (*dto.UserHeroProperty, error) {
	result := new(dto.UserHeroProperty)
	// 先获取英雄表的基础属性和主键id
	heroInfo, err := service.FindHeroBaseProperty(resourceHeroId)
	if err != nil {
		return nil, errors.New("未找到对应英雄")
	}
	result.BaseHp += int32(heroInfo.Life)
	result.BaseSan += int32(heroInfo.Reason)
	result.BaseStrength += int32(heroInfo.Power)
	result.BaseAgility += int32(heroInfo.Agile)
	result.BaseKnowledge += int32(heroInfo.Knowledge)
	result.BaseWillPower += int32(heroInfo.Will)

	// 增加偏移量
	offset, err := service.UpdatePropertyOffsetByType(userId, heroInfo.Id, property, isIncr)
	if err != nil {
		return nil, err
	}
	result.BaseHp += int32(offset.LifeOffset)
	result.BaseSan += int32(offset.ReasonOffset)
	result.BaseStrength += int32(offset.PowerOffset)
	result.BaseAgility += int32(offset.AgileOffset)
	result.BaseKnowledge += int32(offset.KnowledgeOffset)
	result.BaseWillPower += int32(offset.WillOffset)
	result.AvailablePoints = int32(offset.TotalOffset - offset.LifeOffset - offset.ReasonOffset - offset.PowerOffset - offset.AgileOffset - offset.KnowledgeOffset - offset.WillOffset)

	return result, nil
}

// ResetPropertyOffset 重置玩家属性加点并返回现有属性值
func ResetPropertyOffset(userId, resourceHeroId int64) (*dto.UserHeroProperty, error) {
	result := new(dto.UserHeroProperty)
	// 先获取英雄表的基础属性和主键id
	heroInfo, err := service.FindHeroBaseProperty(resourceHeroId)
	if err != nil {
		return nil, errors.New("未找到对应英雄")
	}
	result.BaseHp += int32(heroInfo.Life)
	result.BaseSan += int32(heroInfo.Reason)
	result.BaseStrength += int32(heroInfo.Power)
	result.BaseAgility += int32(heroInfo.Agile)
	result.BaseKnowledge += int32(heroInfo.Knowledge)
	result.BaseWillPower += int32(heroInfo.Will)

	//重置偏移量
	offset, err := service.ResetPropertyOffset(userId, heroInfo.Id)
	if err != nil {
		return nil, err
	}
	result.AvailablePoints = int32(offset.TotalOffset)
	return result, nil
}
