package service

import (
	"errors"
	"fmt"
	"github.com/youngpto/zgz-db/base"
	"github.com/youngpto/zgz-db/enum"
	"github.com/youngpto/zgz-db/model"
	"xorm.io/builder"
)

// FindAllPassive 获取所有被动信息
func FindAllPassive() []model.Passive {
	var result []model.Passive
	_ = base.Engine.Find(&result)
	return result
}

// DropAllPassive 清空被动表
func DropAllPassive() {
	_, _ = base.Engine.Where("1=1").Delete(new(model.Passive))
}

// InsertPassive 新增英雄被动
func InsertPassive(passive []*model.Passive) {
	_, _ = base.Engine.Insert(&passive)
}

// GetPassiveResourceIdByKeyID 根据主键Id解析被动资源ID
func GetPassiveResourceIdByKeyID(pkId ...int64) ([]enum.HeroPassive, error) {
	if len(pkId) == 0 {
		return nil, nil
	}
	session := base.Engine.NewSession()
	defer session.Close()

	if err := session.Begin(); err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	var result []enum.HeroPassive
	for _, id := range pkId {
		passive := model.Passive{}
		_, err := session.ID(id).
			Cols("resource_id").Get(&passive)
		if err != nil {
			return nil, err
		}
		result = append(result, enum.HeroPassive(passive.ResourceId))
	}
	err := session.Commit()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetHeroPassiveId 获取被动主键Id
func GetHeroPassiveId(heroId, resourcePId int) (int64, bool) {
	result := new(model.Passive)
	_, err := base.Engine.Where(builder.Eq{"hero_id": heroId, "resource_id": resourcePId}).
		Get(result)
	if err != nil {
		return 0, false
	}
	return result.Id, true
}

// DropAllPassiveRule 清空被动规则表
func DropAllPassiveRule() {
	_, _ = base.Engine.Where("1=1").Delete(new(model.PassiveRule))
}

// InsertPassiveRule 新增被动规则
func InsertPassiveRule(batch []*model.PassiveRule) {
	_, _ = base.Engine.Insert(batch)
}

// InsertOrUpdateUserPassiveByLevel 插入或更新玩家在确定层级配置的被动
func InsertOrUpdateUserPassiveByLevel(passive *model.UserPassive) error {
	session := base.Engine.NewSession()
	defer session.Close()

	if err := session.Begin(); err != nil {
		fmt.Println(err.Error())
		return err
	}

	// 校验该层级是否能配置当前选择的被动
	var rules []model.PassiveRule
	err := session.Where(builder.Eq{"hero_id": passive.HeroId, "level": passive.Level}).
		Cols("passive_id").
		Find(&rules)
	if err != nil {
		return err
	}

	// 是否匹配上规则限制
	find := false
	for _, rule := range rules {
		if rule.PassiveId == passive.PassiveId {
			find = true
			break
		}
	}
	if !find {
		return errors.New("不允许在当前层级配置的专长")
	}

	// 检查是否已经配置了该层级的专长
	exist, err := session.Table("user_passive").
		Where(builder.Eq{"user_id": passive.UserId,
			"hero_id": passive.HeroId,
			"level":   passive.Level}).
		Exist()
	if err != nil {
		return err
	}

	// 有则修改，无则插入
	if exist {
		_, err := session.Where(builder.Eq{"user_id": passive.UserId,
			"hero_id": passive.HeroId,
			"level":   passive.Level}).
			Cols("passive_id").
			Update(passive)
		if err != nil {
			return err
		}
	} else {
		_, err := session.Insert(passive)
		if err != nil {
			return err
		}
	}
	return session.Commit()
}

// FindUserPassiveByUserAndHero 根据用户和英雄获取被动
func FindUserPassiveByUserAndHero(userId int64, heroId int64) ([]model.UserPassive, error) {
	var result []model.UserPassive
	err := base.Engine.Where(builder.Eq{"user_id": userId, "hero_id": heroId}).
		Cols("level", "passive_id").
		Find(&result)
	return result, err
}
