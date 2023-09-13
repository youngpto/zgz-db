package model

import "time"

// Speciality 专长
type Speciality struct {
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

func (s *Speciality) TableName() string {
	return "speciality"
}

// SpecialityRule 专长规则
type SpecialityRule struct {
	Id           int64     `xorm:"pk autoincr"`             // 主键
	HeroId       int       `xorm:"notnull 'hero_id'"`       // 英雄ID
	Level        int       `xorm:"notnull"`                 // 层级
	SpecialityId int64     `xorm:"notnull 'speciality_id'"` // 专长ID
	CreatedAt    time.Time `xorm:"created"`                 // 创建时间
	UpdatedAt    time.Time `xorm:"updated"`                 // 修改时间
}

func (s *SpecialityRule) TableName() string {
	return "speciality_rule"
}

// UserSpeciality 玩家配置的专长
type UserSpeciality struct {
	Id           int64     `xorm:"pk autoincr"`                         // 主键
	UserId       int64     `xorm:"notnull 'user_id'"`                   // 用户ID
	HeroId       int64     `xorm:"notnull"`                             // 调查员编号
	Level        int       `xorm:"notnull"`                             // 层级
	SpecialityId int64     `xorm:"notnull 'speciality_id'"`             // 专长ID
	TakeAlong    bool      `xorm:"default(false) notnull 'take_along'"` // 是否佩戴
	CreatedAt    time.Time `xorm:"created"`                             // 创建时间
	UpdatedAt    time.Time `xorm:"updated"`                             // 修改时间
}

func (u *UserSpeciality) TableName() string {
	return "user_speciality"
}
