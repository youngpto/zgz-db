package data_fix

import (
	"github.com/youngpto/zgz-db/enum"
	"github.com/youngpto/zgz-db/model"
	"github.com/youngpto/zgz-db/service"
)

func RebuildHeroCardsTable() {
	service.DropAllHeroCards()
	var batch []*model.HeroCards
	var heroId int64

	heroId, _ = service.GetHeroId(int64(enum.Taoist))
	batch = append(batch, &model.HeroCards{
		HeroId: heroId,
		CardId: 34,
	})
	batch = append(batch, &model.HeroCards{
		HeroId: heroId,
		CardId: 1039,
	})
	batch = append(batch, &model.HeroCards{
		HeroId: heroId,
		CardId: 1040,
	})

	heroId, _ = service.GetHeroId(int64(enum.Gambler))
	batch = append(batch, &model.HeroCards{
		HeroId: heroId,
		CardId: 1013,
	})
	batch = append(batch, &model.HeroCards{
		HeroId: heroId,
		CardId: 1015,
	})
	batch = append(batch, &model.HeroCards{
		HeroId: heroId,
		CardId: 1014,
	})

	heroId, _ = service.GetHeroId(int64(enum.Doctor))
	batch = append(batch, &model.HeroCards{
		HeroId: heroId,
		CardId: 1046,
	})
	batch = append(batch, &model.HeroCards{
		HeroId: heroId,
		CardId: 1047,
	})
	batch = append(batch, &model.HeroCards{
		HeroId: heroId,
		CardId: 1048,
	})

	heroId, _ = service.GetHeroId(int64(enum.College))
	batch = append(batch, &model.HeroCards{
		HeroId: heroId,
		CardId: 1024,
	})
	batch = append(batch, &model.HeroCards{
		HeroId: heroId,
		CardId: 2023,
	})
	batch = append(batch, &model.HeroCards{
		HeroId: heroId,
		CardId: 2024,
	})

	heroId, _ = service.GetHeroId(int64(enum.Police))
	batch = append(batch, &model.HeroCards{
		HeroId: heroId,
		CardId: 1004,
	})
	batch = append(batch, &model.HeroCards{
		HeroId: heroId,
		CardId: 2,
	})
	batch = append(batch, &model.HeroCards{
		HeroId: heroId,
		CardId: 2003,
	})

	heroId, _ = service.GetHeroId(int64(enum.Pupil))
	batch = append(batch, &model.HeroCards{
		HeroId: heroId,
		CardId: 1053,
	})
	batch = append(batch, &model.HeroCards{
		HeroId: heroId,
		CardId: 1057,
	})
	batch = append(batch, &model.HeroCards{
		HeroId: heroId,
		CardId: 2055,
	})
	service.InsertHeroCards(batch)
}
