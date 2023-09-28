package data_fix

import (
	"fmt"
	"github.com/youngpto/zgz-db/base"
	"github.com/youngpto/zgz-db/enum"
	"github.com/youngpto/zgz-db/model"
	"github.com/youngpto/zgz-db/service"
	"runtime"
	"sync"
	"sync/atomic"
)

// RebuildUserSpecial 重建用户专长表
func RebuildUserSpecial() {
	service.DropAllUserSpeciality()

	// 所有角色的默认配置
	var heroId, sId int64
	sidList := make(map[int64]int64)
	heroId, _ = service.GetHeroId(0)
	sId, _ = service.GetHeroSpecialityId(int(heroId), int(enum.SpecialityMysterious))
	sidList[heroId] = sId

	heroId, _ = service.GetHeroId(1)
	sId, _ = service.GetHeroSpecialityId(int(heroId), int(enum.SpecialityCharm))
	sidList[heroId] = sId

	heroId, _ = service.GetHeroId(2)
	sId, _ = service.GetHeroSpecialityId(int(heroId), int(enum.SpecialityMedicine))
	sidList[heroId] = sId

	heroId, _ = service.GetHeroId(3)
	sId, _ = service.GetHeroSpecialityId(int(heroId), int(enum.SpecialityNature))
	sidList[heroId] = sId

	heroId, _ = service.GetHeroId(4)
	sId, _ = service.GetHeroSpecialityId(int(heroId), int(enum.SpecialityDeterrence))
	sidList[heroId] = sId

	heroId, _ = service.GetHeroId(5)
	sId, _ = service.GetHeroSpecialityId(int(heroId), int(enum.SpecialityHumanity))
	sidList[heroId] = sId

	// 获取用户总量
	var workGoroutine int32 = 0
	count, _ := base.Engine.Count(new(model.User))
	pageSize := 50 // 每次读取50个用户id
	pageCount := int(count) / pageSize
	if int(count)%pageSize > 1 {
		pageCount++
	}

	var wg sync.WaitGroup
	for i := 0; i < pageCount; i++ {
		for atomic.LoadInt32(&workGoroutine) >= 50 {
			runtime.Gosched()
		}
		wg.Add(1)
		atomic.AddInt32(&workGoroutine, 1)
		go func(pc int) {
			defer wg.Done()
			defer atomic.AddInt32(&workGoroutine, -1)
			_ = base.Engine.Cols("id").
				Limit(pageSize, pc*pageSize).
				Iterate(new(model.User), func(idx int, bean interface{}) error {
					user := bean.(*model.User)
					var batch []*model.UserSpeciality
					for heroId, sid := range sidList {
						batch = append(batch, &model.UserSpeciality{
							UserId:       user.Id,
							HeroId:       heroId,
							Level:        1,
							SpecialityId: sid,
						})
					}
					service.InsertUserSpeciality(batch)
					return nil
				})
			fmt.Println(pc+1, "/", pageCount, "已完成")
		}(i)
	}
	wg.Done()
}

