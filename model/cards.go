package model

import "time"

// UserHeroCards 用户英雄卡牌解锁情况
type UserHeroCards struct {
	Id             int64     `xorm:"pk autoincr"` // 主键
	UserId         int64     `xorm:"notnull"`     // 用户Id
	CardId         int64     `xorm:"notnull"`     // cardsOrderID
	UnlockQuantity int       `xorm:"notnull"`     // 解锁数量
	CreatedAt      time.Time `xorm:"created"`     // 创建时间
	UpdatedAt      time.Time `xorm:"updated"`     // 修改时间
}

func (u *UserHeroCards) TableName() string {
	return "user_hero_cards"
}
