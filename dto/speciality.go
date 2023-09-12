package dto

import "github.com/youngpto/zgz-db/enum"

type UserHeroSpeciality struct {
	Level         int32               // 专长层级
	SpecialityRId enum.HeroSpeciality // 专长的资源ID
}
