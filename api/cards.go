package api

import (
	"github.com/youngpto/zgz-db/dto"
	"github.com/youngpto/zgz-db/service"
)

func CheckCardRule(userId int64, cards []dto.CardInfo) (bool, error) {
	return service.AllExistInUserCards(userId, cards)
}
