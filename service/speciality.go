package service

import (
	"github.com/youngpto/zgz-db/base"
	"github.com/youngpto/zgz-db/model"
	"xorm.io/builder"
)

// DropAllSpeciality 清空专长表
func DropAllSpeciality() {
	_, _ = base.Engine.Where("1=1").Delete(new(model.Speciality))
}

// InsertSpeciality 新增专长
func InsertSpeciality(speciality []*model.Speciality) {
	_, _ = base.Engine.Insert(&speciality)
}

// GetHeroSpecialityId 获取指定英雄专长ID
func GetHeroSpecialityId(heroId, specialityId int) (int64, bool) {
	result := new(model.Speciality)
	_, err := base.Engine.Where(builder.Eq{"hero_id": heroId, "resource_id": specialityId}).
		Get(result)
	if err != nil {
		return 0, false
	}
	return result.Id, true
}

// DropAllSpecialityRule 清空专长规则表
func DropAllSpecialityRule() {
	_, _ = base.Engine.Where("1=1").Delete(new(model.SpecialityRule))
}

// InsertSpecialityRule 新增专长规则
func InsertSpecialityRule(batch []*model.SpecialityRule) {
	_, _ = base.Engine.Insert(&batch)
}

// DropAllUserSpeciality 清空玩家配置的专长表
func DropAllUserSpeciality() {
	_, _ = base.Engine.Where("1=1").Delete(new(model.UserSpeciality))
}

// InsertUserSpeciality 玩家新增专长配置
func InsertUserSpeciality(batch []*model.UserSpeciality) {
	_, _ = base.Engine.Insert(&batch)
}

// FindUserSpecialityByUserAndHero 根据用户和英雄获取专长
func FindUserSpecialityByUserAndHero(userId int64, heroId int64) []model.UserSpeciality {
	var result []model.UserSpeciality
	_ = base.Engine.Where(builder.Eq{"user_id": userId, "hero_id": heroId}).
		Cols("level", "speciality_id").
		Find(&result)
	return result
}
