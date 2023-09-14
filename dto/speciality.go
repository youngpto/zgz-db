package dto

import "github.com/youngpto/zgz-db/enum"

type UserHeroSpeciality struct {
	Level                         int32                 // 专长层级
	TakeAlongSpecialityResourceId enum.HeroSpeciality   // 装备的专长资源ID
	ChoosePool                    []enum.HeroSpeciality // 本层可选的专长资源ID
}
