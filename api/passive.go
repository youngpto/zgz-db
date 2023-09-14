package api

import (
	"github.com/youngpto/zgz-db/dto"
	"github.com/youngpto/zgz-db/enum"
	"github.com/youngpto/zgz-db/model"
	"github.com/youngpto/zgz-db/service"
)

// SetUserPassive 玩家设置指定层级被动
func SetUserPassive(userId int64, resourceHeroId int64, level int, resourcePassiveId enum.HeroPassive) ([]dto.UserHeroPassive, error) {
	heroId, err := service.GetHeroId(resourceHeroId)
	if err != nil {
		return nil, err
	}
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
	data, err := service.FindAllUserPassiveByUser(userId, heroId)
	if err != nil {
		return nil, err
	}
	var tmap = make(map[int]dto.UserHeroPassive)
	for _, datum := range data {
		resourceId, err := service.GetHeroPassiveResourceId(datum.PassiveId)
		if err != nil {
			continue
		}
		if old, ok := tmap[datum.Level]; ok {
			if datum.TakeAlong {
				old.TakeAlongPassiveResourceId = resourceId
			} else {
				old.ChoosePool = append(old.ChoosePool, resourceId)
			}
		} else {
			newDto := dto.UserHeroPassive{
				Level:                      int32(datum.Level),
				TakeAlongPassiveResourceId: 0,
				ChoosePool:                 make([]enum.HeroPassive, 0),
			}
			if datum.TakeAlong {
				newDto.TakeAlongPassiveResourceId = resourceId
			} else {
				newDto.ChoosePool = append(newDto.ChoosePool, resourceId)
			}
			tmap[datum.Level] = newDto
		}
	}
	var result []dto.UserHeroPassive
	for _, passive := range tmap {
		result = append(result, passive)
	}
	return result, nil
}
