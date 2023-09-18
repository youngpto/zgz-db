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
	data, err := service.FindAllUserSpecialityByUserAndHero(userId, heroId)
	if err != nil {
		return nil, err
	}
	var tmap = make(map[int]*dto.UserHeroSpeciality)
	for _, datum := range data {
		resourceId, err := service.GetHeroSpecialityResourceId(datum.SpecialityId)
		if err != nil {
			continue
		}
		if old, ok := tmap[datum.Level]; ok {
			if datum.TakeAlong {
				old.TakeAlongSpecialityResourceId = resourceId
				old.ChoosePool = append(old.ChoosePool, resourceId)
			} else {
				old.ChoosePool = append(old.ChoosePool, resourceId)
			}
		} else {
			newDto := &dto.UserHeroSpeciality{
				Level:                         int32(datum.Level),
				TakeAlongSpecialityResourceId: -1,
				ChoosePool:                    make([]enum.HeroSpeciality, 0),
			}
			if datum.TakeAlong {
				newDto.TakeAlongSpecialityResourceId = resourceId
				newDto.ChoosePool = append(newDto.ChoosePool, resourceId)
			} else {
				newDto.ChoosePool = append(newDto.ChoosePool, resourceId)
			}
			tmap[datum.Level] = newDto
		}
	}
	var result []*dto.UserHeroSpeciality
	for _, speciality := range tmap {
		result = append(result, speciality)
	}
	return result, nil
}

// GetUserHeroSpeciality 获取用户英雄所有专长
func GetUserHeroSpeciality(userId int64, resourceHeroId int64) ([]*dto.UserHeroSpeciality, error) {
	heroId, err := service.GetHeroId(resourceHeroId)
	if err != nil {
		return nil, err
	}
	data, err := service.FindAllUserSpecialityByUserAndHero(userId, heroId)
	if err != nil {
		return nil, err
	}
	var tmap = make(map[int]*dto.UserHeroSpeciality)
	for _, datum := range data {
		resourceId, err := service.GetHeroSpecialityResourceId(datum.SpecialityId)
		if err != nil {
			continue
		}
		if old, ok := tmap[datum.Level]; ok {
			if datum.TakeAlong {
				old.TakeAlongSpecialityResourceId = resourceId
				old.ChoosePool = append(old.ChoosePool, resourceId)
			} else {
				old.ChoosePool = append(old.ChoosePool, resourceId)
			}
		} else {
			newDto := &dto.UserHeroSpeciality{
				Level:                         int32(datum.Level),
				TakeAlongSpecialityResourceId: -1,
				ChoosePool:                    make([]enum.HeroSpeciality, 0),
			}
			if datum.TakeAlong {
				newDto.TakeAlongSpecialityResourceId = resourceId
				newDto.ChoosePool = append(newDto.ChoosePool, resourceId)
			} else {
				newDto.ChoosePool = append(newDto.ChoosePool, resourceId)
			}
			tmap[datum.Level] = newDto
		}
	}
	var result []*dto.UserHeroSpeciality
	for _, speciality := range tmap {
		result = append(result, speciality)
	}
	return result, nil
}
