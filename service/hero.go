package service

import (
	"github.com/youngpto/zgz-db/base"
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
