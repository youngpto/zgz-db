package data_fix

import (
	"fmt"
	"github.com/youngpto/zgz-db/base"
	"github.com/youngpto/zgz-db/enum"
	"github.com/youngpto/zgz-db/model"
	"github.com/youngpto/zgz-db/service"
	"sync/atomic"
	"time"
)

type PageResult struct {
	Page int
	Size int
}

func ResetFirstSpeciality() {
	// customize
	var userHeroSpecialityRuleInfoList []model.SpecialityRuleInfo
	err := base.Engine.Table("speciality_rule").Alias("sr").
		Join("LEFT", []string{"speciality", "s"}, "sr.speciality_id = s.id").
		Where("sr.level = ?", 1).
		Find(&userHeroSpecialityRuleInfoList)
	if err != nil {
		println(err.Error())
		fmt.Println("RA1")
		return
	}
	// customize end

	taskChannel := make(chan PageResult, 100)
	var overTaskCount uint32
	// 获取用户总量
	count, _ := base.Engine.Count(new(model.User))
	pageSize := 10 // 每次读取10个用户id
	pageCount := int(count) / pageSize
	if int(count)%pageSize > 1 {
		pageCount++
	}
	go func() {
		for i := 0; i < pageCount; i++ {
			taskChannel <- PageResult{
				Page: i,
				Size: pageSize,
			}
		}
		close(taskChannel)
	}()

	var workGoroutine int32
	for result := range taskChannel {
		for atomic.LoadInt32(&workGoroutine) > 500 {
			fmt.Println("任务队列满，等待执行，当前任务", result.Page+1)
			time.Sleep(5 * time.Second)
		}
		atomic.AddInt32(&workGoroutine, 1)
		go func(pageResult PageResult) {
			defer atomic.AddInt32(&workGoroutine, -1)
			fmt.Println("执行任务", pageResult.Page+1, "总任务数", pageCount)
			// logic start
			var users []model.User
			_ = base.Engine.Cols("id", "tmp_achieve_num").
				Limit(pageSize, pageResult.Page*pageSize).
				Find(&users)
			for _, user := range users {
				for _, info := range userHeroSpecialityRuleInfoList {
					err = service.UpdateTakeAlongFormUserSpecialityByLevel(&model.UserSpeciality{
						UserId:       user.Id,
						HeroId:       int64(info.SpecialityRule.HeroId),
						Level:        1,
						SpecialityId: info.Speciality.Id,
					})
					if err != nil {
						println(err.Error())
						fmt.Println("RA2")
						continue
					}
				}
			}
			// logic end
			newCount := atomic.AddUint32(&overTaskCount, 1)
			fmt.Println("任务", pageResult.Page+1, "已完成", "当前进度", newCount, "/", pageCount)
		}(result)
	}

	for atomic.LoadInt32(&workGoroutine) > 0 {
		fmt.Println("等待任务结束")
		time.Sleep(5 * time.Second)
	}
}

