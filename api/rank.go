package api

import (
	"github.com/youngpto/zgz-db/model"
	"github.com/youngpto/zgz-db/service"
)

// FindUnclaimedRankReward 获取未领取的升级奖励
func FindUnclaimedRankReward(userId int64, webHeroId int64, currenRank int) ([]model.HeroUpgradeReward, error) {
	return service.FindUnclaimedRankReward(userId, int(webHeroId), currenRank)
}
