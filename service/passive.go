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

// GetHeroPassiveId 获取被动主键Id
func GetHeroPassiveId(heroId, resourcePId int) (int64, error) {
	result := new(model.Passive)
	_, err := base.Engine.Where(builder.Eq{"hero_id": heroId, "resource_id": resourcePId}).
		Get(result)
	if err != nil {
		return 0, err
	}
	return result.Id, nil
}

// GetHeroPassiveResourceId 根据被动ID解析资源ID
func GetHeroPassiveResourceId(pkID int64) (enum.HeroPassive, error) {
	result := new(model.Passive)
	_, err := base.Engine.ID(pkID).
		Get(result)
	if err != nil {
		return 0, err
	}
	return enum.HeroPassive(result.ResourceId), nil
}

// DropAllPassiveRule 清空被动规则表
func DropAllPassiveRule() {
	_, _ = base.Engine.Where("1=1").Delete(new(model.PassiveRule))
}

// InsertPassiveRule 新增被动规则
func InsertPassiveRule(batch []*model.PassiveRule) {
	_, _ = base.Engine.Insert(batch)
}

// InsertNotExistUserPassive 不重复插入玩家被动
func InsertNotExistUserPassive(session *xorm.Session, passive *model.UserPassive) (bool, error) {
	exist, err := session.Where(builder.Eq{
		"user_id":    passive.UserId,
		"level":      passive.Level,
		"hero_id":    passive.HeroId,
		"passive_id": passive.PassiveId},
	).Exist(new(model.UserPassive))
	if err != nil {
		return false, err
	}
	if !exist {
		insert, err := session.Insert(passive)
		if err != nil {
			return false, err
		}
		return insert > 0, nil
	}
	return false, nil
}

// CancelTakeAlongUserPassiveByLevel 取消玩家在对应层级设置的被动
func CancelTakeAlongUserPassiveByLevel(passive *model.UserPassive) error {
	_, err := base.Engine.
		Where(builder.Eq{
			"user_id":    passive.UserId,
			"hero_id":    passive.HeroId,
			"level":      passive.Level,
			"take_along": true,
		}).MustCols("take_along").
		Update(&model.UserPassive{
			TakeAlong: false,
		})
	if err != nil {
		return err
	}
	return nil
}

// UpdateTakeAlongFormUserPassiveByLevel 调整玩家在指定层级选中的被动
func UpdateTakeAlongFormUserPassiveByLevel(passive *model.UserPassive) error {
	session := base.Engine.NewSession()
	defer session.Close()

	if err := session.Begin(); err != nil {
		return err
	}
	// 校验该被动是否符合配置规则
	matchRule, err := session.Where(builder.Eq{
		"hero_id":    passive.HeroId,
		"level":      passive.Level,
		"passive_id": passive.PassiveId,
	}).Exist(new(model.PassiveRule))
	if err != nil {
		return err
	}
	if !matchRule {
		return errors.New("不符合配置规则")
	}
	// 检查是否已经配置了被动
	exist, err := session.Table("user_passive").
		Where(builder.Eq{
			"user_id":    passive.UserId,
			"hero_id":    passive.HeroId,
			"level":      passive.Level,
			"take_along": true,
		}).Exist()
	if err != nil {
		return err
	}
	// 如果存在的话则取消
	if exist {
		_, err := session.
			Where(builder.Eq{
				"user_id":    passive.UserId,
				"hero_id":    passive.HeroId,
				"level":      passive.Level,
				"take_along": true,
			}).MustCols("take_along").
			Update(&model.UserPassive{
				TakeAlong: false,
			})
		if err != nil {
			return err
		}
	}
	_, err = session.
		Where(builder.Eq{
			"user_id":    passive.UserId,
			"hero_id":    passive.HeroId,
			"level":      passive.Level,
			"passive_id": passive.PassiveId,
		}).MustCols("take_along").
		Update(&model.UserPassive{
			TakeAlong: true,
		})
	if err != nil {
		return err
	}
	return session.Commit()
}

// FindAllUserPassiveByUser 获取玩家所有被动
func FindAllUserPassiveByUser(userId int64, heroId int64) ([]*dto.UserHeroPassive, error) {
	tempMap := make(map[int]*dto.UserHeroPassive)

	var userHeroPassiveInfoList []model.UserHeroPassiveInfo
	err := base.Engine.Table("user_passive").Alias("up").
		Join("LEFT", []string{"passive", "p"}, "up.passive_id = p.id").
		Where("up.user_id = ? AND up.hero_id = ? AND up.take_along = 1", userId, heroId).
		Find(&userHeroPassiveInfoList)
	if err != nil {
		return nil, err
	}
	for _, passive := range userHeroPassiveInfoList {
		tempMap[passive.Level] = &dto.UserHeroPassive{
			Level:                      int32(passive.Level),
			TakeAlongPassiveResourceId: enum.HeroPassive(passive.ResourceId),
			ChoosePool:                 make([]enum.HeroPassive, 0),
		}
	}

	var userHeroPassiveRuleInfoList []model.PassiveRuleInfo
	err = base.Engine.Table("passive_rule").Alias("pr").
		Join("LEFT", []string{"passive", "p"}, "pr.passive_id = p.id").
		Where("pr.hero_id = ?", heroId).
		Find(&userHeroPassiveRuleInfoList)
	if err != nil {
		return nil, err
	}
	for _, passiveRuleInfo := range userHeroPassiveRuleInfoList {
		if old, ok := tempMap[passiveRuleInfo.Level]; ok {
			old.ChoosePool = append(old.ChoosePool, enum.HeroPassive(passiveRuleInfo.ResourceId))
		} else {
			newDto := &dto.UserHeroPassive{
				Level:                      int32(passiveRuleInfo.Level),
				TakeAlongPassiveResourceId: -1,
				ChoosePool:                 make([]enum.HeroPassive, 0),
			}
			newDto.ChoosePool = append(newDto.ChoosePool, enum.HeroPassive(passiveRuleInfo.ResourceId))
			tempMap[passiveRuleInfo.Level] = newDto
		}
	}

	var result []*dto.UserHeroPassive
	for _, passiveInfo := range tempMap {
		result = append(result, passiveInfo)
	}
	return result, nil
}
