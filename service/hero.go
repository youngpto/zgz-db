package service

import (
	"fmt"
	"github.com/youngpto/zgz-db/base"
	"github.com/youngpto/zgz-db/model"
	"xorm.io/builder"
)

func GetHeroId(resourceId int64) (int64, bool) {
	result := new(model.HeroProperty)
	ok, err := base.Engine.Table("hero").
		Where(builder.Eq{"resource_id": resourceId}).
		Cols("id").
		Get(result)
	if err != nil {
		fmt.Println(err.Error())
		return 0, false
	}
	return result.Id, ok
}
