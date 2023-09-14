package service

import (
	"github.com/go-xorm/xorm"
	"github.com/youngpto/zgz-db/base"
	"github.com/youngpto/zgz-db/model"
	"xorm.io/builder"
)

func DropAllRank() {
	_, _ = base.Engine.Where("1=1").Delete(new(model.Rank))
}

func InsertRank(rank []*model.Rank) {
	_, _ = base.Engine.Insert(&rank)
}

// FindRankGTELevel 获取大于等于当前等级的升级路线
func FindRankGTELevel(currentLevel int) ([]model.Rank, error) {
	var result []model.Rank
	err := base.Engine.Where(builder.Gte{"rank": currentLevel}).
		Cols("rank", "experience").
		Find(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateUserRank 更新玩家英雄等级
func UpdateUserRank(rank *model.UserHeroRank) error {
	_, err := base.Engine.ID(rank.Id).Update(rank)
	if err != nil {
		return err
	}
	return nil
}

// FindUnclaimedRankReward 获取未领取的升级奖励
func FindUnclaimedRankReward(userId int64, webHeroId int, currentRank int) ([]model.HeroUpgradeReward, error) {
	session := base.Engine.NewSession()
	defer session.Close()

	if err := session.Begin(); err != nil {
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

// ReceiveReward 领取奖励
func ReceiveReward(userId int64, rewards []model.HeroUpgradeReward) error {
	for _, reward := range rewards {
		_ = receiveReward(userId, reward)
	}
	return nil
}

// 领取奖励内部方法
func receiveReward(userId int64, reward model.HeroUpgradeReward) error {
	session := base.Engine.NewSession()
	defer session.Close()

	if err := session.Begin(); err != nil {
		return err
	}

	switch reward.Type {
	case 1: // 专长
		// 根据英雄Id和专长资源ID获取数据库中的专长
		speciality := new(model.Speciality)
		_, err := base.Engine.
			Where(builder.Eq{
				"hero_id":     reward.HeroId,
				"resource_id": reward.ObjectId,
			}).
			Get(speciality)
		if err != nil {
			return err
		}
		// 给玩家添加对应的专长，且设置默认值
		_, err = InsertNotExistUserSpeciality(session, &model.UserSpeciality{
			UserId:       userId,
			HeroId:       int64(reward.HeroId),
			Level:        reward.ObjectValue,
			SpecialityId: speciality.Id,
			TakeAlong:    false,
		})
		if err != nil {
			return err
		}
	case 2: // 卡牌[暂无]
		return nil
	case 3: // 被动
		// 根据英雄Id和被动资源ID获取数据库中的被动
		passive := new(model.Passive)
		_, err := base.Engine.
			Where(builder.Eq{
				"hero_id":     reward.HeroId,
				"resource_id": reward.ObjectId,
			}).
			Get(passive)
		if err != nil {
			return err
		}
		// 给玩家添加对应的专长，且设置默认值
		_, err = InsertNotExistUserPassive(session, &model.UserPassive{
			UserId:    userId,
			HeroId:    int64(reward.HeroId),
			Level:     reward.ObjectValue,
			PassiveId: passive.Id,
			TakeAlong: false,
		})
		if err != nil {
			return err
		}
	case 4: // 自定义属性点
		err := UpdateTotalOffsetToUserHero(session, userId, int64(reward.HeroId), reward.ObjectValue)
		if err != nil {
			return err
		}
	}
	// 添加一条奖励领取记录
	_ = InsertUserHeroUpgradeRewardLog(session, &model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: reward.Id,
	})
	return session.Commit()
}

func DropAllHeroUpgradeReward() {
	_, _ = base.Engine.Where("1=1").Delete(new(model.HeroUpgradeReward))
}

func InsertHeroUpgradeReward(heroUpgradeReward []*model.HeroUpgradeReward) {
	_, _ = base.Engine.Insert(&heroUpgradeReward)
}

func InsertUserHeroUpgradeRewardLog(session *xorm.Session, log *model.UserHeroUpgradeRewardLog) error {
	_, err := session.Where(
		builder.NotExists(builder.
			Select("id").
			From("user_hero_upgrade_reward_log").
			Where(builder.Eq{
				"user_id":                log.UserId,
				"hero_upgrade_reward_id": log.HeroUpgradeRewardId},
			)),
	).Insert(log)
	if err != nil {
		return err
	}
	return nil
}
