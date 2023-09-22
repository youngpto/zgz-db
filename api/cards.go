package api

import (
	"github.com/youngpto/zgz-db/base"
	"github.com/youngpto/zgz-db/dto"
	"github.com/youngpto/zgz-db/service"
)

func GetUserHeroCardsInfo(userId int64, resourceHeroId int64) (map[int64]dto.CardInfo, error) {
	heroId, err := service.GetHeroId(resourceHeroId)
	if err != nil {
		return nil, err
	}
	session := base.Engine.NewSession()
	defer session.Close()

	if err := session.Begin(); err != nil {
		return nil, err
	}
	cardUnlockInfos, err := service.GetUserHeroCardsInfo(session, userId, heroId)
	if err != nil {
		return nil, err
	}
	var result = make(map[int64]dto.CardInfo)
	for _, info := range cardUnlockInfos {
		result[info.CardId] = dto.CardInfo{
			CardId: info.CardId,
			Number: int32(info.UnlockQuantity),
		}
	}
	err = session.Commit()
	if err != nil {
		return nil, err
	}
	return result, nil
}
