package service

import (
	"fmt"
	"github.com/youngpto/zgz-db/base"
	"github.com/youngpto/zgz-db/model"
	"xorm.io/builder"
)

// FindUnclaimedRankReward 获取未领取的升级奖励
func FindUnclaimedRankReward(userId int64, webHeroId int, currentRank int) ([]model.HeroUpgradeReward, error) {
	session := base.Engine.NewSession()
	defer session.Close()

	if err := session.Begin(); err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	var claimed []model.UserHeroUpgradeRewardLog
	err := session.Where(builder.Eq{"user_id": userId}).
		Cols("hero_upgrade_reward_id").
		Find(&claimed)
	if err != nil {
		return nil, err
	}
	var claimedIds []int64
	for _, log := range claimed {
		claimedIds = append(claimedIds, log.HeroUpgradeRewardId)
	}

	var result []model.HeroUpgradeReward
	err = session.Where(builder.Eq{"hero_id": webHeroId}.
		And(builder.Lte{"rank": currentRank}.
			And(builder.NotIn("id", claimedIds)))).
		Cols("id", "category", "type", "object_id", "object_value").Find(result)
	if err != nil {
		return nil, err
	}

	err = session.Commit()
	if err != nil {
		return nil, err
	}

	return result, nil
}
