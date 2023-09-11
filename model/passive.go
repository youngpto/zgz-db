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
