package service

import (
	"errors"
	"fmt"
	"github.com/youngpto/zgz-db/base"
	"github.com/youngpto/zgz-db/enum"
	"github.com/youngpto/zgz-db/model"
	"xorm.io/builder"
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

// UpdateTotalOffsetToUserHero 修改玩家英雄对应的总偏移值
func UpdateTotalOffsetToUserHero(userId int64, heroId int64, addNumber int) error {
	session := base.Engine.NewSession()
	defer session.Close()

	if err := session.Begin(); err != nil {
		fmt.Println(err.Error())
		return err
	}

	// 获取已经存储的总偏移值
	userHeroPropertyOffset := new(model.PropertyOffset)
	_, err := session.Table("user_hero").
		Where(builder.Eq{"user_id": userId, "hero_id": heroId}).
		Cols("id", "total_offset").
		Get(userHeroPropertyOffset)
	if err != nil {
		fmt.Println(err.Error())
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
	return session.Commit()
}

// UpdatePropertyOffsetByType 更新玩家英雄六维属性的某一项
func UpdatePropertyOffsetByType(userId, heroId int64, property enum.HeroProperty, isIncr bool) (*model.PropertyOffset, error) {
	session := base.Engine.NewSession()
	defer session.Close()

	if err := session.Begin(); err != nil {
		fmt.Println(err.Error())
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
		fmt.Println(err.Error())
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
func FindHeroBaseProperty(resourceHeroId int64) (*model.HeroProperty, bool) {
	result := new(model.HeroProperty)
	ok, err := base.Engine.Table("hero").
		Where(builder.Eq{"resource_id": resourceHeroId}).
		Cols("id", "resource_id", "life", "reason", "power", "agile", "knowledge", "will").
		Get(result)
	if err != nil {
		fmt.Println(err.Error())
	}
	return result, ok
}

// UpdateHeroBaseProperty 更新英雄的基础六维属性
func UpdateHeroBaseProperty(resourceHeroId int64, property *model.HeroProperty) bool {
	ok, err := base.Engine.Table("hero").
		Where(builder.Eq{"resource_id": resourceHeroId}).
		Update(property)
	if err != nil {
		fmt.Println(err)
	}
	return ok > 0
}
