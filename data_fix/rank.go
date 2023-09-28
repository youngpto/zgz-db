package data_fix

import (
	"fmt"
	"github.com/youngpto/zgz-db/api"
	"github.com/youngpto/zgz-db/base"
	"github.com/youngpto/zgz-db/enum"
	"github.com/youngpto/zgz-db/model"
	"github.com/youngpto/zgz-db/service"
	"runtime"
	"sync"
	"sync/atomic"
	"xorm.io/builder"
)

func InitReleaseRank() {
	// 获取用户总量
	var workGoroutine int32 = 0
	count, _ := base.Engine.
		Where(builder.Gte{"tmp_achieve_num": 1}).
		Count(new(model.User))
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
			_ = base.Engine.
				Where(builder.Gte{"tmp_achieve_num": 1}).
				Limit(pageSize, pc*pageSize).
				Iterate(new(model.User), func(idx int, bean interface{}) error {
					user := bean.(*model.User)
					for i := 0; i < 6; i++ {
						_, ch, _ := api.HeroGainExperience(user.Id, int64(i), float64(200*user.TmpAchieveNum))
						for i2 := range ch {
							if i2 == 0 {
								fmt.Println("ERROR Exp")
							}
						}
					}
					return nil
				})
			fmt.Println(pc+1, "/", pageCount, "已完成")
		}(i)
	}
	wg.Done()
}

func RebuildRankTable() {
	service.DropAllRank()
	var batch []*model.Rank

	batch = append(batch, &model.Rank{
		Name:       "0",
		Rank:       0,
		Experience: 200,
	})
	batch = append(batch, &model.Rank{
		Name:       "1",
		Rank:       1,
		Experience: 300,
	})
	batch = append(batch, &model.Rank{
		Name:       "2",
		Rank:       2,
		Experience: 400,
	})
	batch = append(batch, &model.Rank{
		Name:       "3",
		Rank:       3,
		Experience: 500,
	})
	batch = append(batch, &model.Rank{
		Name:       "4",
		Rank:       4,
		Experience: 600,
	})
	batch = append(batch, &model.Rank{
		Name:       "5",
		Rank:       5,
		Experience: 700,
	})
	batch = append(batch, &model.Rank{
		Name:       "6",
		Rank:       6,
		Experience: 800,
	})
	batch = append(batch, &model.Rank{
		Name:       "7",
		Rank:       7,
		Experience: 900,
	})
	batch = append(batch, &model.Rank{
		Name:       "8",
		Rank:       8,
		Experience: 1000,
	})
	batch = append(batch, &model.Rank{
		Name:       "9",
		Rank:       9,
		Experience: 0,
	})

	service.InsertRank(batch)
}