func InsertMercenaryHero() {
	err := service.InsertHero([]*model.Hero{
		{
			ResourceId: int(enum.Mercenary),
			Name:       "佣兵",
			CodeName:   "mercenary",
			Life:       9,
			Reason:     5,
			Power:      4,
			Agile:      4,
			Knowledge:  2,
			Will:       2,
			Config:     "{\n    \"hero_id\":1,\n    \"m_cardDeckId\": 1,\n    \"m_cardDeckName\": \"Default\",\n    \"m_serverId\": 0,\n    \"m_cardDeckList\": {\n        \"data\": [\n            {\n                \"m_id\": 61,\n                \"m_cardCount\": 1\n            },\n            {\n                \"m_id\": 62,\n                \"m_cardCount\": 1\n            },\n            {\n                \"m_id\": 123,\n                \"m_cardCount\": 1\n            },\n            {\n                \"m_id\": 211,\n                \"m_cardCount\": 1\n            },\n            {\n                \"m_id\": 212,\n                \"m_cardCount\": 1\n            },\n            {\n                \"m_id\": 1061,\n                \"m_cardCount\": 2\n            },\n            {\n                \"m_id\": 1062,\n                \"m_cardCount\": 2\n            },\n            {\n                \"m_id\": 1107,\n                \"m_cardCount\": 1\n            },\n            {\n                \"m_id\": 1125,\n                \"m_cardCount\": 2\n            },\n            {\n                \"m_id\": 1124,\n                \"m_cardCount\": 1\n            },\n            {\n                \"m_id\": 1202,\n                \"m_cardCount\": 2\n            },\n            {\n                \"m_id\": 1211,\n                \"m_cardCount\": 2\n            },\n            {\n                \"m_id\": 1212,\n                \"m_cardCount\": 2\n            },\n            {\n                \"m_id\": 2061,\n                \"m_cardCount\": 2\n            },\n            {\n                \"m_id\": 2062,\n                \"m_cardCount\": 2\n            },\n            {\n                \"m_id\": 2101,\n                \"m_cardCount\": 2\n            },\n            {\n                \"m_id\": 2124,\n                \"m_cardCount\": 1\n            },\n            {\n                \"m_id\": 2211,\n                \"m_cardCount\": 2\n            },\n            {\n                \"m_id\": 2212,\n                \"m_cardCount\": 2\n            }\n        ]\n    }\n}\n",
		},
	})
	if err != nil {
		return
	}
}

