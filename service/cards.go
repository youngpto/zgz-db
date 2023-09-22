package service

import (
	"errors"
	"github.com/go-xorm/xorm"
	"github.com/youngpto/zgz-db/base"
	"github.com/youngpto/zgz-db/dto"
	"github.com/youngpto/zgz-db/model"
	"xorm.io/builder"
)

func InsertOrUpdateUserHeroCards(session *xorm.Session, cards *model.UserHeroCards) error {
	exist, err := session.Table("user_hero_cards").
		Where(builder.Eq{
			"user_id": cards.UserId,
			"card_id": cards.CardId,
		}).Exist()
	if err != nil {
		return err
	}
	if exist {
		_, err := session.Where(builder.Eq{
			"user_id": cards.UserId,
			"card_id": cards.CardId,
		}).MustCols("unlock_quantity").
			Update(&model.UserHeroCards{
				UnlockQuantity: cards.UnlockQuantity,
			})
		if err != nil {
			return err
		}
	} else {
		_, err := session.Insert(&model.UserHeroCards{
			UserId:         cards.UserId,
			CardId:         cards.CardId,
			UnlockQuantity: cards.UnlockQuantity,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

// AllExistInUserCards 传入的卡牌玩家是否都拥有
func AllExistInUserCards(userId int64, cards []dto.CardInfo) (bool, error) {
	session := base.Engine.NewSession()
	defer session.Close()

	if err := session.Begin(); err != nil {
		return false, err
	}

	for _, card := range cards {
		mCard := new(model.UserHeroCards)
		_, err := session.Where(builder.Eq{
			"user_id": userId,
			"card_id": card.CardId,
		}).Get(mCard)
		if err != nil {
			return false, err
		}
		if mCard.UnlockQuantity < int(card.Number) {
			return false, errors.New("存在非法卡牌")
		}
	}
	err := session.Commit()
	if err != nil {
		return false, err
	}
	return true, nil
}
