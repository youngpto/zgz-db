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

	// 获取玩家未解锁的英雄
	heros, err := service.GetAllHero()
	if err != nil {
		return nil, err
	}
H:
	for _, hero := range heros {
		for _, info := range result {
			if int(info.HeroId) == hero.ResourceId {
				continue H
			}
		}
		specialities, err := service.FindAllUserSpecialityByUserAndHero(userId, hero.Id)
		if err != nil {
			return nil, err
		}
		passives, err := service.FindAllUserPassiveByUser(userId, hero.Id)
		if err != nil {
			return nil, err
		}
		result = append(result, &dto.HeroInfo{
			HeroId:            int(hero.ResourceId),
			GainExperienceRes: dto.GainExperienceRes{},
			UserHeroProperty: dto.UserHeroProperty{
				BaseHp:        int32(hero.Life),
				BaseSan:       int32(hero.Reason),
				BaseStrength:  int32(hero.Power),
				BaseAgility:   int32(hero.Agile),
				BaseKnowledge: int32(hero.Knowledge),
				BaseWillPower: int32(hero.Will),
			},
			Speciality: specialities,
			Passive:    passives,
		})
	}
	return result, nil
}