// RebuildSpecialRuleTable 重建专长规则表
func RebuildSpecialRuleTable() {
	service.DropAllSpecialityRule()
	var batch []*model.SpecialityRule
	var heroId, sId int64

	heroId, _ = service.GetHeroId(0)
	sId, _ = service.GetHeroSpecialityId(int(heroId), int(enum.SpecialityMysterious))
	batch = append(batch, &model.SpecialityRule{
		HeroId:       int(heroId),
		Level:        1,
		SpecialityId: sId,
	})
	sId, _ = service.GetHeroSpecialityId(int(heroId), int(enum.SpecialityBeastTraining))
	batch = append(batch, &model.SpecialityRule{
		HeroId:       int(heroId),
		Level:        2,
		SpecialityId: sId,
	})
	sId, _ = service.GetHeroSpecialityId(int(heroId), int(enum.SpecialityNature))
	batch = append(batch, &model.SpecialityRule{
		HeroId:       int(heroId),
		Level:        2,
		SpecialityId: sId,
	})
	sId, _ = service.GetHeroSpecialityId(int(heroId), int(enum.SpecialityNegotiate))
	batch = append(batch, &model.SpecialityRule{
		HeroId:       int(heroId),
		Level:        3,
		SpecialityId: sId,
	})
	sId, _ = service.GetHeroSpecialityId(int(heroId), int(enum.SpecialityMedicine))
	batch = append(batch, &model.SpecialityRule{
		HeroId:       int(heroId),
		Level:        3,
		SpecialityId: sId,
	})

	heroId, _ = service.GetHeroId(1)
	sId, _ = service.GetHeroSpecialityId(int(heroId), int(enum.SpecialityCharm))
	batch = append(batch, &model.SpecialityRule{
		HeroId:       int(heroId),
		Level:        1,
		SpecialityId: sId,
	})
	sId, _ = service.GetHeroSpecialityId(int(heroId), int(enum.SpecialityCraftyHand))
	batch = append(batch, &model.SpecialityRule{
		HeroId:       int(heroId),
		Level:        2,
		SpecialityId: sId,
	})
	sId, _ = service.GetHeroSpecialityId(int(heroId), int(enum.SpecialityVerbalTrick))
	batch = append(batch, &model.SpecialityRule{
		HeroId:       int(heroId),
		Level:        2,
		SpecialityId: sId,
	})
	sId, _ = service.GetHeroSpecialityId(int(heroId), int(enum.SpecialityPerformance))
	batch = append(batch, &model.SpecialityRule{
		HeroId:       int(heroId),
		Level:        3,
		SpecialityId: sId,
	})
	sId, _ = service.GetHeroSpecialityId(int(heroId), int(enum.SpecialityNegotiate))
	batch = append(batch, &model.SpecialityRule{
		HeroId:       int(heroId),
		Level:        3,
		SpecialityId: sId,
	})

	heroId, _ = service.GetHeroId(2)
	sId, _ = service.GetHeroSpecialityId(int(heroId), int(enum.SpecialityMedicine))
	batch = append(batch, &model.SpecialityRule{
		HeroId:       int(heroId),
		Level:        1,
		SpecialityId: sId,
	})
	sId, _ = service.GetHeroSpecialityId(int(heroId), int(enum.SpecialityDeterrence))
	batch = append(batch, &model.SpecialityRule{
		HeroId:       int(heroId),
		Level:        2,
		SpecialityId: sId,
	})
	sId, _ = service.GetHeroSpecialityId(int(heroId), int(enum.SpecialityHumanity))
	batch = append(batch, &model.SpecialityRule{
		HeroId:       int(heroId),
		Level:        2,
		SpecialityId: sId,
	})
	sId, _ = service.GetHeroSpecialityId(int(heroId), int(enum.SpecialityCraftyHand))
	batch = append(batch, &model.SpecialityRule{
		HeroId:       int(heroId),
		Level:        3,
		SpecialityId: sId,
	})
	sId, _ = service.GetHeroSpecialityId(int(heroId), int(enum.SpecialityMysterious))
	batch = append(batch, &model.SpecialityRule{
		HeroId:       int(heroId),
		Level:        3,
		SpecialityId: sId,
	})

	heroId, _ = service.GetHeroId(3)
	sId, _ = service.GetHeroSpecialityId(int(heroId), int(enum.SpecialityNature))
	batch = append(batch, &model.SpecialityRule{
		HeroId:       int(heroId),
		Level:        1,
		SpecialityId: sId,
	})
	sId, _ = service.GetHeroSpecialityId(int(heroId), int(enum.SpecialityComputer))
	batch = append(batch, &model.SpecialityRule{
		HeroId:       int(heroId),
		Level:        2,
		SpecialityId: sId,
	})
	sId, _ = service.GetHeroSpecialityId(int(heroId), int(enum.SpecialityNegotiate))
	batch = append(batch, &model.SpecialityRule{
		HeroId:       int(heroId),
		Level:        2,
		SpecialityId: sId,
	})
	sId, _ = service.GetHeroSpecialityId(int(heroId), int(enum.SpecialityHumanity))
	batch = append(batch, &model.SpecialityRule{
		HeroId:       int(heroId),
		Level:        3,
		SpecialityId: sId,
	})
	sId, _ = service.GetHeroSpecialityId(int(heroId), int(enum.SpecialityVerbalTrick))
	batch = append(batch, &model.SpecialityRule{
		HeroId:       int(heroId),
		Level:        3,
		SpecialityId: sId,
	})

	heroId, _ = service.GetHeroId(4)
	sId, _ = service.GetHeroSpecialityId(int(heroId), int(enum.SpecialityDeterrence))
	batch = append(batch, &model.SpecialityRule{
		HeroId:       int(heroId),
		Level:        1,
		SpecialityId: sId,
	})
	sId, _ = service.GetHeroSpecialityId(int(heroId), int(enum.SpecialityMedicine))
	batch = append(batch, &model.SpecialityRule{
		HeroId:       int(heroId),
		Level:        2,
		SpecialityId: sId,
	})
	sId, _ = service.GetHeroSpecialityId(int(heroId), int(enum.SpecialityBeastTraining))
	batch = append(batch, &model.SpecialityRule{
		HeroId:       int(heroId),
		Level:        2,
		SpecialityId: sId,
	})
	sId, _ = service.GetHeroSpecialityId(int(heroId), int(enum.SpecialityCharm))
	batch = append(batch, &model.SpecialityRule{
		HeroId:       int(heroId),
		Level:        3,
		SpecialityId: sId,
	})
	sId, _ = service.GetHeroSpecialityId(int(heroId), int(enum.SpecialityComputer))
	batch = append(batch, &model.SpecialityRule{
		HeroId:       int(heroId),
		Level:        3,
		SpecialityId: sId,
	})

	heroId, _ = service.GetHeroId(5)
	sId, _ = service.GetHeroSpecialityId(int(heroId), int(enum.SpecialityHumanity))
	batch = append(batch, &model.SpecialityRule{
		HeroId:       int(heroId),
		Level:        1,
		SpecialityId: sId,
	})
	sId, _ = service.GetHeroSpecialityId(int(heroId), int(enum.SpecialityCharm))
	batch = append(batch, &model.SpecialityRule{
		HeroId:       int(heroId),
		Level:        2,
		SpecialityId: sId,
	})
	sId, _ = service.GetHeroSpecialityId(int(heroId), int(enum.SpecialityPerformance))
	batch = append(batch, &model.SpecialityRule{
		HeroId:       int(heroId),
		Level:        2,
		SpecialityId: sId,
	})
	sId, _ = service.GetHeroSpecialityId(int(heroId), int(enum.SpecialityComputer))
	batch = append(batch, &model.SpecialityRule{
		HeroId:       int(heroId),
		Level:        3,
		SpecialityId: sId,
	})
	sId, _ = service.GetHeroSpecialityId(int(heroId), int(enum.SpecialityNegotiate))
	batch = append(batch, &model.SpecialityRule{
		HeroId:       int(heroId),
		Level:        3,
		SpecialityId: sId,
	})
	service.InsertSpecialityRule(batch)
}

