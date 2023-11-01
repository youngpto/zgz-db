package api

import (
	"github.com/youngpto/zgz-db/dto"
	"github.com/youngpto/zgz-db/enum"
	"github.com/youngpto/zgz-db/model"
	"github.com/youngpto/zgz-db/service"
)

// SetUserSpeciality 玩家设置指定层级专长
func SetUserSpeciality(userId int64, resourceHeroId int64, level int, resourceSpecialityId enum.HeroSpeciality) ([]*dto.UserHeroSpeciality, error) {
	heroId, err := service.GetHeroId(resourceHeroId)
	if err != nil {
		return nil, err
	}
	if resourceSpecialityId == -1 {
		return nil, nil
		//err := service.CancelTakeAlongUserSpecialityByLevel(&model.UserSpeciality{
		//	UserId: userId,
		//	HeroId: heroId,
		//	Level:  level,
		//})
		//if err != nil {
		//	return nil, err
		//}
	} else {
		specialityId, err := service.GetHeroSpecialityId(int(heroId), int(resourceSpecialityId))
		if err != nil {
			return nil, err
		}
		err = service.UpdateTakeAlongFormUserSpecialityByLevel(&model.UserSpeciality{
			UserId:       userId,
			HeroId:       heroId,
			Level:        level,
			SpecialityId: specialityId,
		})
		if err != nil {
			return nil, err
		}
	}
	return service.FindAllUserSpecialityByUserAndHero(userId, heroId)
}

// GetUserHeroSpeciality 获取用户英雄所有专长
func GetUserHeroSpeciality(userId int64, resourceHeroId int64) ([]*dto.UserHeroSpeciality, error) {
	heroId, err := service.GetHeroId(resourceHeroId)
	if err != nil {
		return nil, err
	}
	return service.FindAllUserSpecialityByUserAndHero(userId, heroId)
}
