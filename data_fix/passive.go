package data_fix

import (
	"github.com/youngpto/zgz-db/enum"
	"github.com/youngpto/zgz-db/model"
	"github.com/youngpto/zgz-db/service"
)

// RebuildPassiveTable 重建被动表
func RebuildPassiveTable() {
	var text = []string{"手牌上限+3。", "开局受到1点伤害、1点恐惧，额外抽3牌。",
		"开局时，将1张<逢凶化吉>加入你的手牌。", "开局受到1点伤害、1点恐惧，装备槽+1。",
		"你的首个回合开始时，从卡组检索1牌。", "每局首次被击败时，保留1点生命/理智。"}
	var batch []*model.Passive
	for i := 0; i < 6; i++ {
		heroId, _ := service.GetHeroId(int64(i))
		batch = append(batch, &model.Passive{
			ResourceId:  int(enum.HeroPassive1),
			HeroId:      int(heroId),
			Name:        text[0],
			Desc:        text[0],
			EnglishName: text[0],
			EnglishDesc: text[0],
		})
		batch = append(batch, &model.Passive{
			ResourceId:  int(enum.HeroPassive2),
			HeroId:      int(heroId),
			Name:        text[1],
			Desc:        text[1],
			EnglishName: text[1],
			EnglishDesc: text[1],
		})
		batch = append(batch, &model.Passive{
			ResourceId:  int(enum.HeroPassive3),
			HeroId:      int(heroId),
			Name:        text[2],
			Desc:        text[2],
			EnglishName: text[2],
			EnglishDesc: text[2],
		})
		batch = append(batch, &model.Passive{
			ResourceId:  int(enum.HeroPassive4),
			HeroId:      int(heroId),
			Name:        text[3],
			Desc:        text[3],
			EnglishName: text[3],
			EnglishDesc: text[3],
		})
		batch = append(batch, &model.Passive{
			ResourceId:  int(enum.HeroPassive5),
			HeroId:      int(heroId),
			Name:        text[4],
			Desc:        text[4],
			EnglishName: text[4],
			EnglishDesc: text[4],
		})
		batch = append(batch, &model.Passive{
			ResourceId:  int(enum.HeroPassive6),
			HeroId:      int(heroId),
			Name:        text[5],
			Desc:        text[5],
			EnglishName: text[5],
			EnglishDesc: text[5],
		})
	}
	service.InsertPassive(batch)
}

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