// RebuildSpecialTable 重建专长表
func RebuildSpecialTable() {
	service.DropAllSpeciality()
	var batch []*model.Speciality

	heroId, _ := service.GetHeroId(0)
	batch = append(batch, &model.Speciality{
		ResourceId:  int(enum.SpecialityMysterious),
		HeroId:      int(heroId),
		Name:        "神秘",
		Desc:        "神秘",
		EnglishName: "神秘",
		EnglishDesc: "神秘",
	})
	batch = append(batch, &model.Speciality{
		ResourceId:  int(enum.SpecialityBeastTraining),
		HeroId:      int(heroId),
		Name:        "驯兽",
		Desc:        "驯兽",
		EnglishName: "驯兽",
		EnglishDesc: "驯兽",
	})
	batch = append(batch, &model.Speciality{
		ResourceId:  int(enum.SpecialityNature),
		HeroId:      int(heroId),
		Name:        "自然",
		Desc:        "自然",
		EnglishName: "自然",
		EnglishDesc: "自然",
	})
	batch = append(batch, &model.Speciality{
		ResourceId:  int(enum.SpecialityNegotiate),
		HeroId:      int(heroId),
		Name:        "交涉",
		Desc:        "交涉",
		EnglishName: "交涉",
		EnglishDesc: "交涉",
	})
	batch = append(batch, &model.Speciality{
		ResourceId:  int(enum.SpecialityMedicine),
		HeroId:      int(heroId),
		Name:        "医药",
		Desc:        "医药",
		EnglishName: "医药",
		EnglishDesc: "医药",
	})

	heroId, _ = service.GetHeroId(1)
	batch = append(batch, &model.Speciality{
		ResourceId:  int(enum.SpecialityCharm),
		HeroId:      int(heroId),
		Name:        "魅力",
		Desc:        "魅力",
		EnglishName: "魅力",
		EnglishDesc: "魅力",
	})
	batch = append(batch, &model.Speciality{
		ResourceId:  int(enum.SpecialityCraftyHand),
		HeroId:      int(heroId),
		Name:        "巧手",
		Desc:        "巧手",
		EnglishName: "巧手",
		EnglishDesc: "巧手",
	})
	batch = append(batch, &model.Speciality{
		ResourceId:  int(enum.SpecialityVerbalTrick),
		HeroId:      int(heroId),
		Name:        "话术",
		Desc:        "话术",
		EnglishName: "话术",
		EnglishDesc: "话术",
	})
	batch = append(batch, &model.Speciality{
		ResourceId:  int(enum.SpecialityPerformance),
		HeroId:      int(heroId),
		Name:        "表演",
		Desc:        "表演",
		EnglishName: "表演",
		EnglishDesc: "表演",
	})
	batch = append(batch, &model.Speciality{
		ResourceId:  int(enum.SpecialityNegotiate),
		HeroId:      int(heroId),
		Name:        "交涉",
		Desc:        "交涉",
		EnglishName: "交涉",
		EnglishDesc: "交涉",
	})

	heroId, _ = service.GetHeroId(2)
	batch = append(batch, &model.Speciality{
		ResourceId:  int(enum.SpecialityMedicine),
		HeroId:      int(heroId),
		Name:        "医药",
		Desc:        "医药",
		EnglishName: "医药",
		EnglishDesc: "医药",
	})
	batch = append(batch, &model.Speciality{
		ResourceId:  int(enum.SpecialityDeterrence),
		HeroId:      int(heroId),
		Name:        "威慑",
		Desc:        "威慑",
		EnglishName: "威慑",
		EnglishDesc: "威慑",
	})
	batch = append(batch, &model.Speciality{
		ResourceId:  int(enum.SpecialityHumanity),
		HeroId:      int(heroId),
		Name:        "人文",
		Desc:        "人文",
		EnglishName: "人文",
		EnglishDesc: "人文",
	})
	batch = append(batch, &model.Speciality{
		ResourceId:  int(enum.SpecialityCraftyHand),
		HeroId:      int(heroId),
		Name:        "巧手",
		Desc:        "巧手",
		EnglishName: "巧手",
		EnglishDesc: "巧手",
	})
	batch = append(batch, &model.Speciality{
		ResourceId:  int(enum.SpecialityMysterious),
		HeroId:      int(heroId),
		Name:        "神秘",
		Desc:        "神秘",
		EnglishName: "神秘",
		EnglishDesc: "神秘",
	})

	heroId, _ = service.GetHeroId(3)
	batch = append(batch, &model.Speciality{
		ResourceId:  int(enum.SpecialityNature),
		HeroId:      int(heroId),
		Name:        "自然",
		Desc:        "自然",
		EnglishName: "自然",
		EnglishDesc: "自然",
	})
	batch = append(batch, &model.Speciality{
		ResourceId:  int(enum.SpecialityComputer),
		HeroId:      int(heroId),
		Name:        "电脑",
		Desc:        "电脑",
		EnglishName: "电脑",
		EnglishDesc: "电脑",
	})
	batch = append(batch, &model.Speciality{
		ResourceId:  int(enum.SpecialityNegotiate),
		HeroId:      int(heroId),
		Name:        "交涉",
		Desc:        "交涉",
		EnglishName: "交涉",
		EnglishDesc: "交涉",
	})
	batch = append(batch, &model.Speciality{
		ResourceId:  int(enum.SpecialityHumanity),
		HeroId:      int(heroId),
		Name:        "人文",
		Desc:        "人文",
		EnglishName: "人文",
		EnglishDesc: "人文",
	})
	batch = append(batch, &model.Speciality{
		ResourceId:  int(enum.SpecialityVerbalTrick),
		HeroId:      int(heroId),
		Name:        "话术",
		Desc:        "话术",
		EnglishName: "话术",
		EnglishDesc: "话术",
	})

	heroId, _ = service.GetHeroId(4)
	batch = append(batch, &model.Speciality{
		ResourceId:  int(enum.SpecialityDeterrence),
		HeroId:      int(heroId),
		Name:        "威慑",
		Desc:        "威慑",
		EnglishName: "威慑",
		EnglishDesc: "威慑",
	})
	batch = append(batch, &model.Speciality{
		ResourceId:  int(enum.SpecialityMedicine),
		HeroId:      int(heroId),
		Name:        "医药",
		Desc:        "医药",
		EnglishName: "医药",
		EnglishDesc: "医药",
	})
	batch = append(batch, &model.Speciality{
		ResourceId:  int(enum.SpecialityBeastTraining),
		HeroId:      int(heroId),
		Name:        "驯兽",
		Desc:        "驯兽",
		EnglishName: "驯兽",
		EnglishDesc: "驯兽",
	})
	batch = append(batch, &model.Speciality{
		ResourceId:  int(enum.SpecialityCharm),
		HeroId:      int(heroId),
		Name:        "魅力",
		Desc:        "魅力",
		EnglishName: "魅力",
		EnglishDesc: "魅力",
	})
	batch = append(batch, &model.Speciality{
		ResourceId:  int(enum.SpecialityComputer),
		HeroId:      int(heroId),
		Name:        "电脑",
		Desc:        "电脑",
		EnglishName: "电脑",
		EnglishDesc: "电脑",
	})

	heroId, _ = service.GetHeroId(5)
	batch = append(batch, &model.Speciality{
		ResourceId:  int(enum.SpecialityHumanity),
		HeroId:      int(heroId),
		Name:        "人文",
		Desc:        "人文",
		EnglishName: "人文",
		EnglishDesc: "人文",
	})
	batch = append(batch, &model.Speciality{
		ResourceId:  int(enum.SpecialityCharm),
		HeroId:      int(heroId),
		Name:        "魅力",
		Desc:        "魅力",
		EnglishName: "魅力",
		EnglishDesc: "魅力",
	})
	batch = append(batch, &model.Speciality{
		ResourceId:  int(enum.SpecialityPerformance),
		HeroId:      int(heroId),
		Name:        "表演",
		Desc:        "表演",
		EnglishName: "表演",
		EnglishDesc: "表演",
	})
	batch = append(batch, &model.Speciality{
		ResourceId:  int(enum.SpecialityComputer),
		HeroId:      int(heroId),
		Name:        "电脑",
		Desc:        "电脑",
		EnglishName: "电脑",
		EnglishDesc: "电脑",
	})
	batch = append(batch, &model.Speciality{
		ResourceId:  int(enum.SpecialityNegotiate),
		HeroId:      int(heroId),
		Name:        "交涉",
		Desc:        "交涉",
		EnglishName: "交涉",
		EnglishDesc: "交涉",
	})
	service.InsertSpeciality(batch)
}
