package service

import (
	"errors"
	"github.com/go-xorm/xorm"
	"github.com/youngpto/zgz-db/base"
	"github.com/youngpto/zgz-db/enum"
	"github.com/youngpto/zgz-db/model"
	"xorm.io/builder"
)

// FindPropertyOffsetByUserAndHero 查询玩家对应英雄的六维偏移值
func FindPropertyOffsetByUserAndHero(userId int64, heroId int64) (*model.PropertyOffset, error) {
	result := new(model.PropertyOffset)
	_, err := base.Engine.Table("user_hero").
		Where(builder.Eq{"user_id": userId, "hero_id": heroId}).
		Cols("id", "total_offset", "life_offset", "reason_offset",
			"power_offset", "agile_offset", "knowledge_offset", "will_offset").
		Get(result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateTotalOffsetToUserHero 修改玩家英雄对应的总偏移值
func UpdateTotalOffsetToUserHero(session *xorm.Session, userId int64, heroId int64, addNumber int) error {
	// 获取已经存储的总偏移值
	userHeroPropertyOffset := new(model.PropertyOffset)
	_, err := session.Table("user_hero").
		Where(builder.Eq{"user_id": userId, "hero_id": heroId}).
		Cols("id", "total_offset").
		Get(userHeroPropertyOffset)
	if err != nil {
		return err
	}

	// 修改总值
	userHeroPropertyOffset.TotalOffset += addNumber
	_, err = session.Table("user_hero").
		ID(userHeroPropertyOffset.Id).
		Update(userHeroPropertyOffset)
	if err != nil {
		return err
	}
	return nil
}

// UpdatePropertyOffsetByType 更新玩家英雄六维属性的某一项
func UpdatePropertyOffsetByType(userId, heroId int64, property enum.HeroProperty, isIncr bool) (*model.PropertyOffset, error) {
	session := base.Engine.NewSession()
	defer session.Close()

	if err := session.Begin(); err != nil {
		return nil, err
	}

	// 获取已经存储的六维属性偏移值
	userHeroPropertyOffset := new(model.PropertyOffset)
	_, err := session.Table("user_hero").
		Where(builder.Eq{"user_id": userId, "hero_id": heroId}).
		Cols("id", "total_offset", "life_offset", "reason_offset",
			"power_offset", "agile_offset", "knowledge_offset", "will_offset").
		Get(userHeroPropertyOffset)
	if err != nil {
		return nil, err
	}

	number, ok := userHeroPropertyOffset.IncrOrDecrByType(property, isIncr)
	if !ok {
		return nil, errors.New("不合法的点数")
	}

	str := model.ConvEnum2Str(property)
	if str == "" {
		return nil, errors.New("不合法的参数项")
	}
	_, err = session.Table("user_hero").
		ID(userHeroPropertyOffset.Id).
		Update(map[string]interface{}{str: number})
	if err != nil {
		return nil, err
	}
	userHeroPropertyOffset.ResetPropertyByType(property, number)
	err = session.Commit()
	if err != nil {
		return nil, err
	}
	return userHeroPropertyOffset, nil
}

// ResetPropertyOffset 重置玩家属性加点
func ResetPropertyOffset(userId, heroId int64) error {
	_, err := base.Engine.Table("user_hero").
		Where(builder.Eq{"user_id": userId, "hero_id": heroId}).
		MustCols("life_offset", "reason_offset", "power_offset",
			"agile_offset", "knowledge_offset", "will_offset").
		Update(new(model.PropertyOffset))
	if err != nil {
		return err
	}
	return nil
}

// FindHeroBaseProperty 查询英雄的基础六维属性
func FindHeroBaseProperty(resourceHeroId int64) (*model.HeroProperty, error) {
	result := new(model.HeroProperty)
	_, err := base.Engine.Table("hero").
		Where(builder.Eq{"resource_id": resourceHeroId}).
		Cols("id", "resource_id", "life", "reason", "power", "agile", "knowledge", "will").
		Get(result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateHeroBaseProperty 更新英雄的基础六维属性
func UpdateHeroBaseProperty(resourceHeroId int64, property *model.HeroProperty) error {
	_, err := base.Engine.Table("hero").
		Where(builder.Eq{"resource_id": resourceHeroId}).
		Update(property)
	if err != nil {
		return err
	}
	return nil
}
