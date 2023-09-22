package service

import (
	"github.com/go-xorm/xorm"
	"github.com/youngpto/zgz-db/base"
	"github.com/youngpto/zgz-db/model"
	"xorm.io/builder"
)

func DropAllHeroCards() {
	_, _ = base.Engine.
		Where("1=1").
		Delete(new(model.HeroCards))
}

func InsertHeroCards(cards []*model.HeroCards) {
	_, _ = base.Engine.Insert(&cards)
}

func InsertOrUpdateUserHeroCards(session *xorm.Session, cards *model.UserHeroCards) error {
	exist, err := session.Table("user_hero_cards").
		Where(builder.Eq{
			"user_id":      cards.UserId,
			"hero_id":      cards.HeroId,
			"hero_card_id": cards.HeroCardId,
		}).Exist()
	if err != nil {
		return err
	}
	if exist {
		_, err := session.Where(builder.Eq{
			"user_id":      cards.UserId,
			"hero_id":      cards.HeroId,
			"hero_card_id": cards.HeroCardId,
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
			HeroId:         cards.HeroId,
			HeroCardId:     cards.HeroCardId,
			UnlockQuantity: cards.UnlockQuantity,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

// GetUserHeroCardsInfo 获取玩家英雄卡牌信息
func GetUserHeroCardsInfo(session *xorm.Session, userId int64, heroId int64) ([]model.UserHeroCardUnlockInfo, error) {
	var result []model.UserHeroCardUnlockInfo
	err := session.Table("hero_cards").Alias("hc").
		Join("LEFT", []string{"user_hero_cards", "uhc"}, "hc.id = uhc.hero_card_id").
		Where("hc.hero_id = ? AND uhc.user_id = ?", heroId, userId).
		Find(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
