package service

import (
	"zgz-db/base"
	"zgz-db/model"
)

func FindAllPassive() []model.Passive {
	var result []model.Passive
	_ = base.Engine.Find(&result)
	return result
}

func InsertPassive(passive []*model.Passive) {
	base.Engine.Insert(&passive)
}
