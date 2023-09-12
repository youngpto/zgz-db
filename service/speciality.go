package service

import (
	"errors"
	"fmt"
	"github.com/youngpto/zgz-db/base"
	"github.com/youngpto/zgz-db/enum"
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

// GetSpecialityResourceIdByKeyID 根据主键Id解析资源ID
func GetSpecialityResourceIdByKeyID(pkId ...int64) ([]enum.HeroSpeciality, error) {
	if len(pkId) == 0 {
		return nil, nil
	}
	session := base.Engine.NewSession()
	defer session.Close()

	if err := session.Begin(); err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	var result []enum.HeroSpeciality
	for _, id := range pkId {
		speciality := model.Speciality{}
		_, err := session.ID(id).
			Cols("resource_id").Get(&speciality)
		if err != nil {
			return nil, err
		}
		result = append(result, enum.HeroSpeciality(speciality.ResourceId))
	}
	err := session.Commit()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetHeroSpecialityId 获取指定英雄专长ID
func GetHeroSpecialityId(heroId, resourceSId int) (int64, bool) {
	result := new(model.Speciality)
	_, err := base.Engine.Where(builder.Eq{"hero_id": heroId, "resource_id": resourceSId}).
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

// InsertOrUpdateUserSpecialityByLevel 插入或更新玩家在确定层级配置的专长
func InsertOrUpdateUserSpecialityByLevel(speciality *model.UserSpeciality) error {
	session := base.Engine.NewSession()
	defer session.Close()

	if err := session.Begin(); err != nil {
		fmt.Println(err.Error())
		return err
	}

	// 校验该层级是否能配置当前选择的专长
	var rules []model.SpecialityRule
	err := session.Where(builder.Eq{"hero_id": speciality.HeroId, "level": speciality.Level}).
		Cols("speciality_id").
		Find(&rules)
	if err != nil {
		return err
	}

	// 是否匹配上规则限制
	find := false
	for _, rule := range rules {
		if rule.SpecialityId == speciality.SpecialityId {
			find = true
			break
		}
	}
	if !find {
		return errors.New("不允许在当前层级配置的专长")
	}

	// 检查是否已经配置了该层级的专长
	exist, err := session.Table("user_speciality").
		Where(builder.Eq{"user_id": speciality.UserId,
			"hero_id": speciality.HeroId,
			"level":   speciality.Level}).
		Exist()
	if err != nil {
		return err
	}

	// 有则修改，无则插入
	if exist {
		_, err := session.Where(builder.Eq{"user_id": speciality.UserId,
			"hero_id": speciality.HeroId,
			"level":   speciality.Level}).
			Cols("speciality_id").
			Update(speciality)
		if err != nil {
			return err
		}
	} else {
		_, err := session.Insert(speciality)
		if err != nil {
			return err
		}
	}
	return session.Commit()
}

// FindUserSpecialityByUserAndHero 根据用户和英雄获取专长
func FindUserSpecialityByUserAndHero(userId int64, heroId int64) ([]model.UserSpeciality, error) {
	var result []model.UserSpeciality
	err := base.Engine.Where(builder.Eq{"user_id": userId, "hero_id": heroId}).
		Cols("level", "speciality_id").
		Find(&result)
	return result, err
}
