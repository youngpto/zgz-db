package service

import (
	"errors"
	"github.com/go-xorm/xorm"
	"github.com/youngpto/zgz-db/base"
	"github.com/youngpto/zgz-db/dto"
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
	// 默认配置一个专长
	exist, _ := session.Where(builder.Eq{
		"user_id":    speciality.UserId,
		"level":      speciality.Level,
		"hero_id":    speciality.HeroId,
		"take_along": true,
	}).Exist(new(model.UserSpeciality))
	if !exist {
		speciality.TakeAlong = true
	}
	exist, err := session.Where(builder.Eq{
		"user_id":       speciality.UserId,
		"level":         speciality.Level,
		"hero_id":       speciality.HeroId,
		"speciality_id": speciality.SpecialityId},
	).Exist(new(model.UserSpeciality))
	if err != nil {
		return false, err
	}
	if !exist {
		insert, err := session.Insert(speciality)
		if err != nil {
			return false, err
		}
		return insert > 0, nil
	}
	return false, nil
}

// CancelTakeAlongUserSpecialityByLevel 取消玩家在对应层级设置的专长
func CancelTakeAlongUserSpecialityByLevel(speciality *model.UserSpeciality) error {
	_, err := base.Engine.
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
	return nil
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
	} else {
		_, err = session.Insert(&model.UserSpeciality{
			UserId:       speciality.UserId,
			HeroId:       speciality.HeroId,
			Level:        speciality.Level,
			SpecialityId: speciality.SpecialityId,
			TakeAlong:    true,
		})
		if err != nil {
			return err
		}
	}
	if err != nil {
		return err
	}
	return session.Commit()
}

// FindAllUserSpecialityByUserAndHero 获取玩家所有专长
func FindAllUserSpecialityByUserAndHero(userId int64, heroId int64) ([]*dto.UserHeroSpeciality, error) {
	tempMap := make(map[int]*dto.UserHeroSpeciality)

	var userHeroSpecialityInfoList []model.UserHeroSpecialityInfo
	err := base.Engine.Table("user_speciality").Alias("us").
		Join("LEFT", []string{"speciality", "s"}, "us.speciality_id = s.id").
		Where("us.user_id = ? AND us.hero_id = ? AND us.take_along = 1", userId, heroId).
		Find(&userHeroSpecialityInfoList)
	if err != nil {
		return nil, err
	}
	for _, speciality := range userHeroSpecialityInfoList {
		tempMap[speciality.Level] = &dto.UserHeroSpeciality{
			Level:                         int32(speciality.Level),
			TakeAlongSpecialityResourceId: enum.HeroSpeciality(speciality.ResourceId),
			ChoosePool:                    make([]enum.HeroSpeciality, 0),
		}
	}

	var userHeroSpecialityRuleInfoList []model.SpecialityRuleInfo
	err = base.Engine.Table("speciality_rule").Alias("sr").
		Join("LEFT", []string{"speciality", "s"}, "sr.speciality_id = s.id").
		Where("sr.hero_id = ?", heroId).
		Find(&userHeroSpecialityRuleInfoList)
	if err != nil {
		return nil, err
	}
	for _, specialityRuleInfo := range userHeroSpecialityRuleInfoList {
		if old, ok := tempMap[specialityRuleInfo.Level]; ok {
			old.ChoosePool = append(old.ChoosePool, enum.HeroSpeciality(specialityRuleInfo.ResourceId))
		} else {
			newDto := &dto.UserHeroSpeciality{
				Level:                         int32(specialityRuleInfo.Level),
				TakeAlongSpecialityResourceId: -1,
				ChoosePool:                    make([]enum.HeroSpeciality, 0),
			}
			newDto.ChoosePool = append(newDto.ChoosePool, enum.HeroSpeciality(specialityRuleInfo.ResourceId))
			tempMap[specialityRuleInfo.Level] = newDto
		}
	}

	var result []*dto.UserHeroSpeciality
	for _, speciality := range tempMap {
		result = append(result, speciality)
	}
	return result, nil
}
