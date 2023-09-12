package api

import (
	"errors"
	"github.com/youngpto/zgz-db/dto"
	"github.com/youngpto/zgz-db/enum"
	"github.com/youngpto/zgz-db/model"
	"github.com/youngpto/zgz-db/service"
)

// SetUserSpeciality 设置玩家专长
func SetUserSpeciality(userId int64, resourceHeroId int64, level int, resourceSpecialityId enum.HeroSpeciality) ([]dto.UserHeroSpeciality, error) {
	// 获取英雄主键Id
	heroId, ok := service.GetHeroId(resourceHeroId)
	if !ok {
		return nil, errors.New("未找到调查员类型")
	}
	// 获取专长主键Id
	specialityId, ok := service.GetHeroSpecialityId(int(heroId), int(resourceSpecialityId))
	if !ok {
		return nil, errors.New("该专长无法被该调查员使用")
	}
	// 设置专长
	speciality := &model.UserSpeciality{
		UserId:       userId,
		HeroId:       heroId,
		Level:        level,
		SpecialityId: specialityId,
	}
	err := service.InsertOrUpdateUserSpecialityByLevel(speciality)
	if err != nil {
		return nil, err
	}
	// 获取该玩家该英雄所有专长设置
	res, err := service.FindUserSpecialityByUserAndHero(userId, heroId)
	if err != nil {
		return nil, err
	}
	var sPkId []int64
	for _, re := range res {
		sPkId = append(sPkId, re.SpecialityId)
	}
	resourceIds, err := service.GetSpecialityResourceIdByKeyID(sPkId...)
	if err != nil {
		return nil, err
	}
	var result []dto.UserHeroSpeciality
	for index, re := range res {
		result = append(result, dto.UserHeroSpeciality{
			Level:         int32(re.Level),
			SpecialityRId: resourceIds[index],
		})
	}
	return result, nil
}
