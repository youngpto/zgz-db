package model

import "time"

type User struct {
	Id            int64     `xorm:"pk autoincr"`       // 主键
	Name          string    `xorm:"notnull"`           // 用户名
	TmpAchieveNum int       `xorm:"'tmp_achieve_num'"` // 统计steam成就数量
	CreatedAt     time.Time `xorm:"created"`           // 创建时间
	UpdatedAt     time.Time `xorm:"updated"`           // 修改时间
}

func (u *User) TableName() string {
	return "users"
}

// UserHeroRank 玩家英雄等级
type UserHeroRank struct {
	Id             int64     `xorm:"pk autoincr"`               // 主键
	UserId         int64     `xorm:"notnull 'user_id'"`         // 用户ID
	HeroId         int64     `xorm:"notnull 'hero_id'"`         // web端英雄Id
	Rank           int       `xorm:"'rank'"`                    // 英雄等级
	ExperiencePool float64   `xorm:"notnull 'experience_pool'"` // 经验池
	CreatedAt      time.Time `xorm:"created"`                   // 创建时间
	UpdatedAt      time.Time `xorm:"updated"`                   // 修改时间
}

// UserHeroRankAndPropertyOffset 玩家英雄等级和配置的属性偏移值
type UserHeroRankAndPropertyOffset struct {
	Id              int64     `xorm:"pk autoincr"`               // 主键
	UserId          int64     `xorm:"notnull 'user_id'"`         // 用户ID
	HeroId          int64     `xorm:"notnull 'hero_id'"`         // web端英雄Id
	Rank            int       `xorm:"'rank'"`                    // 英雄等级
	ExperiencePool  float64   `xorm:"notnull 'experience_pool'"` // 经验池
	TotalOffset     int       `xorm:"'total_offset'"`            // 总偏移量
	LifeOffset      int       `xorm:"'life_offset'"`             // 生命偏移
	ReasonOffset    int       `xorm:"'reason_offset'"`           // 理智偏移
	PowerOffset     int       `xorm:"'power_offset'"`            // 力量偏移
	AgileOffset     int       `xorm:"'agile_offset'"`            // 敏捷偏移
	KnowledgeOffset int       `xorm:"'knowledge_offset'"`        // 知识偏移
	WillOffset      int       `xorm:"'will_offset'"`             // 意志偏移
	CreatedAt       time.Time `xorm:"created"`                   // 创建时间
	UpdatedAt       time.Time `xorm:"updated"`                   // 修改时间
}
