package dto

import "github.com/youngpto/zgz-db/enum"

type UserHeroPassive struct {
	Level                      int32              // 被动层级
	TakeAlongPassiveResourceId enum.HeroPassive   // 装备的被动资源ID
	ChoosePool                 []enum.HeroPassive // 本层可选的被动资源ID
}