func UpdateMercenaryInfo() {
	heroId, _ := service.GetHeroId(int64(enum.Mercenary))

	// 更新专长表
	var batchSpeciality []*model.Speciality
	batchSpeciality = append(batchSpeciality, &model.Speciality{
		ResourceId:  int(enum.SpecialityCraftyHand),
		HeroId:      int(heroId),
		Name:        "巧手",
		Desc:        "巧手",
		EnglishName: "巧手",
		EnglishDesc: "巧手",
	})
	batchSpeciality = append(batchSpeciality, &model.Speciality{
		ResourceId:  int(enum.SpecialityDeterrence),
		HeroId:      int(heroId),
		Name:        "威慑",
		Desc:        "威慑",
		EnglishName: "威慑",
		EnglishDesc: "威慑",
	})
	batchSpeciality = append(batchSpeciality, &model.Speciality{
		ResourceId:  int(enum.SpecialityNegotiate),
		HeroId:      int(heroId),
		Name:        "交涉",
		Desc:        "交涉",
		EnglishName: "交涉",
		EnglishDesc: "交涉",
	})
	batchSpeciality = append(batchSpeciality, &model.Speciality{
		ResourceId:  int(enum.SpecialityBeastTraining),
		HeroId:      int(heroId),
		Name:        "驯兽",
		Desc:        "驯兽",
		EnglishName: "驯兽",
		EnglishDesc: "驯兽",
	})
	batchSpeciality = append(batchSpeciality, &model.Speciality{
		ResourceId:  int(enum.SpecialityMedicine),
		HeroId:      int(heroId),
		Name:        "医药",
		Desc:        "医药",
		EnglishName: "医药",
		EnglishDesc: "医药",
	})
	service.InsertSpeciality(batchSpeciality)

	// 更新专长规则表
	var batchSpecialityRule []*model.SpecialityRule
	var sId int64
	sId, _ = service.GetHeroSpecialityId(int(heroId), int(enum.SpecialityCraftyHand))
	batchSpecialityRule = append(batchSpecialityRule, &model.SpecialityRule{
		HeroId:       int(heroId),
		Level:        1,
		SpecialityId: sId,
	})
	sId, _ = service.GetHeroSpecialityId(int(heroId), int(enum.SpecialityDeterrence))
	batchSpecialityRule = append(batchSpecialityRule, &model.SpecialityRule{
		HeroId:       int(heroId),
		Level:        2,
		SpecialityId: sId,
	})
	sId, _ = service.GetHeroSpecialityId(int(heroId), int(enum.SpecialityNegotiate))
	batchSpecialityRule = append(batchSpecialityRule, &model.SpecialityRule{
		HeroId:       int(heroId),
		Level:        2,
		SpecialityId: sId,
	})
	sId, _ = service.GetHeroSpecialityId(int(heroId), int(enum.SpecialityBeastTraining))
	batchSpecialityRule = append(batchSpecialityRule, &model.SpecialityRule{
		HeroId:       int(heroId),
		Level:        3,
		SpecialityId: sId,
	})
	sId, _ = service.GetHeroSpecialityId(int(heroId), int(enum.SpecialityMedicine))
	batchSpecialityRule = append(batchSpecialityRule, &model.SpecialityRule{
		HeroId:       int(heroId),
		Level:        3,
		SpecialityId: sId,
	})
	service.InsertSpecialityRule(batchSpecialityRule)

	// 更新被动表
	var text = []string{"手牌上限+3。", "开局受到1点伤害、1点恐惧，额外抽3牌。",
		"开局时，将1张<逢凶化吉>加入你的手牌。", "开局受到1点伤害、1点恐惧，装备槽+1。",
		"你的首个回合开始时，从卡组检索1牌。", "每局首次被击败时，保留1点生命/理智。"}
	var batchPassive []*model.Passive
	batchPassive = append(batchPassive, &model.Passive{
		ResourceId:  int(enum.HeroPassive1),
		HeroId:      int(heroId),
		Name:        text[0],
		Desc:        text[0],
		EnglishName: text[0],
		EnglishDesc: text[0],
	})
	batchPassive = append(batchPassive, &model.Passive{
		ResourceId:  int(enum.HeroPassive2),
		HeroId:      int(heroId),
		Name:        text[1],
		Desc:        text[1],
		EnglishName: text[1],
		EnglishDesc: text[1],
	})
	batchPassive = append(batchPassive, &model.Passive{
		ResourceId:  int(enum.HeroPassive3),
		HeroId:      int(heroId),
		Name:        text[2],
		Desc:        text[2],
		EnglishName: text[2],
		EnglishDesc: text[2],
	})
	batchPassive = append(batchPassive, &model.Passive{
		ResourceId:  int(enum.HeroPassive4),
		HeroId:      int(heroId),
		Name:        text[3],
		Desc:        text[3],
		EnglishName: text[3],
		EnglishDesc: text[3],
	})
	batchPassive = append(batchPassive, &model.Passive{
		ResourceId:  int(enum.HeroPassive5),
		HeroId:      int(heroId),
		Name:        text[4],
		Desc:        text[4],
		EnglishName: text[4],
		EnglishDesc: text[4],
	})
	batchPassive = append(batchPassive, &model.Passive{
		ResourceId:  int(enum.HeroPassive6),
		HeroId:      int(heroId),
		Name:        text[5],
		Desc:        text[5],
		EnglishName: text[5],
		EnglishDesc: text[5],
	})
	service.InsertPassive(batchPassive)

	// 更新被动规则表
	var batchPassiveRule []*model.PassiveRule
	var pId int64
	pId, _ = service.GetHeroPassiveId(int(heroId), int(enum.HeroPassive1))
	batchPassiveRule = append(batchPassiveRule, &model.PassiveRule{
		HeroId:    int(heroId),
		Level:     1,
		PassiveId: pId,
	})
	pId, _ = service.GetHeroPassiveId(int(heroId), int(enum.HeroPassive2))
	batchPassiveRule = append(batchPassiveRule, &model.PassiveRule{
		HeroId:    int(heroId),
		Level:     1,
		PassiveId: pId,
	})
	pId, _ = service.GetHeroPassiveId(int(heroId), int(enum.HeroPassive3))
	batchPassiveRule = append(batchPassiveRule, &model.PassiveRule{
		HeroId:    int(heroId),
		Level:     1,
		PassiveId: pId,
	})
	pId, _ = service.GetHeroPassiveId(int(heroId), int(enum.HeroPassive4))
	batchPassiveRule = append(batchPassiveRule, &model.PassiveRule{
		HeroId:    int(heroId),
		Level:     2,
		PassiveId: pId,
	})
	pId, _ = service.GetHeroPassiveId(int(heroId), int(enum.HeroPassive5))
	batchPassiveRule = append(batchPassiveRule, &model.PassiveRule{
		HeroId:    int(heroId),
		Level:     2,
		PassiveId: pId,
	})
	pId, _ = service.GetHeroPassiveId(int(heroId), int(enum.HeroPassive6))
	batchPassiveRule = append(batchPassiveRule, &model.PassiveRule{
		HeroId:    int(heroId),
		Level:     2,
		PassiveId: pId,
	})
	service.InsertPassiveRule(batchPassiveRule)

	// 更新升级奖励表
	var batchHeroUpgradeReward []*model.HeroUpgradeReward
	batchHeroUpgradeReward = append(batchHeroUpgradeReward, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        1,
		Category:    1,
		ObjectId:    int(enum.SpecialityDeterrence),
		ObjectValue: 2,
	})
	batchHeroUpgradeReward = append(batchHeroUpgradeReward, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        1,
		Category:    1,
		ObjectId:    int(enum.SpecialityNegotiate),
		ObjectValue: 2,
	})
	batchHeroUpgradeReward = append(batchHeroUpgradeReward, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        2,
		Category:    4,
		ObjectValue: 1,
	})
	batchHeroUpgradeReward = append(batchHeroUpgradeReward, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        3,
		Category:    2,
		ObjectId:    63,
		ObjectValue: 1,
	})
	batchHeroUpgradeReward = append(batchHeroUpgradeReward, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        4,
		Category:    3,
		ObjectId:    int(enum.HeroPassive1),
		ObjectValue: 1,
	})
	batchHeroUpgradeReward = append(batchHeroUpgradeReward, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        4,
		Category:    3,
		ObjectId:    int(enum.HeroPassive2),
		ObjectValue: 1,
	})
	batchHeroUpgradeReward = append(batchHeroUpgradeReward, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        4,
		Category:    3,
		ObjectId:    int(enum.HeroPassive3),
		ObjectValue: 1,
	})
	batchHeroUpgradeReward = append(batchHeroUpgradeReward, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        5,
		Category:    4,
		ObjectValue: 1,
	})
	batchHeroUpgradeReward = append(batchHeroUpgradeReward, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        6,
		Category:    2,
		ObjectId:    1063,
		ObjectValue: 2,
	})
	batchHeroUpgradeReward = append(batchHeroUpgradeReward, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        7,
		Category:    1,
		ObjectId:    int(enum.SpecialityBeastTraining),
		ObjectValue: 3,
	})
	batchHeroUpgradeReward = append(batchHeroUpgradeReward, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        7,
		Category:    1,
		ObjectId:    int(enum.SpecialityMedicine),
		ObjectValue: 3,
	})
	batchHeroUpgradeReward = append(batchHeroUpgradeReward, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        8,
		Category:    3,
		ObjectId:    int(enum.HeroPassive4),
		ObjectValue: 2,
	})
	batchHeroUpgradeReward = append(batchHeroUpgradeReward, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        8,
		Category:    3,
		ObjectId:    int(enum.HeroPassive5),
		ObjectValue: 2,
	})
	batchHeroUpgradeReward = append(batchHeroUpgradeReward, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        8,
		Category:    3,
		ObjectId:    int(enum.HeroPassive6),
		ObjectValue: 2,
	})
	batchHeroUpgradeReward = append(batchHeroUpgradeReward, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        9,
		Category:    2,
		ObjectId:    1064,
		ObjectValue: 2,
	})
	service.InsertHeroUpgradeReward(batchHeroUpgradeReward)
}
