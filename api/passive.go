package api

import (
	"github.com/youngpto/zgz-db/dto"
	"github.com/youngpto/zgz-db/enum"
	"github.com/youngpto/zgz-db/model"
	"github.com/youngpto/zgz-db/service"
)

// SetUserPassive 玩家设置指定层级被动
func SetUserPassive(userId int64, resourceHeroId int64, level int, resourcePassiveId enum.HeroPassive) ([]*dto.UserHeroPassive, error) {
	heroId, err := service.GetHeroId(resourceHeroId)
	if err != nil {
		return nil, err
	}
	if resourcePassiveId == -1 {
		err := service.CancelTakeAlongUserPassiveByLevel(&model.UserPassive{
			UserId: userId,
			HeroId: heroId,
			Level:  level,
		})
		if err != nil {
			return nil, err
		}
	} else {
		passiveId, err := service.GetHeroPassiveId(int(heroId), int(resourcePassiveId))
		if err != nil {
			return nil, err
		}
		err = service.UpdateTakeAlongFormUserPassiveByLevel(&model.UserPassive{
			UserId:    userId,
			HeroId:    heroId,
			Level:     level,
			PassiveId: passiveId,
		})
		if err != nil {
			return nil, err
		}
	}
	return service.FindAllUserPassiveByUser(userId, heroId)
}

// GetUserHeroPassive 获取玩家英雄被动
func GetUserHeroPassive(userId int64, resourceHeroId int64) ([]*dto.UserHeroPassive, error) {
	heroId, err := service.GetHeroId(resourceHeroId)
	if err != nil {
		return nil, err
	}
	return service.FindAllUserPassiveByUser(userId, heroId)
}
