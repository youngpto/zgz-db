package model

import "time"

// HeroProperty 英雄设置的默认属性
type HeroProperty struct {
	Id         int64     `xorm:"pk autoincr"`   // 主键
	ResourceId int       `xorm:"'resource_id'"` // 资源ID
	Life       int       // 基础生命
	Reason     int       // 基础理智
	Power      int       // 基础力量
	Agile      int       // 基础敏捷
	Knowledge  int       // 基础知识
	Will       int       // 基础意志
	CreatedAt  time.Time `xorm:"created"` // 创建时间
	UpdatedAt  time.Time `xorm:"updated"` // 修改时间
}
