package service

import (
	"github.com/youngpto/zgz-db/base"
	"github.com/youngpto/zgz-db/dto"
	"github.com/youngpto/zgz-db/enum"
	"github.com/youngpto/zgz-db/model"
	"xorm.io/builder"
)

// GetHeroId 根据资源ID获取英雄表主键
func GetHeroId(resourceId int64) (int64, error) {
	result := new(model.HeroProperty)
	_, err := base.Engine.Table("hero").
		Where(builder.Eq{"resource_id": resourceId}).
		Cols("id").
		Get(result)
	if err != nil {
		return 0, err
	}
	return result.Id, nil
}

// GetAllHeroByUser 获取玩家所有的英雄
func GetAllHeroByUser(userId int64) ([]model.UserHeroRank, error) {
	var result []model.UserHeroRank
	err := base.Engine.Table("user_hero").
		Where(builder.Eq{"user_id": userId}).
		Cols("hero_id").
		Find(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetUserHeroRankInfo 获取玩家对应英雄的等级情况
func GetUserHeroRankInfo(userId int64, heroId int64) (*model.UserHeroRank, error) {
	result := new(model.UserHeroRank)
	_, err := base.Engine.Table("user_hero").
		Where(builder.Eq{"user_id": userId, "hero_id": heroId}).
		Cols("id", "user_id", "hero_id", "rank", "experience_pool").
		Get(result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetPropertyAndPassiveAndSpecialityByUser 查询用户配置的英雄属性和被动以及专长
func GetPropertyAndPassiveAndSpecialityByUser(userId int64, heroId int64) (*dto.HeroInfo, error) {
	session := base.Engine.NewSession()
	defer session.Close()

	if err := session.Begin(); err != nil {
		return nil, err
	}
	// 查询用户等级和属性偏移值
	userHeroRankAndPropertyOffset := model.UserHeroRankAndPropertyOffset{}
	_, err := session.Table("user_hero").
		Where(builder.Eq{
			"user_id": userId,
			"hero_id": heroId,
		}).Get(&userHeroRankAndPropertyOffset)
	if err != nil {
		return nil, err
	}
	// 根据等级查询当前经验值上限
	rank := model.Rank{}
	_, err = session.Where(builder.Eq{"rank": userHeroRankAndPropertyOffset.Rank}).
		Cols("experience").
		Get(&rank)
	if err != nil {
		return nil, err
	}
	// 根据HeroID查询英雄基础属性
	heroProperty := model.HeroProperty{}
	_, err = session.Table("hero").
		ID(heroId).
		Get(&heroProperty)
	if err != nil {
		return nil, err
	}
	// 获取专长
	var speciality []model.UserSpeciality
	err = base.Engine.Where(builder.Eq{
		"user_id": userId,
		"hero_id": heroId,
	}).Cols("level", "speciality_id", "take_along").
		Find(&speciality)
	if err != nil {
		return nil, err
	}
	// 获取被动
	var passive []model.UserPassive
	err = base.Engine.Where(builder.Eq{
		"user_id": userId,
		"hero_id": heroId,
	}).Cols("level", "passive_id", "take_along").
		Find(&passive)
	if err != nil {
		return nil, err
	}

	result := &dto.HeroInfo{
		GainExperienceRes: dto.GainExperienceRes{
			CurrentLevel:  userHeroRankAndPropertyOffset.Rank,
			CurrentExp:    userHeroRankAndPropertyOffset.ExperiencePool,
			CurrentMaxExp: rank.Experience,
		},
		UserHeroProperty: dto.UserHeroProperty{
			BaseHp:        int32(heroProperty.Life + userHeroRankAndPropertyOffset.LifeOffset),
			BaseSan:       int32(heroProperty.Reason + userHeroRankAndPropertyOffset.ReasonOffset),
			BaseStrength:  int32(heroProperty.Power + userHeroRankAndPropertyOffset.PowerOffset),
			BaseAgility:   int32(heroProperty.Agile + userHeroRankAndPropertyOffset.AgileOffset),
			BaseKnowledge: int32(heroProperty.Knowledge + userHeroRankAndPropertyOffset.KnowledgeOffset),
			BaseWillPower: int32(heroProperty.Will + userHeroRankAndPropertyOffset.WillOffset),
			AvailablePoints: int32(userHeroRankAndPropertyOffset.TotalOffset -
				userHeroRankAndPropertyOffset.LifeOffset -
				userHeroRankAndPropertyOffset.ReasonOffset -
				userHeroRankAndPropertyOffset.PowerOffset -
				userHeroRankAndPropertyOffset.AgileOffset -
				userHeroRankAndPropertyOffset.KnowledgeOffset -
				userHeroRankAndPropertyOffset.WillOffset),
		},
		Speciality: make([]dto.UserHeroSpeciality, 0),
		Passive:    make([]dto.UserHeroPassive, 0),
	}

	var sMap = make(map[int]dto.UserHeroSpeciality)
	for _, userSpeciality := range speciality {
		resourceID, _ := GetHeroSpecialityResourceId(userSpeciality.SpecialityId)
		if old, ok := sMap[userSpeciality.Level]; ok {
			if userSpeciality.TakeAlong {
				old.TakeAlongSpecialityResourceId = resourceID
			} else {
				old.ChoosePool = append(old.ChoosePool, resourceID)
			}
		} else {
			newDto := dto.UserHeroSpeciality{
				Level:                         int32(userSpeciality.Level),
				TakeAlongSpecialityResourceId: 0,
				ChoosePool:                    make([]enum.HeroSpeciality, 0),
			}
			if userSpeciality.TakeAlong {
				newDto.TakeAlongSpecialityResourceId = resourceID
			} else {
				newDto.ChoosePool = append(newDto.ChoosePool, resourceID)
			}
			sMap[userSpeciality.Level] = newDto
		}
	}
	for _, heroSpeciality := range sMap {
		result.Speciality = append(result.Speciality, heroSpeciality)
	}

	var pMap = make(map[int]dto.UserHeroPassive)
	for _, userPassive := range passive {
		resourceID, _ := GetHeroPassiveResourceId(userPassive.PassiveId)
		if old, ok := pMap[userPassive.Level]; ok {
			if userPassive.TakeAlong {
				old.TakeAlongPassiveResourceId = resourceID
			} else {
				old.ChoosePool = append(old.ChoosePool, resourceID)
			}
		} else {
			newDto := dto.UserHeroPassive{
				Level:                      int32(userPassive.Level),
				TakeAlongPassiveResourceId: 0,
				ChoosePool:                 make([]enum.HeroPassive, 0),
			}
			if userPassive.TakeAlong {
				newDto.TakeAlongPassiveResourceId = resourceID
			} else {
				newDto.ChoosePool = append(newDto.ChoosePool, resourceID)
			}
			pMap[userPassive.Level] = newDto
		}
	}
	for _, heroPassive := range pMap {
		result.Passive = append(result.Passive, heroPassive)
	}

	err = session.Commit()
	if err != nil {
		return nil, err
	}
	return result, nil
}