func RebuildHeroUpgradeRewardTable() {
	service.DropAllHeroUpgradeReward()
	var heroId int64
	var batch []*model.HeroUpgradeReward

	heroId, _ = service.GetHeroId(0)
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        1,
		Category:    1,
		ObjectId:    int(enum.SpecialityBeastTraining),
		ObjectValue: 2,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        1,
		Category:    1,
		ObjectId:    int(enum.SpecialityNature),
		ObjectValue: 2,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        2,
		Category:    4,
		ObjectValue: 1,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        3,
		Category:    2,
		ObjectId:    34,
		ObjectValue: 1,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        4,
		Category:    3,
		ObjectId:    int(enum.HeroPassive1),
		ObjectValue: 1,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        4,
		Category:    3,
		ObjectId:    int(enum.HeroPassive2),
		ObjectValue: 1,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        4,
		Category:    3,
		ObjectId:    int(enum.HeroPassive3),
		ObjectValue: 1,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        5,
		Category:    4,
		ObjectValue: 1,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        6,
		Category:    2,
		ObjectId:    1039,
		ObjectValue: 2,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        7,
		Category:    1,
		ObjectId:    int(enum.SpecialityNegotiate),
		ObjectValue: 3,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        7,
		Category:    1,
		ObjectId:    int(enum.SpecialityMedicine),
		ObjectValue: 3,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        8,
		Category:    3,
		ObjectId:    int(enum.HeroPassive4),
		ObjectValue: 2,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        8,
		Category:    3,
		ObjectId:    int(enum.HeroPassive5),
		ObjectValue: 2,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        8,
		Category:    3,
		ObjectId:    int(enum.HeroPassive6),
		ObjectValue: 2,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        9,
		Category:    2,
		ObjectId:    1040,
		ObjectValue: 2,
	})

	heroId, _ = service.GetHeroId(1)
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        1,
		Category:    1,
		ObjectId:    int(enum.SpecialityCraftyHand),
		ObjectValue: 2,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        1,
		Category:    1,
		ObjectId:    int(enum.SpecialityVerbalTrick),
		ObjectValue: 2,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        2,
		Category:    4,
		ObjectValue: 1,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        3,
		Category:    2,
		ObjectId:    1013,
		ObjectValue: 2,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        4,
		Category:    3,
		ObjectId:    int(enum.HeroPassive1),
		ObjectValue: 1,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        4,
		Category:    3,
		ObjectId:    int(enum.HeroPassive2),
		ObjectValue: 1,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        4,
		Category:    3,
		ObjectId:    int(enum.HeroPassive3),
		ObjectValue: 1,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        5,
		Category:    4,
		ObjectValue: 1,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        6,
		Category:    2,
		ObjectId:    1015,
		ObjectValue: 2,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        7,
		Category:    1,
		ObjectId:    int(enum.SpecialityPerformance),
		ObjectValue: 3,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        7,
		Category:    1,
		ObjectId:    int(enum.SpecialityNegotiate),
		ObjectValue: 3,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        8,
		Category:    3,
		ObjectId:    int(enum.HeroPassive4),
		ObjectValue: 2,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        8,
		Category:    3,
		ObjectId:    int(enum.HeroPassive5),
		ObjectValue: 2,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        8,
		Category:    3,
		ObjectId:    int(enum.HeroPassive6),
		ObjectValue: 2,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        9,
		Category:    2,
		ObjectId:    1014,
		ObjectValue: 1,
	})

	heroId, _ = service.GetHeroId(2)
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        1,
		Category:    1,
		ObjectId:    int(enum.SpecialityDeterrence),
		ObjectValue: 2,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        1,
		Category:    1,
		ObjectId:    int(enum.SpecialityHumanity),
		ObjectValue: 2,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        2,
		Category:    4,
		ObjectValue: 1,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        3,
		Category:    2,
		ObjectId:    1046,
		ObjectValue: 2,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        4,
		Category:    3,
		ObjectId:    int(enum.HeroPassive1),
		ObjectValue: 1,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        4,
		Category:    3,
		ObjectId:    int(enum.HeroPassive2),
		ObjectValue: 1,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        4,
		Category:    3,
		ObjectId:    int(enum.HeroPassive3),
		ObjectValue: 1,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        5,
		Category:    4,
		ObjectValue: 1,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        6,
		Category:    2,
		ObjectId:    1047,
		ObjectValue: 2,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        7,
		Category:    1,
		ObjectId:    int(enum.SpecialityCraftyHand),
		ObjectValue: 3,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        7,
		Category:    1,
		ObjectId:    int(enum.SpecialityMysterious),
		ObjectValue: 3,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        8,
		Category:    3,
		ObjectId:    int(enum.HeroPassive4),
		ObjectValue: 2,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        8,
		Category:    3,
		ObjectId:    int(enum.HeroPassive5),
		ObjectValue: 2,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        8,
		Category:    3,
		ObjectId:    int(enum.HeroPassive6),
		ObjectValue: 2,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        9,
		Category:    2,
		ObjectId:    1048,
		ObjectValue: 1,
	})

	heroId, _ = service.GetHeroId(3)
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        1,
		Category:    1,
		ObjectId:    int(enum.SpecialityComputer),
		ObjectValue: 2,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        1,
		Category:    1,
		ObjectId:    int(enum.SpecialityNegotiate),
		ObjectValue: 2,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        2,
		Category:    4,
		ObjectValue: 1,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        3,
		Category:    2,
		ObjectId:    1024,
		ObjectValue: 2,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        4,
		Category:    3,
		ObjectId:    int(enum.HeroPassive1),
		ObjectValue: 1,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        4,
		Category:    3,
		ObjectId:    int(enum.HeroPassive2),
		ObjectValue: 1,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        4,
		Category:    3,
		ObjectId:    int(enum.HeroPassive3),
		ObjectValue: 1,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        5,
		Category:    4,
		ObjectValue: 1,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        6,
		Category:    2,
		ObjectId:    2023,
		ObjectValue: 2,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        7,
		Category:    1,
		ObjectId:    int(enum.SpecialityHumanity),
		ObjectValue: 3,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        7,
		Category:    1,
		ObjectId:    int(enum.SpecialityVerbalTrick),
		ObjectValue: 3,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        8,
		Category:    3,
		ObjectId:    int(enum.HeroPassive4),
		ObjectValue: 2,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        8,
		Category:    3,
		ObjectId:    int(enum.HeroPassive5),
		ObjectValue: 2,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        8,
		Category:    3,
		ObjectId:    int(enum.HeroPassive6),
		ObjectValue: 2,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        9,
		Category:    2,
		ObjectId:    2024,
		ObjectValue: 1,
	})

	heroId, _ = service.GetHeroId(4)
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        1,
		Category:    1,
		ObjectId:    int(enum.SpecialityMedicine),
		ObjectValue: 2,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        1,
		Category:    1,
		ObjectId:    int(enum.SpecialityBeastTraining),
		ObjectValue: 2,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        2,
		Category:    4,
		ObjectValue: 1,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        3,
		Category:    2,
		ObjectId:    1004,
		ObjectValue: 2,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        4,
		Category:    3,
		ObjectId:    int(enum.HeroPassive1),
		ObjectValue: 1,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        4,
		Category:    3,
		ObjectId:    int(enum.HeroPassive2),
		ObjectValue: 1,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        4,
		Category:    3,
		ObjectId:    int(enum.HeroPassive3),
		ObjectValue: 1,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        5,
		Category:    4,
		ObjectValue: 1,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        6,
		Category:    2,
		ObjectId:    2,
		ObjectValue: 1,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        7,
		Category:    1,
		ObjectId:    int(enum.SpecialityCharm),
		ObjectValue: 3,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        7,
		Category:    1,
		ObjectId:    int(enum.SpecialityComputer),
		ObjectValue: 3,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        8,
		Category:    3,
		ObjectId:    int(enum.HeroPassive4),
		ObjectValue: 2,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        8,
		Category:    3,
		ObjectId:    int(enum.HeroPassive5),
		ObjectValue: 2,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        8,
		Category:    3,
		ObjectId:    int(enum.HeroPassive6),
		ObjectValue: 2,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        9,
		Category:    2,
		ObjectId:    2003,
		ObjectValue: 2,
	})

	heroId, _ = service.GetHeroId(5)
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        1,
		Category:    1,
		ObjectId:    int(enum.SpecialityCharm),
		ObjectValue: 2,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        1,
		Category:    1,
		ObjectId:    int(enum.SpecialityPerformance),
		ObjectValue: 2,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        2,
		Category:    4,
		ObjectValue: 1,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        3,
		Category:    2,
		ObjectId:    1053,
		ObjectValue: 2,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        4,
		Category:    3,
		ObjectId:    int(enum.HeroPassive1),
		ObjectValue: 1,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        4,
		Category:    3,
		ObjectId:    int(enum.HeroPassive2),
		ObjectValue: 1,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        4,
		Category:    3,
		ObjectId:    int(enum.HeroPassive3),
		ObjectValue: 1,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        5,
		Category:    4,
		ObjectValue: 1,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        6,
		Category:    2,
		ObjectId:    1057,
		ObjectValue: 2,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        7,
		Category:    1,
		ObjectId:    int(enum.SpecialityComputer),
		ObjectValue: 3,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        7,
		Category:    1,
		ObjectId:    int(enum.SpecialityNegotiate),
		ObjectValue: 3,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        8,
		Category:    3,
		ObjectId:    int(enum.HeroPassive4),
		ObjectValue: 2,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        8,
		Category:    3,
		ObjectId:    int(enum.HeroPassive5),
		ObjectValue: 2,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        8,
		Category:    3,
		ObjectId:    int(enum.HeroPassive6),
		ObjectValue: 2,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        9,
		Category:    2,
		ObjectId:    2055,
		ObjectValue: 1,
	})

	service.InsertHeroUpgradeReward(batch)
}

