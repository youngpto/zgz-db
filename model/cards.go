package model

import "time"

// HeroCards 英雄卡牌配置表
type HeroCards struct {
	Id        int64     `xorm:"pk autoincr"`       // 主键
	HeroId    int64     `xorm:"notnull 'hero_id'"` // 英雄ID
	CardId    int64     `xorm:"notnull 'card_id'"` // 卡牌cid
	CreatedAt time.Time `xorm:"created"`           // 创建时间
	UpdatedAt time.Time `xorm:"updated"`           // 修改时间
}

func (h *HeroCards) TableName() string {
	return "hero_cards"
}

// UserHeroCards 用户英雄卡牌解锁情况
type UserHeroCards struct {
	Id             int64     `xorm:"pk autoincr"`       // 主键
	UserId         int64     `xorm:"notnull"`           // 用户Id
	HeroId         int64     `xorm:"notnull 'hero_id'"` // 英雄ID
	HeroCardId     int64     `xorm:"notnull"`           // 英雄卡牌配置表主键
	UnlockQuantity int       `xorm:"notnull"`           // 解锁数量
	CreatedAt      time.Time `xorm:"created"`           // 创建时间
	UpdatedAt      time.Time `xorm:"updated"`           // 修改时间
}

func (u *UserHeroCards) TableName() string {
	return "user_hero_cards"
}

type UserHeroCardUnlockInfo struct {
	HeroCards     `xorm:"extends"` // 卡牌基本信息
	UserHeroCards `xorm:"extends"` // 用户英雄的卡牌解锁信息
}
