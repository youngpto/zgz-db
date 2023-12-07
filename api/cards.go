package api

import (
	"github.com/youngpto/zgz-db/dto"
	"github.com/youngpto/zgz-db/service"
)

func CheckCardRule(userId int64, cards []dto.CardInfo) (bool, error) {
	return service.AllExistInUserCards(userId, cards)
}

func GetUserCardsFromDLC(userId int64) ([]int32, error) {
	userCards, err := service.GetUserCards(userId)
	if err != nil {
		return nil, err
	}
	var result []int32
	for _, card := range userCards {
		switch card.CardId {
		case 63, 1063, 1064: // 佣兵升级卡牌
			continue
		case 34, 1039, 1040: // 道士升级卡牌
			continue
		case 1013, 1015, 1014: // 赌徒升级卡牌
			continue
		case 1046, 1047, 1048: // 医生升级卡牌
			continue
		case 1024, 2023, 2024: // 大学生升级卡牌
			continue
		case 1004, 2, 2003: // 警察升级卡牌
			continue
		case 1053, 1057, 2055: // 小学生升级卡牌
			continue
		default:
			result = append(result, int32(card.CardId))
		}
	}
	return result, nil
}
