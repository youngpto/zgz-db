package data_fix

import (
	"github.com/youngpto/zgz-db/enum"
	"github.com/youngpto/zgz-db/model"
	"github.com/youngpto/zgz-db/service"
)

// RebuildPassiveRuleTable 重建被动规则表
func RebuildPassiveRuleTable() {
	service.DropAllPassiveRule()

	var batch []*model.PassiveRule
	var heroId, pId int64

	heroId, _ = service.GetHeroId(0)
	pId, _ = service.GetHeroPassiveId(int(heroId), int(enum.HeroPassive1))
	batch = append(batch, &model.PassiveRule{
		HeroId:    int(heroId),
		Level:     1,
		PassiveId: pId,
	})
	pId, _ = service.GetHeroPassiveId(int(heroId), int(enum.HeroPassive2))
	batch = append(batch, &model.PassiveRule{
		HeroId:    int(heroId),
		Level:     1,
		PassiveId: pId,
	})
	pId, _ = service.GetHeroPassiveId(int(heroId), int(enum.HeroPassive3))
	batch = append(batch, &model.PassiveRule{
		HeroId:    int(heroId),
		Level:     1,
		PassiveId: pId,
	})
	pId, _ = service.GetHeroPassiveId(int(heroId), int(enum.HeroPassive4))
	batch = append(batch, &model.PassiveRule{
		HeroId:    int(heroId),
		Level:     2,
		PassiveId: pId,
	})
	pId, _ = service.GetHeroPassiveId(int(heroId), int(enum.HeroPassive5))
	batch = append(batch, &model.PassiveRule{
		HeroId:    int(heroId),
		Level:     2,
		PassiveId: pId,
	})
	pId, _ = service.GetHeroPassiveId(int(heroId), int(enum.HeroPassive6))
	batch = append(batch, &model.PassiveRule{
		HeroId:    int(heroId),
		Level:     2,
		PassiveId: pId,
	})

	heroId, _ = service.GetHeroId(1)
	pId, _ = service.GetHeroPassiveId(int(heroId), int(enum.HeroPassive1))
	batch = append(batch, &model.PassiveRule{
		HeroId:    int(heroId),
		Level:     1,
		PassiveId: pId,
	})
	pId, _ = service.GetHeroPassiveId(int(heroId), int(enum.HeroPassive2))
	batch = append(batch, &model.PassiveRule{
		HeroId:    int(heroId),
		Level:     1,
		PassiveId: pId,
	})
	pId, _ = service.GetHeroPassiveId(int(heroId), int(enum.HeroPassive3))
	batch = append(batch, &model.PassiveRule{
		HeroId:    int(heroId),
		Level:     1,
		PassiveId: pId,
	})
	pId, _ = service.GetHeroPassiveId(int(heroId), int(enum.HeroPassive4))
	batch = append(batch, &model.PassiveRule{
		HeroId:    int(heroId),
		Level:     2,
		PassiveId: pId,
	})
	pId, _ = service.GetHeroPassiveId(int(heroId), int(enum.HeroPassive5))
	batch = append(batch, &model.PassiveRule{
		HeroId:    int(heroId),
		Level:     2,
		PassiveId: pId,
	})
	pId, _ = service.GetHeroPassiveId(int(heroId), int(enum.HeroPassive6))
	batch = append(batch, &model.PassiveRule{
		HeroId:    int(heroId),
		Level:     2,
		PassiveId: pId,
	})

	heroId, _ = service.GetHeroId(2)
	pId, _ = service.GetHeroPassiveId(int(heroId), int(enum.HeroPassive1))
	batch = append(batch, &model.PassiveRule{
		HeroId:    int(heroId),
		Level:     1,
		PassiveId: pId,
	})
	pId, _ = service.GetHeroPassiveId(int(heroId), int(enum.HeroPassive2))
	batch = append(batch, &model.PassiveRule{
		HeroId:    int(heroId),
		Level:     1,
		PassiveId: pId,
	})
	pId, _ = service.GetHeroPassiveId(int(heroId), int(enum.HeroPassive3))
	batch = append(batch, &model.PassiveRule{
		HeroId:    int(heroId),
		Level:     1,
		PassiveId: pId,
	})
	pId, _ = service.GetHeroPassiveId(int(heroId), int(enum.HeroPassive4))
	batch = append(batch, &model.PassiveRule{
		HeroId:    int(heroId),
		Level:     2,
		PassiveId: pId,
	})
	pId, _ = service.GetHeroPassiveId(int(heroId), int(enum.HeroPassive5))
	batch = append(batch, &model.PassiveRule{
		HeroId:    int(heroId),
		Level:     2,
		PassiveId: pId,
	})
	pId, _ = service.GetHeroPassiveId(int(heroId), int(enum.HeroPassive6))
	batch = append(batch, &model.PassiveRule{
		HeroId:    int(heroId),
		Level:     2,
		PassiveId: pId,
	})

	heroId, _ = service.GetHeroId(3)
	pId, _ = service.GetHeroPassiveId(int(heroId), int(enum.HeroPassive1))
	batch = append(batch, &model.PassiveRule{
		HeroId:    int(heroId),
		Level:     1,
		PassiveId: pId,
	})
	pId, _ = service.GetHeroPassiveId(int(heroId), int(enum.HeroPassive2))
	batch = append(batch, &model.PassiveRule{
		HeroId:    int(heroId),
		Level:     1,
		PassiveId: pId,
	})
	pId, _ = service.GetHeroPassiveId(int(heroId), int(enum.HeroPassive3))
	batch = append(batch, &model.PassiveRule{
		HeroId:    int(heroId),
		Level:     1,
		PassiveId: pId,
	})
	pId, _ = service.GetHeroPassiveId(int(heroId), int(enum.HeroPassive4))
	batch = append(batch, &model.PassiveRule{
		HeroId:    int(heroId),
		Level:     2,
		PassiveId: pId,
	})
	pId, _ = service.GetHeroPassiveId(int(heroId), int(enum.HeroPassive5))
	batch = append(batch, &model.PassiveRule{
		HeroId:    int(heroId),
		Level:     2,
		PassiveId: pId,
	})
	pId, _ = service.GetHeroPassiveId(int(heroId), int(enum.HeroPassive6))
	batch = append(batch, &model.PassiveRule{
		HeroId:    int(heroId),
		Level:     2,
		PassiveId: pId,
	})

	heroId, _ = service.GetHeroId(4)
	pId, _ = service.GetHeroPassiveId(int(heroId), int(enum.HeroPassive1))
	batch = append(batch, &model.PassiveRule{
		HeroId:    int(heroId),
		Level:     1,
		PassiveId: pId,
	})
	pId, _ = service.GetHeroPassiveId(int(heroId), int(enum.HeroPassive2))
	batch = append(batch, &model.PassiveRule{
		HeroId:    int(heroId),
		Level:     1,
		PassiveId: pId,
	})
	pId, _ = service.GetHeroPassiveId(int(heroId), int(enum.HeroPassive3))
	batch = append(batch, &model.PassiveRule{
		HeroId:    int(heroId),
		Level:     1,
		PassiveId: pId,
	})
	pId, _ = service.GetHeroPassiveId(int(heroId), int(enum.HeroPassive4))
	batch = append(batch, &model.PassiveRule{
		HeroId:    int(heroId),
		Level:     2,
		PassiveId: pId,
	})
	pId, _ = service.GetHeroPassiveId(int(heroId), int(enum.HeroPassive5))
	batch = append(batch, &model.PassiveRule{
		HeroId:    int(heroId),
		Level:     2,
		PassiveId: pId,
	})
	pId, _ = service.GetHeroPassiveId(int(heroId), int(enum.HeroPassive6))
	batch = append(batch, &model.PassiveRule{
		HeroId:    int(heroId),
		Level:     2,
		PassiveId: pId,
	})

	heroId, _ = service.GetHeroId(5)
	pId, _ = service.GetHeroPassiveId(int(heroId), int(enum.HeroPassive1))
	batch = append(batch, &model.PassiveRule{
		HeroId:    int(heroId),
		Level:     1,
		PassiveId: pId,
	})
	pId, _ = service.GetHeroPassiveId(int(heroId), int(enum.HeroPassive2))
	batch = append(batch, &model.PassiveRule{
		HeroId:    int(heroId),
		Level:     1,
		PassiveId: pId,
	})
	pId, _ = service.GetHeroPassiveId(int(heroId), int(enum.HeroPassive3))
	batch = append(batch, &model.PassiveRule{
		HeroId:    int(heroId),
		Level:     1,
		PassiveId: pId,
	})
	pId, _ = service.GetHeroPassiveId(int(heroId), int(enum.HeroPassive4))
	batch = append(batch, &model.PassiveRule{
		HeroId:    int(heroId),
		Level:     2,
		PassiveId: pId,
	})
	pId, _ = service.GetHeroPassiveId(int(heroId), int(enum.HeroPassive5))
	batch = append(batch, &model.PassiveRule{
		HeroId:    int(heroId),
		Level:     2,
		PassiveId: pId,
	})
	pId, _ = service.GetHeroPassiveId(int(heroId), int(enum.HeroPassive6))
	batch = append(batch, &model.PassiveRule{
		HeroId:    int(heroId),
		Level:     2,
		PassiveId: pId,
	})
	service.InsertPassiveRule(batch)
}
