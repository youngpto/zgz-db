package api

import (
	"github.com/youngpto/zgz-db/dto"
	"github.com/youngpto/zgz-db/service"
)

// GetUserHeroGrowthAttribute 获取玩家英雄所有成长相关属性
func GetUserHeroGrowthAttribute(userId int64, resourceHeroId int64) (*dto.HeroInfo, error) {
	heroId, err := service.GetHeroId(resourceHeroId)
	if err != nil {
		return nil, err
	}
	return service.GetPropertyAndPassiveAndSpecialityByUser(userId, heroId)
}

// GetUserAllHeroGrowthAttribute 获取玩家所有英雄的所有成长相关属性
func GetUserAllHeroGrowthAttribute(userId int64) ([]*dto.HeroInfo, error) {
	heroRanks, err := service.GetAllHeroByUser(userId)
	if err != nil {
		return nil, err
	}
	var result []*dto.HeroInfo
	for _, rank := range heroRanks {
		heroInfo, err := service.GetPropertyAndPassiveAndSpecialityByUser(userId, rank.HeroId)
		if err != nil {
			continue
		}
		result = append(result, heroInfo)
	}
	return result, nil
}
