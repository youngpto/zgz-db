package model

import "time"

type User struct {
	Id        int64     `xorm:"pk autoincr"` // 主键
	Name      string    `xorm:"notnull"`     // 用户名
	CreatedAt time.Time `xorm:"created"`     // 创建时间
	UpdatedAt time.Time `xorm:"updated"`     // 修改时间
}

func (u *User) TableName() string {
	return "users"
}
