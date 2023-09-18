package model

import "time"

// Passive 英雄被动设置
type Passive struct {
	Id          int64     `xorm:"pk autoincr"`                         // 主键
	ResourceId  int       `xorm:"notnull 'resource_id'"`               // 资源ID
	HeroId      int       `xorm:"notnull 'hero_id'"`                   // 英雄ID
	Name        string    `xorm:"varchar(191) notnull"`                // 被动名称
	Desc        string    `xorm:"varchar(191) notnull"`                // 被动描述
	EnglishName string    `xorm:"varchar(191) notnull 'english_name'"` // 英文被动名称
	EnglishDesc string    `xorm:"varchar(191) notnull 'english_desc'"` // 英文被动描述
	CreatedAt   time.Time `xorm:"created"`                             // 创建时间
	UpdatedAt   time.Time `xorm:"updated"`                             // 修改时间
}

func (p *Passive) TableName() string {
	return "passive"
}

// PassiveRule 英雄被动规则
type PassiveRule struct {
	Id        int64     `xorm:"pk autoincr"` // 主键
	HeroId    int       `xorm:"notnull"`     // 英雄Id
	Level     int       `xorm:"notnull"`     // 层级
	PassiveId int64     `xorm:"notnull"`     // 被动主键ID
	CreatedAt time.Time `xorm:"created"`     // 创建时间
	UpdatedAt time.Time `xorm:"updated"`     // 修改时间
}

func (p *PassiveRule) TableName() string {
	return "passive_rule"
}

type UserPassive struct {
	Id        int64     `xorm:"pk autoincr"`                         // 主键
	UserId    int64     `xorm:"notnull 'user_id'"`                   // 用户ID
	HeroId    int64     `xorm:"notnull"`                             // 调查员编号
	Level     int       `xorm:"notnull"`                             // 层级
	PassiveId int64     `xorm:"notnull 'passive_id'"`                // 被动ID
	TakeAlong bool      `xorm:"default(false) notnull 'take_along'"` // 是否佩戴
	CreatedAt time.Time `xorm:"created"`                             // 创建时间
	UpdatedAt time.Time `xorm:"updated"`                             // 修改时间
}

func (u *UserPassive) TableName() string {
	return "user_passive"
}

type PassiveRuleInfo struct {
	PassiveRule `xorm:"extends"` // 被动规则配置
	Passive     `xorm:"extends"` // 被动基本信息
}

type UserHeroPassiveInfo struct {
	UserPassive `xorm:"extends"` // 用户被动配置
	Passive     `xorm:"extends"` // 被动基本信息
}