func AddCardsReward() {
	var heroId int64
	var batch []*model.HeroUpgradeReward

	heroId, _ = service.GetHeroId(0)
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        3,
		Category:    2,
		ObjectId:    34,
		ObjectValue: 1,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        6,
		Category:    2,
		ObjectId:    1039,
		ObjectValue: 2,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        9,
		Category:    2,
		ObjectId:    1040,
		ObjectValue: 2,
	})

	heroId, _ = service.GetHeroId(1)
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        3,
		Category:    2,
		ObjectId:    1013,
		ObjectValue: 2,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        6,
		Category:    2,
		ObjectId:    1015,
		ObjectValue: 2,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        9,
		Category:    2,
		ObjectId:    1014,
		ObjectValue: 1,
	})

	heroId, _ = service.GetHeroId(2)
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        3,
		Category:    2,
		ObjectId:    1046,
		ObjectValue: 2,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        6,
		Category:    2,
		ObjectId:    1047,
		ObjectValue: 2,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        9,
		Category:    2,
		ObjectId:    1048,
		ObjectValue: 1,
	})

	heroId, _ = service.GetHeroId(3)
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        3,
		Category:    2,
		ObjectId:    1024,
		ObjectValue: 2,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        6,
		Category:    2,
		ObjectId:    2023,
		ObjectValue: 2,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        9,
		Category:    2,
		ObjectId:    2024,
		ObjectValue: 1,
	})

	heroId, _ = service.GetHeroId(4)
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        3,
		Category:    2,
		ObjectId:    1004,
		ObjectValue: 2,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        6,
		Category:    2,
		ObjectId:    2,
		ObjectValue: 1,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        9,
		Category:    2,
		ObjectId:    2003,
		ObjectValue: 2,
	})

	heroId, _ = service.GetHeroId(5)
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        3,
		Category:    2,
		ObjectId:    1053,
		ObjectValue: 2,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        6,
		Category:    2,
		ObjectId:    1057,
		ObjectValue: 2,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        9,
		Category:    2,
		ObjectId:    2055,
		ObjectValue: 1,
	})

	service.InsertHeroUpgradeReward(batch)
}
