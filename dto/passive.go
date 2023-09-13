package dto

import "github.com/youngpto/zgz-db/enum"

type UserHeroPassive struct {
	Level             int32            // 被动层级
	PassiveResourceId enum.HeroPassive // 被动的资源ID
}
