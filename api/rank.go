package api

import (
	"errors"
	"github.com/youngpto/zgz-db/dto"
	"github.com/youngpto/zgz-db/model"
	"github.com/youngpto/zgz-db/service"
)

// HeroGainExperience 英雄获得经验
func HeroGainExperience(userId int64, resourceHeroId int64, experience float64) (*dto.GainExperienceRes, chan int, error) {
	// 解析英雄ID
	heroId, err := service.GetHeroId(resourceHeroId)
	if err != nil {
		return nil, nil, errors.New("未找到调查员类型")
	}
	// 获取玩家英雄等级和累计经验池
	userHeroRankInfo, err := service.GetUserHeroRankInfo(userId, heroId)
	if err != nil {
		return nil, nil, err
	}
	currentLevel := userHeroRankInfo.Rank
	currentExp := userHeroRankInfo.ExperiencePool
	var addLevel int   // 增长的等级
	var addExp float64 // 增长的经验
	result := new(dto.GainExperienceRes)
	result.OriginLevel = currentLevel
	result.CurrentLevel = currentLevel
	result.OriginExp = currentExp
	result.CurrentExp = currentExp

	// 获取升级路线
	route, err := service.FindRankGTELevel(currentLevel)
	if err != nil {
		return result, nil, nil
	}
	var upgradeRoute = make(map[int]model.Rank)
	for _, rank := range route {
		upgradeRoute[rank.Rank] = rank
	}
	if rank, ok := upgradeRoute[currentLevel]; ok {
		result.OriginMaxExp = rank.Experience
		result.CurrentMaxExp = rank.Experience
	}

	// 计算经验值
	residual := experience
	for residual > 0 {
		// 检查当前等级累计经验还差多少
		if rank, ok := upgradeRoute[currentLevel]; ok {
			// 升级经验为0代表满级
			if rank.Experience == 0 {
				break
			}
			// 差多少经验升级
			differenceExp := rank.Experience - currentExp
			// 足够升级
			if residual >= differenceExp {
				currentExp = 0
				residual -= differenceExp
				// 检查是否有下一个等级
				if _, ok := upgradeRoute[currentLevel+1]; ok {
					// 有则升级
					currentLevel++
					addLevel++
				} else {
					// 没有配置下一个等级
					break
				}
				addExp += differenceExp
			} else {
				// 不足够升级
				addExp += residual
				currentExp += residual
				residual = 0
			}
		} else {
			// 找不到当前等级的配置，直接退出
			break
		}
	}

	// 如果本次产生了升级
	if addLevel > 0 || addExp > 0 {
		err := service.UpdateUserRank(&model.UserHeroRank{
			UserId:         userHeroRankInfo.UserId,
			HeroId:         userHeroRankInfo.HeroId,
			Rank:           currentLevel,
			ExperiencePool: currentExp,
		})
		if err != nil {
			return result, nil, nil
		}
		result.CurrentLevel = currentLevel
		result.CurrentExp = currentExp
		if rank, ok := upgradeRoute[currentLevel]; ok {
			result.CurrentMaxExp = rank.Experience
		}
	}

	ch := make(chan int, 10)
	// 后台自动领取升级奖励
	go func() {
		defer close(ch)
		unclaimedRankReward, err := service.FindUnclaimedRankReward(userHeroRankInfo.UserId, int(userHeroRankInfo.HeroId), currentLevel)
		if err != nil {
			return
		}
		rewardType := make(map[int]struct{})
		for _, reward := range unclaimedRankReward {
			rewardType[reward.Category] = struct{}{}
		}
		_ = service.ReceiveReward(userId, unclaimedRankReward)
		for rt, _ := range rewardType {
			ch <- rt
		}
	}()

	return result, ch, nil
}

// GetUnclaimedRankReward 领取未领取的升级奖励
func GetUnclaimedRankReward(userId int64, resourceHeroId int64, currenRank int) error {
	heroId, err := service.GetHeroId(resourceHeroId)
	if err != nil {
		return errors.New("未找到调查员类型")
	}
	// 后台自动领取升级奖励
	go func() {
		unclaimedRankReward, err := service.FindUnclaimedRankReward(userId, int(heroId), currenRank)
		if err != nil {
			return
		}
		_ = service.ReceiveReward(userId, unclaimedRankReward)
	}()
	return nil
}
