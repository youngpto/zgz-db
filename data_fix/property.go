package data_fix

import (
	"zgz-db/model"
	"zgz-db/service"
)

func FixHeroBaseProperty() {
	service.UpdateHeroBaseProperty(0, &model.HeroProperty{
		Life:      6,
		Reason:    8,
		Power:     3,
		Agile:     2,
		Knowledge: 2,
		Will:      5,
	})
	service.UpdateHeroBaseProperty(1, &model.HeroProperty{
		Life:      7,
		Reason:    7,
		Power:     3,
		Agile:     4,
		Knowledge: 3,
		Will:      2,
	})
	service.UpdateHeroBaseProperty(2, &model.HeroProperty{
		Life:      6,
		Reason:    8,
		Power:     2,
		Agile:     2,
		Knowledge: 4,
		Will:      4,
	})
	service.UpdateHeroBaseProperty(3, &model.HeroProperty{
		Life:      5,
		Reason:    9,
		Power:     2,
		Agile:     2,
		Knowledge: 3,
		Will:      5,
	})
	service.UpdateHeroBaseProperty(4, &model.HeroProperty{
		Life:      8,
		Reason:    6,
		Power:     5,
		Agile:     2,
		Knowledge: 3,
		Will:      2,
	})
	service.UpdateHeroBaseProperty(5, &model.HeroProperty{
		Life:      7,
		Reason:    7,
		Power:     3,
		Agile:     3,
		Knowledge: 3,
		Will:      3,
	})
}
