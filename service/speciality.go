package service

import (
	"errors"
	"github.com/go-xorm/xorm"
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

// GetHeroSpecialityId 获取指定英雄专长ID
func GetHeroSpecialityId(heroId, resourceSId int) (int64, error) {
	result := new(model.Speciality)
	_, err := base.Engine.Where(builder.Eq{"hero_id": heroId, "resource_id": resourceSId}).
		Get(result)
	if err != nil {
		return 0, err
	}
	return result.Id, nil
}

// GetHeroSpecialityResourceId 根据专长ID解析资源ID
func GetHeroSpecialityResourceId(pkID int64) (enum.HeroSpeciality, error) {
	result := new(model.Speciality)
	_, err := base.Engine.ID(pkID).
		Get(result)
	if err != nil {
		return 0, err
	}
	return enum.HeroSpeciality(result.ResourceId), nil
}

// DropAllSpecialityRule 清空专长规则表
func DropAllSpecialityRule() {
	_, _ = base.Engine.Where("1=1").Delete(new(model.SpecialityRule))
}

// InsertSpecialityRule 新增专长规则
func InsertSpecialityRule(batch []*model.SpecialityRule) {
	_, _ = base.Engine.Insert(&batch)
}

// DropAllUserSpeciality 清空玩家的专长表
func DropAllUserSpeciality() {
	_, _ = base.Engine.Where("1=1").Delete(new(model.UserSpeciality))
}

// InsertUserSpeciality 玩家新增专长
func InsertUserSpeciality(batch []*model.UserSpeciality) {
	_, _ = base.Engine.Insert(&batch)
}

// InsertNotExistUserSpeciality 不重复插入玩家专长
func InsertNotExistUserSpeciality(session *xorm.Session, speciality *model.UserSpeciality) (bool, error) {
	insert, err := session.Where(
		builder.NotExists(builder.
			Select("id").
			From("user_speciality").
			Where(builder.Eq{
				"user_id":       speciality.UserId,
				"level":         speciality.Level,
				"hero_id":       speciality.HeroId,
				"speciality_id": speciality.SpecialityId},
			)),
	).Insert(speciality)
	if err != nil {
		return false, err
	}
	return insert > 0, nil
}

// UpdateTakeAlongFormUserSpecialityByLevel 调整玩家在指定层级选中的专长
func UpdateTakeAlongFormUserSpecialityByLevel(speciality *model.UserSpeciality) error {
	session := base.Engine.NewSession()
	defer session.Close()

	if err := session.Begin(); err != nil {
		return err
	}
	// 校验该专长是否符合配置规则
	matchRule, err := session.Where(builder.Eq{
		"hero_id":       speciality.HeroId,
		"level":         speciality.Level,
		"speciality_id": speciality.SpecialityId,
	}).Exist(new(model.SpecialityRule))
	if err != nil {
		return err
	}
	if !matchRule {
		return errors.New("不符合配置规则")
	}
	// 检查是否已经配置了专长
	exist, err := session.Table("user_speciality").
		Where(builder.Eq{
			"user_id":    speciality.UserId,
			"hero_id":    speciality.HeroId,
			"level":      speciality.Level,
			"take_along": true,
		}).Exist()
	if err != nil {
		return err
	}
	// 如果存在的话则取消
	if exist {
		_, err := session.
			Where(builder.Eq{
				"user_id":    speciality.UserId,
				"hero_id":    speciality.HeroId,
				"level":      speciality.Level,
				"take_along": true,
			}).MustCols("take_along").
			Update(&model.UserSpeciality{
				TakeAlong: false,
			})
		if err != nil {
			return err
		}

	}
	_, err = session.
		Where(builder.Eq{
			"user_id":       speciality.UserId,
			"hero_id":       speciality.HeroId,
			"level":         speciality.Level,
			"speciality_id": speciality.SpecialityId,
		}).MustCols("take_along").
		Update(&model.UserSpeciality{
			TakeAlong: true,
		})
	if err != nil {
		return err
	}
	return session.Commit()
}

// FindAllUserSpecialityByUserAndHero 获取玩家所有专长
func FindAllUserSpecialityByUserAndHero(userId int64, heroId int64) ([]model.UserSpeciality, error) {
	var result []model.UserSpeciality
	err := base.Engine.Where(builder.Eq{
		"user_id": userId,
		"hero_id": heroId,
	}).Cols("level", "speciality_id", "take_along").
		Find(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
