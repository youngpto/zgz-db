package service

import (
	"fmt"
	"xorm.io/builder"
	"zgz-db/base"
	"zgz-db/model"
)

// FindPropertyOffsetByUserAndHero 查询玩家对应英雄的六维偏移值
func FindPropertyOffsetByUserAndHero(userId int64, heroId int64) (*model.PropertyOffset, bool) {
	result := new(model.PropertyOffset)
	ok, err := base.Engine.Table("user_hero").
		Where(builder.Eq{"user_id": userId, "hero_id": heroId}).
		Cols("id", "total_offset", "life_offset", "reason_offset",
			"power_offset", "agile_offset", "knowledge_offset", "will_offset").
		Get(result)
	if err != nil {
		fmt.Println(err.Error())
	}
	return result, ok
}

// FindHeroBaseProperty 查询英雄的基础六维属性
func FindHeroBaseProperty(heroId int64) (*model.HeroProperty, bool) {
	result := new(model.HeroProperty)
	ok, err := base.Engine.Table("hero").
		Where(builder.Eq{"resource_id": heroId}).
		Cols("id", "resource_id", "life", "reason", "power", "agile", "knowledge", "will").
		Get(result)
	if err != nil {
		fmt.Println(err.Error())
	}
	return result, ok
}

// UpdateHeroBaseProperty 更新英雄的基础六维属性
func UpdateHeroBaseProperty(heroId int64, property *model.HeroProperty) bool {
	ok, err := base.Engine.Table("hero").
		Where(builder.Eq{"resource_id": heroId}).
		Update(property)
	if err != nil {
		fmt.Println(err)
	}
	return ok > 0
}
