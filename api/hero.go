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
