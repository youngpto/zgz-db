package model

import "time"

// HeroUpgradeReward 玩家升级奖励
type HeroUpgradeReward struct {
	Id          int64     `xorm:"pk autoincr"`    // 主键
	HeroId      int       `xorm:"notnull"`        // 英雄ID
	Rank        int       `xorm:"notnull"`        // 等级
	Category    int       `xorm:"notnull"`        // 奖励分类 1专长 2卡牌 3被动  4自定义属性点
	Type        int       `xorm:"'type'"`         // 特殊分类类型 4自定义属性点(1生命,2理智,3力量,4敏捷,5知识,6意志)
	ObjectId    int       `xorm:"'object_id'"`    // 对象ID(分类为专长 卡牌 被动时使用)
	ObjectValue int       `xorm:"'object_value'"` // 对象值(自定义属性点取值该字段)
	CreatedAt   time.Time `xorm:"created"`        // 创建时间
	UpdatedAt   time.Time `xorm:"updated"`        // 修改时间
}

func (h *HeroUpgradeReward) TableName() string {
	return "hero_upgrade_reward"
}

// UserHeroUpgradeRewardLog 玩家升级奖励领取记录
type UserHeroUpgradeRewardLog struct {
	Id                  int64     `xorm:"pk autoincr"`                      // 主键
	UserId              int64     `xorm:"notnull"`                          // 用户ID
	HeroUpgradeRewardId int64     `xorm:"notnull 'hero_upgrade_reward_id'"` // 英雄升级奖励ID
	CreatedAt           time.Time `xorm:"created"`                          // 创建时间
	UpdatedAt           time.Time `xorm:"updated"`                          // 修改时间
}

func (u *UserHeroUpgradeRewardLog) TableName() string {
	return "user_hero_upgrade_reward_log"
}
