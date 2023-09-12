package api

import (
	"errors"
	"github.com/youngpto/zgz-db/dto"
	"github.com/youngpto/zgz-db/enum"
	"github.com/youngpto/zgz-db/model"
	"github.com/youngpto/zgz-db/service"
)

// SetUserPassive 设置玩家被动
func SetUserPassive(userId int64, resourceHeroId int64, level int, resourcePassiveId enum.HeroPassive) ([]dto.UserHeroPassive, error) {
	// 获取英雄主键Id
	heroId, ok := service.GetHeroId(resourceHeroId)
	if !ok {
		return nil, errors.New("未找到调查员类型")
	}
	// 获取专长主键Id
	passiveId, ok := service.GetHeroPassiveId(int(heroId), int(resourcePassiveId))
	if !ok {
		return nil, errors.New("该专长无法被该调查员使用")
	}
	// 设置专长
	passive := &model.UserPassive{
		UserId:    userId,
		HeroId:    heroId,
		Level:     level,
		PassiveId: passiveId,
	}
	err := service.InsertOrUpdateUserPassiveByLevel(passive)
	if err != nil {
		return nil, err
	}
	// 获取该玩家该英雄所有专长设置
	res, err := service.FindUserPassiveByUserAndHero(userId, heroId)
	if err != nil {
		return nil, err
	}
	var sPkId []int64
	for _, re := range res {
		sPkId = append(sPkId, re.PassiveId)
	}
	resourceIds, err := service.GetPassiveResourceIdByKeyID(sPkId...)
	if err != nil {
		return nil, err
	}
	var result []dto.UserHeroPassive
	for index, re := range res {
		result = append(result, dto.UserHeroPassive{
			Level:         int32(re.Level),
			SpecialityRId: resourceIds[index],
		})
	}
	return result, nil
}
