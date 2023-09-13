package data_fix

import (
	"github.com/youngpto/zgz-db/enum"
	"github.com/youngpto/zgz-db/model"
	"github.com/youngpto/zgz-db/service"
)

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

	service.InsertHeroUpgradeReward(batch)
}
