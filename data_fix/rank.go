package data_fix

import (
	"fmt"
	"github.com/youngpto/zgz-db/base"
	"github.com/youngpto/zgz-db/enum"
	"github.com/youngpto/zgz-db/model"
	"github.com/youngpto/zgz-db/service"
	"xorm.io/builder"
)

func Reward0(userId int64, exp float64) {
	session := base.Engine.NewSession()
	_ = session.Begin()

	_, _ = session.Table("user_hero").
		Where(builder.Eq{"user_id": userId}).
		MustCols("experience_pool").
		Update(&model.UserHeroRank{
			ExperiencePool: exp,
		})
	var batchUserSpeciality []model.UserSpeciality
	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       1,
		Level:        1,
		SpecialityId: 21,
		TakeAlong:    true,
	})
	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       2,
		Level:        1,
		SpecialityId: 1,
		TakeAlong:    true,
	})
	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       3,
		Level:        1,
		SpecialityId: 6,
		TakeAlong:    true,
	})
	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       4,
		Level:        1,
		SpecialityId: 11,
		TakeAlong:    true,
	})
	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       5,
		Level:        1,
		SpecialityId: 16,
		TakeAlong:    true,
	})
	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       6,
		Level:        1,
		SpecialityId: 26,
		TakeAlong:    true,
	})
	_, _ = session.Insert(&batchUserSpeciality)

	_ = session.Commit()
}

func Reward1(userId int64, exp float64) {
	session := base.Engine.NewSession()
	_ = session.Begin()

	_, _ = session.Table("user_hero").
		Where(builder.Eq{"user_id": userId}).
		MustCols("rank", "experience_pool").
		Update(&model.UserHeroRank{
			Rank:           1,
			ExperiencePool: exp,
		})
	var batchUserSpeciality []model.UserSpeciality
	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       1,
		Level:        1,
		SpecialityId: 21,
		TakeAlong:    true,
	})
	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       2,
		Level:        1,
		SpecialityId: 1,
		TakeAlong:    true,
	})
	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       3,
		Level:        1,
		SpecialityId: 6,
		TakeAlong:    true,
	})
	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       4,
		Level:        1,
		SpecialityId: 11,
		TakeAlong:    true,
	})
	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       5,
		Level:        1,
		SpecialityId: 16,
		TakeAlong:    true,
	})
	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       6,
		Level:        1,
		SpecialityId: 26,
		TakeAlong:    true,
	})

	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       1,
		Level:        2,
		SpecialityId: 22,
		TakeAlong:    true,
	})
	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       1,
		Level:        2,
		SpecialityId: 23,
		TakeAlong:    false,
	})

	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       2,
		Level:        2,
		SpecialityId: 3,
		TakeAlong:    true,
	})
	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       2,
		Level:        2,
		SpecialityId: 2,
		TakeAlong:    false,
	})

	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       3,
		Level:        2,
		SpecialityId: 8,
		TakeAlong:    true,
	})
	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       3,
		Level:        2,
		SpecialityId: 7,
		TakeAlong:    false,
	})

	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       4,
		Level:        2,
		SpecialityId: 13,
		TakeAlong:    true,
	})
	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       4,
		Level:        2,
		SpecialityId: 12,
		TakeAlong:    false,
	})

	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       5,
		Level:        2,
		SpecialityId: 18,
		TakeAlong:    true,
	})
	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       5,
		Level:        2,
		SpecialityId: 17,
		TakeAlong:    false,
	})

	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       6,
		Level:        2,
		SpecialityId: 28,
		TakeAlong:    true,
	})
	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       6,
		Level:        2,
		SpecialityId: 27,
		TakeAlong:    false,
	})
	_, _ = session.Insert(&batchUserSpeciality)

	var batchLog []model.UserHeroUpgradeRewardLog
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 1,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 2,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 16,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 17,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 31,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 32,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 46,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 47,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 61,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 62,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 76,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 77,
	})
	_, _ = session.Insert(&batchLog)

	_ = session.Commit()
}

func Reward2(userId int64, exp float64) {
	session := base.Engine.NewSession()
	_ = session.Begin()

	_, _ = session.Table("user_hero").
		Where(builder.Eq{"user_id": userId}).
		MustCols("rank", "experience_pool", "total_offset").
		Update(&model.UserHeroRankAndPropertyOffset{
			Rank:           2,
			ExperiencePool: exp,
			TotalOffset:    1,
		})
	var batchUserSpeciality []model.UserSpeciality
	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       1,
		Level:        1,
		SpecialityId: 21,
		TakeAlong:    true,
	})
	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       2,
		Level:        1,
		SpecialityId: 1,
		TakeAlong:    true,
	})
	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       3,
		Level:        1,
		SpecialityId: 6,
		TakeAlong:    true,
	})
	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       4,
		Level:        1,
		SpecialityId: 11,
		TakeAlong:    true,
	})
	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       5,
		Level:        1,
		SpecialityId: 16,
		TakeAlong:    true,
	})
	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       6,
		Level:        1,
		SpecialityId: 26,
		TakeAlong:    true,
	})

	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       1,
		Level:        2,
		SpecialityId: 22,
		TakeAlong:    true,
	})
	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       1,
		Level:        2,
		SpecialityId: 23,
		TakeAlong:    false,
	})

	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       2,
		Level:        2,
		SpecialityId: 3,
		TakeAlong:    true,
	})
	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       2,
		Level:        2,
		SpecialityId: 2,
		TakeAlong:    false,
	})

	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       3,
		Level:        2,
		SpecialityId: 8,
		TakeAlong:    true,
	})
	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       3,
		Level:        2,
		SpecialityId: 7,
		TakeAlong:    false,
	})

	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       4,
		Level:        2,
		SpecialityId: 13,
		TakeAlong:    true,
	})
	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       4,
		Level:        2,
		SpecialityId: 12,
		TakeAlong:    false,
	})

	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       5,
		Level:        2,
		SpecialityId: 18,
		TakeAlong:    true,
	})
	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       5,
		Level:        2,
		SpecialityId: 17,
		TakeAlong:    false,
	})

	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       6,
		Level:        2,
		SpecialityId: 28,
		TakeAlong:    true,
	})
	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       6,
		Level:        2,
		SpecialityId: 27,
		TakeAlong:    false,
	})
	_, _ = session.Insert(&batchUserSpeciality)

	var batchLog []model.UserHeroUpgradeRewardLog
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 1,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 2,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 16,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 17,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 31,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 32,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 46,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 47,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 61,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 62,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 76,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 77,
	})

	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 3,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 18,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 33,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 48,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 63,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 78,
	})
	_, _ = session.Insert(&batchLog)

	_ = session.Commit()
}

func Reward3(userId int64, exp float64) {
	session := base.Engine.NewSession()
	_ = session.Begin()

	_, _ = session.Table("user_hero").
		Where(builder.Eq{"user_id": userId}).
		MustCols("rank", "experience_pool", "total_offset").
		Update(&model.UserHeroRankAndPropertyOffset{
			Rank:           3,
			ExperiencePool: exp,
			TotalOffset:    1,
		})
	var batchUserSpeciality []model.UserSpeciality
	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       1,
		Level:        1,
		SpecialityId: 21,
		TakeAlong:    true,
	})
	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       2,
		Level:        1,
		SpecialityId: 1,
		TakeAlong:    true,
	})
	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       3,
		Level:        1,
		SpecialityId: 6,
		TakeAlong:    true,
	})
	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       4,
		Level:        1,
		SpecialityId: 11,
		TakeAlong:    true,
	})
	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       5,
		Level:        1,
		SpecialityId: 16,
		TakeAlong:    true,
	})
	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       6,
		Level:        1,
		SpecialityId: 26,
		TakeAlong:    true,
	})

	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       1,
		Level:        2,
		SpecialityId: 22,
		TakeAlong:    true,
	})
	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       1,
		Level:        2,
		SpecialityId: 23,
		TakeAlong:    false,
	})

	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       2,
		Level:        2,
		SpecialityId: 3,
		TakeAlong:    true,
	})
	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       2,
		Level:        2,
		SpecialityId: 2,
		TakeAlong:    false,
	})

	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       3,
		Level:        2,
		SpecialityId: 8,
		TakeAlong:    true,
	})
	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       3,
		Level:        2,
		SpecialityId: 7,
		TakeAlong:    false,
	})

	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       4,
		Level:        2,
		SpecialityId: 13,
		TakeAlong:    true,
	})
	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       4,
		Level:        2,
		SpecialityId: 12,
		TakeAlong:    false,
	})

	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       5,
		Level:        2,
		SpecialityId: 18,
		TakeAlong:    true,
	})
	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       5,
		Level:        2,
		SpecialityId: 17,
		TakeAlong:    false,
	})

	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       6,
		Level:        2,
		SpecialityId: 28,
		TakeAlong:    true,
	})
	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       6,
		Level:        2,
		SpecialityId: 27,
		TakeAlong:    false,
	})
	_, _ = session.Insert(&batchUserSpeciality)

	var batchUserCards []model.UserHeroCards
	batchUserCards = append(batchUserCards, model.UserHeroCards{
		UserId:         userId,
		CardId:         34,
		UnlockQuantity: 1,
	})
	batchUserCards = append(batchUserCards, model.UserHeroCards{
		UserId:         userId,
		CardId:         1013,
		UnlockQuantity: 2,
	})
	batchUserCards = append(batchUserCards, model.UserHeroCards{
		UserId:         userId,
		CardId:         1046,
		UnlockQuantity: 2,
	})
	batchUserCards = append(batchUserCards, model.UserHeroCards{
		UserId:         userId,
		CardId:         1024,
		UnlockQuantity: 2,
	})
	batchUserCards = append(batchUserCards, model.UserHeroCards{
		UserId:         userId,
		CardId:         1004,
		UnlockQuantity: 2,
	})
	batchUserCards = append(batchUserCards, model.UserHeroCards{
		UserId:         userId,
		CardId:         1053,
		UnlockQuantity: 2,
	})
	_, _ = session.Insert(&batchUserCards)

	var batchLog []model.UserHeroUpgradeRewardLog
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 1,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 2,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 16,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 17,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 31,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 32,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 46,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 47,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 61,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 62,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 76,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 77,
	})

	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 3,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 18,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 33,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 48,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 63,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 78,
	})

	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 4,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 19,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 34,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 49,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 64,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 79,
	})
	_, _ = session.Insert(&batchLog)

	_ = session.Commit()
}

func Reward4(userId int64, exp float64) {
	session := base.Engine.NewSession()
	_ = session.Begin()

	_, _ = session.Table("user_hero").
		Where(builder.Eq{"user_id": userId}).
		MustCols("rank", "experience_pool", "total_offset").
		Update(&model.UserHeroRankAndPropertyOffset{
			Rank:           4,
			ExperiencePool: exp,
			TotalOffset:    1,
		})
	var batchUserSpeciality []model.UserSpeciality
	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       1,
		Level:        1,
		SpecialityId: 21,
		TakeAlong:    true,
	})
	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       2,
		Level:        1,
		SpecialityId: 1,
		TakeAlong:    true,
	})
	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       3,
		Level:        1,
		SpecialityId: 6,
		TakeAlong:    true,
	})
	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       4,
		Level:        1,
		SpecialityId: 11,
		TakeAlong:    true,
	})
	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       5,
		Level:        1,
		SpecialityId: 16,
		TakeAlong:    true,
	})
	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       6,
		Level:        1,
		SpecialityId: 26,
		TakeAlong:    true,
	})

	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       1,
		Level:        2,
		SpecialityId: 22,
		TakeAlong:    true,
	})
	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       1,
		Level:        2,
		SpecialityId: 23,
		TakeAlong:    false,
	})

	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       2,
		Level:        2,
		SpecialityId: 3,
		TakeAlong:    true,
	})
	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       2,
		Level:        2,
		SpecialityId: 2,
		TakeAlong:    false,
	})

	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       3,
		Level:        2,
		SpecialityId: 8,
		TakeAlong:    true,
	})
	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       3,
		Level:        2,
		SpecialityId: 7,
		TakeAlong:    false,
	})

	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       4,
		Level:        2,
		SpecialityId: 13,
		TakeAlong:    true,
	})
	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       4,
		Level:        2,
		SpecialityId: 12,
		TakeAlong:    false,
	})

	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       5,
		Level:        2,
		SpecialityId: 18,
		TakeAlong:    true,
	})
	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       5,
		Level:        2,
		SpecialityId: 17,
		TakeAlong:    false,
	})

	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       6,
		Level:        2,
		SpecialityId: 28,
		TakeAlong:    true,
	})
	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       6,
		Level:        2,
		SpecialityId: 27,
		TakeAlong:    false,
	})
	_, _ = session.Insert(&batchUserSpeciality)

	var batchUserCards []model.UserHeroCards
	batchUserCards = append(batchUserCards, model.UserHeroCards{
		UserId:         userId,
		CardId:         34,
		UnlockQuantity: 1,
	})
	batchUserCards = append(batchUserCards, model.UserHeroCards{
		UserId:         userId,
		CardId:         1013,
		UnlockQuantity: 2,
	})
	batchUserCards = append(batchUserCards, model.UserHeroCards{
		UserId:         userId,
		CardId:         1046,
		UnlockQuantity: 2,
	})
	batchUserCards = append(batchUserCards, model.UserHeroCards{
		UserId:         userId,
		CardId:         1024,
		UnlockQuantity: 2,
	})
	batchUserCards = append(batchUserCards, model.UserHeroCards{
		UserId:         userId,
		CardId:         1004,
		UnlockQuantity: 2,
	})
	batchUserCards = append(batchUserCards, model.UserHeroCards{
		UserId:         userId,
		CardId:         1053,
		UnlockQuantity: 2,
	})
	_, _ = session.Insert(&batchUserCards)

	var batchUserPassive []model.UserPassive
	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    1,
		Level:     1,
		PassiveId: 25,
	})
	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    1,
		Level:     1,
		PassiveId: 26,
	})
	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    1,
		Level:     1,
		PassiveId: 27,
	})

	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    2,
		Level:     1,
		PassiveId: 1,
	})
	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    2,
		Level:     1,
		PassiveId: 2,
	})
	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    2,
		Level:     1,
		PassiveId: 3,
	})

	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    3,
		Level:     1,
		PassiveId: 7,
	})
	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    3,
		Level:     1,
		PassiveId: 8,
	})
	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    3,
		Level:     1,
		PassiveId: 9,
	})

	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    4,
		Level:     1,
		PassiveId: 13,
	})
	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    4,
		Level:     1,
		PassiveId: 14,
	})
	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    4,
		Level:     1,
		PassiveId: 15,
	})

	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    5,
		Level:     1,
		PassiveId: 19,
	})
	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    5,
		Level:     1,
		PassiveId: 20,
	})
	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    5,
		Level:     1,
		PassiveId: 21,
	})

	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    6,
		Level:     1,
		PassiveId: 31,
	})
	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    6,
		Level:     1,
		PassiveId: 32,
	})
	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    6,
		Level:     1,
		PassiveId: 33,
	})
	_, _ = session.Insert(&batchUserPassive)

	var batchLog []model.UserHeroUpgradeRewardLog
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 1,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 2,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 16,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 17,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 31,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 32,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 46,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 47,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 61,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 62,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 76,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 77,
	})

	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 3,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 18,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 33,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 48,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 63,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 78,
	})

	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 4,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 19,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 34,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 49,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 64,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 79,
	})

	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 5,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 6,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 7,
	})

	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 20,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 21,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 22,
	})

	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 35,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 36,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 37,
	})

	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 50,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 51,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 52,
	})

	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 65,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 66,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 67,
	})

	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 80,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 81,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 82,
	})

	_, _ = session.Insert(&batchLog)

	_ = session.Commit()
}

func Reward5(userId int64, exp float64) {
	session := base.Engine.NewSession()
	_ = session.Begin()

	_, _ = session.Table("user_hero").
		Where(builder.Eq{"user_id": userId}).
		MustCols("rank", "experience_pool", "total_offset").
		Update(&model.UserHeroRankAndPropertyOffset{
			Rank:           5,
			ExperiencePool: exp,
			TotalOffset:    2,
		})
	var batchUserSpeciality []model.UserSpeciality
	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       1,
		Level:        1,
		SpecialityId: 21,
		TakeAlong:    true,
	})
	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       2,
		Level:        1,
		SpecialityId: 1,
		TakeAlong:    true,
	})
	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       3,
		Level:        1,
		SpecialityId: 6,
		TakeAlong:    true,
	})
	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       4,
		Level:        1,
		SpecialityId: 11,
		TakeAlong:    true,
	})
	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       5,
		Level:        1,
		SpecialityId: 16,
		TakeAlong:    true,
	})
	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       6,
		Level:        1,
		SpecialityId: 26,
		TakeAlong:    true,
	})

	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       1,
		Level:        2,
		SpecialityId: 22,
		TakeAlong:    true,
	})
	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       1,
		Level:        2,
		SpecialityId: 23,
		TakeAlong:    false,
	})

	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       2,
		Level:        2,
		SpecialityId: 3,
		TakeAlong:    true,
	})
	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       2,
		Level:        2,
		SpecialityId: 2,
		TakeAlong:    false,
	})

	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       3,
		Level:        2,
		SpecialityId: 8,
		TakeAlong:    true,
	})
	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       3,
		Level:        2,
		SpecialityId: 7,
		TakeAlong:    false,
	})

	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       4,
		Level:        2,
		SpecialityId: 13,
		TakeAlong:    true,
	})
	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       4,
		Level:        2,
		SpecialityId: 12,
		TakeAlong:    false,
	})

	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       5,
		Level:        2,
		SpecialityId: 18,
		TakeAlong:    true,
	})
	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       5,
		Level:        2,
		SpecialityId: 17,
		TakeAlong:    false,
	})

	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       6,
		Level:        2,
		SpecialityId: 28,
		TakeAlong:    true,
	})
	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       6,
		Level:        2,
		SpecialityId: 27,
		TakeAlong:    false,
	})
	_, _ = session.Insert(&batchUserSpeciality)

	var batchUserCards []model.UserHeroCards
	batchUserCards = append(batchUserCards, model.UserHeroCards{
		UserId:         userId,
		CardId:         34,
		UnlockQuantity: 1,
	})
	batchUserCards = append(batchUserCards, model.UserHeroCards{
		UserId:         userId,
		CardId:         1013,
		UnlockQuantity: 2,
	})
	batchUserCards = append(batchUserCards, model.UserHeroCards{
		UserId:         userId,
		CardId:         1046,
		UnlockQuantity: 2,
	})
	batchUserCards = append(batchUserCards, model.UserHeroCards{
		UserId:         userId,
		CardId:         1024,
		UnlockQuantity: 2,
	})
	batchUserCards = append(batchUserCards, model.UserHeroCards{
		UserId:         userId,
		CardId:         1004,
		UnlockQuantity: 2,
	})
	batchUserCards = append(batchUserCards, model.UserHeroCards{
		UserId:         userId,
		CardId:         1053,
		UnlockQuantity: 2,
	})
	_, _ = session.Insert(&batchUserCards)

	var batchUserPassive []model.UserPassive
	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    1,
		Level:     1,
		PassiveId: 25,
	})
	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    1,
		Level:     1,
		PassiveId: 26,
	})
	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    1,
		Level:     1,
		PassiveId: 27,
	})

	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    2,
		Level:     1,
		PassiveId: 1,
	})
	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    2,
		Level:     1,
		PassiveId: 2,
	})
	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    2,
		Level:     1,
		PassiveId: 3,
	})

	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    3,
		Level:     1,
		PassiveId: 7,
	})
	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    3,
		Level:     1,
		PassiveId: 8,
	})
	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    3,
		Level:     1,
		PassiveId: 9,
	})

	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    4,
		Level:     1,
		PassiveId: 13,
	})
	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    4,
		Level:     1,
		PassiveId: 14,
	})
	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    4,
		Level:     1,
		PassiveId: 15,
	})

	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    5,
		Level:     1,
		PassiveId: 19,
	})
	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    5,
		Level:     1,
		PassiveId: 20,
	})
	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    5,
		Level:     1,
		PassiveId: 21,
	})

	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    6,
		Level:     1,
		PassiveId: 31,
	})
	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    6,
		Level:     1,
		PassiveId: 32,
	})
	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    6,
		Level:     1,
		PassiveId: 33,
	})
	_, _ = session.Insert(&batchUserPassive)

	var batchLog []model.UserHeroUpgradeRewardLog
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 1,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 2,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 16,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 17,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 31,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 32,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 46,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 47,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 61,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 62,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 76,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 77,
	})

	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 3,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 18,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 33,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 48,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 63,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 78,
	})

	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 4,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 19,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 34,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 49,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 64,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 79,
	})

	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 5,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 6,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 7,
	})

	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 20,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 21,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 22,
	})

	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 35,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 36,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 37,
	})

	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 50,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 51,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 52,
	})

	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 65,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 66,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 67,
	})

	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 80,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 81,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 82,
	})

	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 8,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 23,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 38,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 53,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 68,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 83,
	})

	_, _ = session.Insert(&batchLog)

	_ = session.Commit()
}

func Reward6(userId int64, exp float64) {
	session := base.Engine.NewSession()
	_ = session.Begin()

	_, _ = session.Table("user_hero").
		Where(builder.Eq{"user_id": userId}).
		MustCols("rank", "experience_pool", "total_offset").
		Update(&model.UserHeroRankAndPropertyOffset{
			Rank:           6,
			ExperiencePool: exp,
			TotalOffset:    2,
		})
	var batchUserSpeciality []model.UserSpeciality
	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       1,
		Level:        1,
		SpecialityId: 21,
		TakeAlong:    true,
	})
	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       2,
		Level:        1,
		SpecialityId: 1,
		TakeAlong:    true,
	})
	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       3,
		Level:        1,
		SpecialityId: 6,
		TakeAlong:    true,
	})
	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       4,
		Level:        1,
		SpecialityId: 11,
		TakeAlong:    true,
	})
	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       5,
		Level:        1,
		SpecialityId: 16,
		TakeAlong:    true,
	})
	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       6,
		Level:        1,
		SpecialityId: 26,
		TakeAlong:    true,
	})

	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       1,
		Level:        2,
		SpecialityId: 22,
		TakeAlong:    true,
	})
	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       1,
		Level:        2,
		SpecialityId: 23,
		TakeAlong:    false,
	})

	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       2,
		Level:        2,
		SpecialityId: 3,
		TakeAlong:    true,
	})
	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       2,
		Level:        2,
		SpecialityId: 2,
		TakeAlong:    false,
	})

	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       3,
		Level:        2,
		SpecialityId: 8,
		TakeAlong:    true,
	})
	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       3,
		Level:        2,
		SpecialityId: 7,
		TakeAlong:    false,
	})

	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       4,
		Level:        2,
		SpecialityId: 13,
		TakeAlong:    true,
	})
	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       4,
		Level:        2,
		SpecialityId: 12,
		TakeAlong:    false,
	})

	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       5,
		Level:        2,
		SpecialityId: 18,
		TakeAlong:    true,
	})
	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       5,
		Level:        2,
		SpecialityId: 17,
		TakeAlong:    false,
	})

	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       6,
		Level:        2,
		SpecialityId: 28,
		TakeAlong:    true,
	})
	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       6,
		Level:        2,
		SpecialityId: 27,
		TakeAlong:    false,
	})
	_, _ = session.Insert(&batchUserSpeciality)

	var batchUserCards []model.UserHeroCards
	batchUserCards = append(batchUserCards, model.UserHeroCards{
		UserId:         userId,
		CardId:         34,
		UnlockQuantity: 1,
	})
	batchUserCards = append(batchUserCards, model.UserHeroCards{
		UserId:         userId,
		CardId:         1013,
		UnlockQuantity: 2,
	})
	batchUserCards = append(batchUserCards, model.UserHeroCards{
		UserId:         userId,
		CardId:         1046,
		UnlockQuantity: 2,
	})
	batchUserCards = append(batchUserCards, model.UserHeroCards{
		UserId:         userId,
		CardId:         1024,
		UnlockQuantity: 2,
	})
	batchUserCards = append(batchUserCards, model.UserHeroCards{
		UserId:         userId,
		CardId:         1004,
		UnlockQuantity: 2,
	})
	batchUserCards = append(batchUserCards, model.UserHeroCards{
		UserId:         userId,
		CardId:         1053,
		UnlockQuantity: 2,
	})

	batchUserCards = append(batchUserCards, model.UserHeroCards{
		UserId:         userId,
		CardId:         1039,
		UnlockQuantity: 2,
	})
	batchUserCards = append(batchUserCards, model.UserHeroCards{
		UserId:         userId,
		CardId:         1015,
		UnlockQuantity: 2,
	})
	batchUserCards = append(batchUserCards, model.UserHeroCards{
		UserId:         userId,
		CardId:         1047,
		UnlockQuantity: 2,
	})
	batchUserCards = append(batchUserCards, model.UserHeroCards{
		UserId:         userId,
		CardId:         2023,
		UnlockQuantity: 2,
	})
	batchUserCards = append(batchUserCards, model.UserHeroCards{
		UserId:         userId,
		CardId:         2,
		UnlockQuantity: 1,
	})
	batchUserCards = append(batchUserCards, model.UserHeroCards{
		UserId:         userId,
		CardId:         1057,
		UnlockQuantity: 2,
	})
	_, _ = session.Insert(&batchUserCards)

	var batchUserPassive []model.UserPassive
	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    1,
		Level:     1,
		PassiveId: 25,
	})
	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    1,
		Level:     1,
		PassiveId: 26,
	})
	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    1,
		Level:     1,
		PassiveId: 27,
	})

	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    2,
		Level:     1,
		PassiveId: 1,
	})
	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    2,
		Level:     1,
		PassiveId: 2,
	})
	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    2,
		Level:     1,
		PassiveId: 3,
	})

	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    3,
		Level:     1,
		PassiveId: 7,
	})
	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    3,
		Level:     1,
		PassiveId: 8,
	})
	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    3,
		Level:     1,
		PassiveId: 9,
	})

	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    4,
		Level:     1,
		PassiveId: 13,
	})
	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    4,
		Level:     1,
		PassiveId: 14,
	})
	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    4,
		Level:     1,
		PassiveId: 15,
	})

	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    5,
		Level:     1,
		PassiveId: 19,
	})
	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    5,
		Level:     1,
		PassiveId: 20,
	})
	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    5,
		Level:     1,
		PassiveId: 21,
	})

	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    6,
		Level:     1,
		PassiveId: 31,
	})
	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    6,
		Level:     1,
		PassiveId: 32,
	})
	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    6,
		Level:     1,
		PassiveId: 33,
	})
	_, _ = session.Insert(&batchUserPassive)

	var batchLog []model.UserHeroUpgradeRewardLog
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 1,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 2,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 16,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 17,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 31,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 32,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 46,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 47,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 61,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 62,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 76,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 77,
	})

	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 3,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 18,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 33,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 48,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 63,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 78,
	})

	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 4,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 19,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 34,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 49,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 64,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 79,
	})

	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 5,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 6,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 7,
	})

	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 20,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 21,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 22,
	})

	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 35,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 36,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 37,
	})

	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 50,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 51,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 52,
	})

	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 65,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 66,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 67,
	})

	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 80,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 81,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 82,
	})

	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 8,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 23,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 38,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 53,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 68,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 83,
	})

	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 9,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 24,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 39,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 54,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 69,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 84,
	})

	_, _ = session.Insert(&batchLog)

	_ = session.Commit()
}

func Reward7(userId int64, exp float64) {
	session := base.Engine.NewSession()
	_ = session.Begin()

	_, _ = session.Table("user_hero").
		Where(builder.Eq{"user_id": userId}).
		MustCols("rank", "experience_pool", "total_offset").
		Update(&model.UserHeroRankAndPropertyOffset{
			Rank:           7,
			ExperiencePool: exp,
			TotalOffset:    2,
		})
	var batchUserSpeciality []model.UserSpeciality
	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       1,
		Level:        1,
		SpecialityId: 21,
		TakeAlong:    true,
	})
	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       2,
		Level:        1,
		SpecialityId: 1,
		TakeAlong:    true,
	})
	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       3,
		Level:        1,
		SpecialityId: 6,
		TakeAlong:    true,
	})
	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       4,
		Level:        1,
		SpecialityId: 11,
		TakeAlong:    true,
	})
	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       5,
		Level:        1,
		SpecialityId: 16,
		TakeAlong:    true,
	})
	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       6,
		Level:        1,
		SpecialityId: 26,
		TakeAlong:    true,
	})

	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       1,
		Level:        2,
		SpecialityId: 22,
		TakeAlong:    true,
	})
	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       1,
		Level:        2,
		SpecialityId: 23,
		TakeAlong:    false,
	})

	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       2,
		Level:        2,
		SpecialityId: 3,
		TakeAlong:    true,
	})
	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       2,
		Level:        2,
		SpecialityId: 2,
		TakeAlong:    false,
	})

	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       3,
		Level:        2,
		SpecialityId: 8,
		TakeAlong:    true,
	})
	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       3,
		Level:        2,
		SpecialityId: 7,
		TakeAlong:    false,
	})

	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       4,
		Level:        2,
		SpecialityId: 13,
		TakeAlong:    true,
	})
	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       4,
		Level:        2,
		SpecialityId: 12,
		TakeAlong:    false,
	})

	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       5,
		Level:        2,
		SpecialityId: 18,
		TakeAlong:    true,
	})
	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       5,
		Level:        2,
		SpecialityId: 17,
		TakeAlong:    false,
	})

	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       6,
		Level:        2,
		SpecialityId: 28,
		TakeAlong:    true,
	})
	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       6,
		Level:        2,
		SpecialityId: 27,
		TakeAlong:    false,
	})

	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       1,
		Level:        3,
		SpecialityId: 24,
		TakeAlong:    true,
	})
	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       1,
		Level:        3,
		SpecialityId: 25,
		TakeAlong:    false,
	})

	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       2,
		Level:        3,
		SpecialityId: 4,
		TakeAlong:    true,
	})
	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       2,
		Level:        3,
		SpecialityId: 5,
		TakeAlong:    false,
	})

	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       3,
		Level:        3,
		SpecialityId: 9,
		TakeAlong:    true,
	})
	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       3,
		Level:        3,
		SpecialityId: 10,
		TakeAlong:    false,
	})

	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       4,
		Level:        3,
		SpecialityId: 14,
		TakeAlong:    true,
	})
	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       4,
		Level:        3,
		SpecialityId: 15,
		TakeAlong:    false,
	})

	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       5,
		Level:        3,
		SpecialityId: 19,
		TakeAlong:    true,
	})
	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       5,
		Level:        3,
		SpecialityId: 20,
		TakeAlong:    false,
	})

	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       6,
		Level:        3,
		SpecialityId: 29,
		TakeAlong:    true,
	})
	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       6,
		Level:        3,
		SpecialityId: 30,
		TakeAlong:    false,
	})
	_, _ = session.Insert(&batchUserSpeciality)

	var batchUserCards []model.UserHeroCards
	batchUserCards = append(batchUserCards, model.UserHeroCards{
		UserId:         userId,
		CardId:         34,
		UnlockQuantity: 1,
	})
	batchUserCards = append(batchUserCards, model.UserHeroCards{
		UserId:         userId,
		CardId:         1013,
		UnlockQuantity: 2,
	})
	batchUserCards = append(batchUserCards, model.UserHeroCards{
		UserId:         userId,
		CardId:         1046,
		UnlockQuantity: 2,
	})
	batchUserCards = append(batchUserCards, model.UserHeroCards{
		UserId:         userId,
		CardId:         1024,
		UnlockQuantity: 2,
	})
	batchUserCards = append(batchUserCards, model.UserHeroCards{
		UserId:         userId,
		CardId:         1004,
		UnlockQuantity: 2,
	})
	batchUserCards = append(batchUserCards, model.UserHeroCards{
		UserId:         userId,
		CardId:         1053,
		UnlockQuantity: 2,
	})

	batchUserCards = append(batchUserCards, model.UserHeroCards{
		UserId:         userId,
		CardId:         1039,
		UnlockQuantity: 2,
	})
	batchUserCards = append(batchUserCards, model.UserHeroCards{
		UserId:         userId,
		CardId:         1015,
		UnlockQuantity: 2,
	})
	batchUserCards = append(batchUserCards, model.UserHeroCards{
		UserId:         userId,
		CardId:         1047,
		UnlockQuantity: 2,
	})
	batchUserCards = append(batchUserCards, model.UserHeroCards{
		UserId:         userId,
		CardId:         2023,
		UnlockQuantity: 2,
	})
	batchUserCards = append(batchUserCards, model.UserHeroCards{
		UserId:         userId,
		CardId:         2,
		UnlockQuantity: 1,
	})
	batchUserCards = append(batchUserCards, model.UserHeroCards{
		UserId:         userId,
		CardId:         1057,
		UnlockQuantity: 2,
	})
	_, _ = session.Insert(&batchUserCards)

	var batchUserPassive []model.UserPassive
	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    1,
		Level:     1,
		PassiveId: 25,
	})
	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    1,
		Level:     1,
		PassiveId: 26,
	})
	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    1,
		Level:     1,
		PassiveId: 27,
	})

	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    2,
		Level:     1,
		PassiveId: 1,
	})
	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    2,
		Level:     1,
		PassiveId: 2,
	})
	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    2,
		Level:     1,
		PassiveId: 3,
	})

	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    3,
		Level:     1,
		PassiveId: 7,
	})
	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    3,
		Level:     1,
		PassiveId: 8,
	})
	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    3,
		Level:     1,
		PassiveId: 9,
	})

	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    4,
		Level:     1,
		PassiveId: 13,
	})
	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    4,
		Level:     1,
		PassiveId: 14,
	})
	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    4,
		Level:     1,
		PassiveId: 15,
	})

	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    5,
		Level:     1,
		PassiveId: 19,
	})
	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    5,
		Level:     1,
		PassiveId: 20,
	})
	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    5,
		Level:     1,
		PassiveId: 21,
	})

	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    6,
		Level:     1,
		PassiveId: 31,
	})
	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    6,
		Level:     1,
		PassiveId: 32,
	})
	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    6,
		Level:     1,
		PassiveId: 33,
	})
	_, _ = session.Insert(&batchUserPassive)

	var batchLog []model.UserHeroUpgradeRewardLog
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 1,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 2,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 16,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 17,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 31,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 32,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 46,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 47,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 61,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 62,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 76,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 77,
	})

	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 3,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 18,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 33,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 48,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 63,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 78,
	})

	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 4,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 19,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 34,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 49,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 64,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 79,
	})

	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 5,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 6,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 7,
	})

	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 20,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 21,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 22,
	})

	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 35,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 36,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 37,
	})

	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 50,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 51,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 52,
	})

	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 65,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 66,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 67,
	})

	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 80,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 81,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 82,
	})

	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 8,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 23,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 38,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 53,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 68,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 83,
	})

	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 9,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 24,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 39,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 54,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 69,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 84,
	})

	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 10,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 11,
	})

	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 25,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 26,
	})

	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 40,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 41,
	})

	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 55,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 56,
	})

	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 70,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 71,
	})

	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 85,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 86,
	})

	_, _ = session.Insert(&batchLog)

	_ = session.Commit()
}

func Reward8(userId int64, exp float64) {
	session := base.Engine.NewSession()
	_ = session.Begin()

	_, _ = session.Table("user_hero").
		Where(builder.Eq{"user_id": userId}).
		MustCols("rank", "experience_pool", "total_offset").
		Update(&model.UserHeroRankAndPropertyOffset{
			Rank:           8,
			ExperiencePool: exp,
			TotalOffset:    2,
		})
	var batchUserSpeciality []model.UserSpeciality
	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       1,
		Level:        1,
		SpecialityId: 21,
		TakeAlong:    true,
	})
	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       2,
		Level:        1,
		SpecialityId: 1,
		TakeAlong:    true,
	})
	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       3,
		Level:        1,
		SpecialityId: 6,
		TakeAlong:    true,
	})
	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       4,
		Level:        1,
		SpecialityId: 11,
		TakeAlong:    true,
	})
	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       5,
		Level:        1,
		SpecialityId: 16,
		TakeAlong:    true,
	})
	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       6,
		Level:        1,
		SpecialityId: 26,
		TakeAlong:    true,
	})

	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       1,
		Level:        2,
		SpecialityId: 22,
		TakeAlong:    true,
	})
	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       1,
		Level:        2,
		SpecialityId: 23,
		TakeAlong:    false,
	})

	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       2,
		Level:        2,
		SpecialityId: 3,
		TakeAlong:    true,
	})
	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       2,
		Level:        2,
		SpecialityId: 2,
		TakeAlong:    false,
	})

	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       3,
		Level:        2,
		SpecialityId: 8,
		TakeAlong:    true,
	})
	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       3,
		Level:        2,
		SpecialityId: 7,
		TakeAlong:    false,
	})

	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       4,
		Level:        2,
		SpecialityId: 13,
		TakeAlong:    true,
	})
	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       4,
		Level:        2,
		SpecialityId: 12,
		TakeAlong:    false,
	})

	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       5,
		Level:        2,
		SpecialityId: 18,
		TakeAlong:    true,
	})
	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       5,
		Level:        2,
		SpecialityId: 17,
		TakeAlong:    false,
	})

	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       6,
		Level:        2,
		SpecialityId: 28,
		TakeAlong:    true,
	})
	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       6,
		Level:        2,
		SpecialityId: 27,
		TakeAlong:    false,
	})

	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       1,
		Level:        3,
		SpecialityId: 24,
		TakeAlong:    true,
	})
	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       1,
		Level:        3,
		SpecialityId: 25,
		TakeAlong:    false,
	})

	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       2,
		Level:        3,
		SpecialityId: 4,
		TakeAlong:    true,
	})
	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       2,
		Level:        3,
		SpecialityId: 5,
		TakeAlong:    false,
	})

	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       3,
		Level:        3,
		SpecialityId: 9,
		TakeAlong:    true,
	})
	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       3,
		Level:        3,
		SpecialityId: 10,
		TakeAlong:    false,
	})

	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       4,
		Level:        3,
		SpecialityId: 14,
		TakeAlong:    true,
	})
	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       4,
		Level:        3,
		SpecialityId: 15,
		TakeAlong:    false,
	})

	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       5,
		Level:        3,
		SpecialityId: 19,
		TakeAlong:    true,
	})
	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       5,
		Level:        3,
		SpecialityId: 20,
		TakeAlong:    false,
	})

	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       6,
		Level:        3,
		SpecialityId: 29,
		TakeAlong:    true,
	})
	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       6,
		Level:        3,
		SpecialityId: 30,
		TakeAlong:    false,
	})
	_, _ = session.Insert(&batchUserSpeciality)

	var batchUserCards []model.UserHeroCards
	batchUserCards = append(batchUserCards, model.UserHeroCards{
		UserId:         userId,
		CardId:         34,
		UnlockQuantity: 1,
	})
	batchUserCards = append(batchUserCards, model.UserHeroCards{
		UserId:         userId,
		CardId:         1013,
		UnlockQuantity: 2,
	})
	batchUserCards = append(batchUserCards, model.UserHeroCards{
		UserId:         userId,
		CardId:         1046,
		UnlockQuantity: 2,
	})
	batchUserCards = append(batchUserCards, model.UserHeroCards{
		UserId:         userId,
		CardId:         1024,
		UnlockQuantity: 2,
	})
	batchUserCards = append(batchUserCards, model.UserHeroCards{
		UserId:         userId,
		CardId:         1004,
		UnlockQuantity: 2,
	})
	batchUserCards = append(batchUserCards, model.UserHeroCards{
		UserId:         userId,
		CardId:         1053,
		UnlockQuantity: 2,
	})

	batchUserCards = append(batchUserCards, model.UserHeroCards{
		UserId:         userId,
		CardId:         1039,
		UnlockQuantity: 2,
	})
	batchUserCards = append(batchUserCards, model.UserHeroCards{
		UserId:         userId,
		CardId:         1015,
		UnlockQuantity: 2,
	})
	batchUserCards = append(batchUserCards, model.UserHeroCards{
		UserId:         userId,
		CardId:         1047,
		UnlockQuantity: 2,
	})
	batchUserCards = append(batchUserCards, model.UserHeroCards{
		UserId:         userId,
		CardId:         2023,
		UnlockQuantity: 2,
	})
	batchUserCards = append(batchUserCards, model.UserHeroCards{
		UserId:         userId,
		CardId:         2,
		UnlockQuantity: 1,
	})
	batchUserCards = append(batchUserCards, model.UserHeroCards{
		UserId:         userId,
		CardId:         1057,
		UnlockQuantity: 2,
	})
	_, _ = session.Insert(&batchUserCards)

	var batchUserPassive []model.UserPassive
	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    1,
		Level:     1,
		PassiveId: 25,
	})
	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    1,
		Level:     1,
		PassiveId: 26,
	})
	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    1,
		Level:     1,
		PassiveId: 27,
	})

	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    2,
		Level:     1,
		PassiveId: 1,
	})
	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    2,
		Level:     1,
		PassiveId: 2,
	})
	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    2,
		Level:     1,
		PassiveId: 3,
	})

	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    3,
		Level:     1,
		PassiveId: 7,
	})
	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    3,
		Level:     1,
		PassiveId: 8,
	})
	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    3,
		Level:     1,
		PassiveId: 9,
	})

	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    4,
		Level:     1,
		PassiveId: 13,
	})
	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    4,
		Level:     1,
		PassiveId: 14,
	})
	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    4,
		Level:     1,
		PassiveId: 15,
	})

	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    5,
		Level:     1,
		PassiveId: 19,
	})
	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    5,
		Level:     1,
		PassiveId: 20,
	})
	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    5,
		Level:     1,
		PassiveId: 21,
	})

	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    6,
		Level:     1,
		PassiveId: 31,
	})
	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    6,
		Level:     1,
		PassiveId: 32,
	})
	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    6,
		Level:     1,
		PassiveId: 33,
	})

	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    1,
		Level:     2,
		PassiveId: 28,
	})
	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    1,
		Level:     2,
		PassiveId: 29,
	})
	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    1,
		Level:     2,
		PassiveId: 30,
	})

	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    2,
		Level:     2,
		PassiveId: 4,
	})
	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    2,
		Level:     2,
		PassiveId: 5,
	})
	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    2,
		Level:     2,
		PassiveId: 6,
	})

	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    3,
		Level:     2,
		PassiveId: 10,
	})
	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    3,
		Level:     2,
		PassiveId: 11,
	})
	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    3,
		Level:     2,
		PassiveId: 12,
	})

	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    4,
		Level:     2,
		PassiveId: 16,
	})
	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    4,
		Level:     2,
		PassiveId: 17,
	})
	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    4,
		Level:     2,
		PassiveId: 18,
	})

	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    5,
		Level:     2,
		PassiveId: 22,
	})
	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    5,
		Level:     2,
		PassiveId: 23,
	})
	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    5,
		Level:     2,
		PassiveId: 24,
	})

	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    6,
		Level:     2,
		PassiveId: 34,
	})
	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    6,
		Level:     2,
		PassiveId: 35,
	})
	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    6,
		Level:     2,
		PassiveId: 36,
	})
	_, _ = session.Insert(&batchUserPassive)

	var batchLog []model.UserHeroUpgradeRewardLog
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 1,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 2,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 16,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 17,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 31,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 32,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 46,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 47,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 61,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 62,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 76,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 77,
	})

	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 3,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 18,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 33,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 48,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 63,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 78,
	})

	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 4,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 19,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 34,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 49,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 64,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 79,
	})

	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 5,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 6,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 7,
	})

	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 20,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 21,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 22,
	})

	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 35,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 36,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 37,
	})

	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 50,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 51,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 52,
	})

	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 65,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 66,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 67,
	})

	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 80,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 81,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 82,
	})

	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 8,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 23,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 38,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 53,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 68,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 83,
	})

	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 9,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 24,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 39,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 54,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 69,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 84,
	})

	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 10,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 11,
	})

	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 25,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 26,
	})

	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 40,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 41,
	})

	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 55,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 56,
	})

	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 70,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 71,
	})

	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 85,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 86,
	})

	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 12,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 13,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 14,
	})

	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 27,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 28,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 29,
	})

	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 42,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 43,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 44,
	})

	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 57,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 58,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 59,
	})

	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 72,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 73,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 74,
	})

	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 87,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 88,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 89,
	})

	_, _ = session.Insert(&batchLog)

	_ = session.Commit()
}

func Reward9(userId int64) {
	session := base.Engine.NewSession()
	_ = session.Begin()

	_, _ = session.Table("user_hero").
		Where(builder.Eq{"user_id": userId}).
		MustCols("rank", "experience_pool", "total_offset").
		Update(&model.UserHeroRankAndPropertyOffset{
			Rank:        9,
			TotalOffset: 2,
		})
	var batchUserSpeciality []model.UserSpeciality
	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       1,
		Level:        1,
		SpecialityId: 21,
		TakeAlong:    true,
	})
	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       2,
		Level:        1,
		SpecialityId: 1,
		TakeAlong:    true,
	})
	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       3,
		Level:        1,
		SpecialityId: 6,
		TakeAlong:    true,
	})
	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       4,
		Level:        1,
		SpecialityId: 11,
		TakeAlong:    true,
	})
	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       5,
		Level:        1,
		SpecialityId: 16,
		TakeAlong:    true,
	})
	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       6,
		Level:        1,
		SpecialityId: 26,
		TakeAlong:    true,
	})

	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       1,
		Level:        2,
		SpecialityId: 22,
		TakeAlong:    true,
	})
	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       1,
		Level:        2,
		SpecialityId: 23,
		TakeAlong:    false,
	})

	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       2,
		Level:        2,
		SpecialityId: 3,
		TakeAlong:    true,
	})
	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       2,
		Level:        2,
		SpecialityId: 2,
		TakeAlong:    false,
	})

	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       3,
		Level:        2,
		SpecialityId: 8,
		TakeAlong:    true,
	})
	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       3,
		Level:        2,
		SpecialityId: 7,
		TakeAlong:    false,
	})

	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       4,
		Level:        2,
		SpecialityId: 13,
		TakeAlong:    true,
	})
	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       4,
		Level:        2,
		SpecialityId: 12,
		TakeAlong:    false,
	})

	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       5,
		Level:        2,
		SpecialityId: 18,
		TakeAlong:    true,
	})
	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       5,
		Level:        2,
		SpecialityId: 17,
		TakeAlong:    false,
	})

	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       6,
		Level:        2,
		SpecialityId: 28,
		TakeAlong:    true,
	})
	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       6,
		Level:        2,
		SpecialityId: 27,
		TakeAlong:    false,
	})

	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       1,
		Level:        3,
		SpecialityId: 24,
		TakeAlong:    true,
	})
	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       1,
		Level:        3,
		SpecialityId: 25,
		TakeAlong:    false,
	})

	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       2,
		Level:        3,
		SpecialityId: 4,
		TakeAlong:    true,
	})
	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       2,
		Level:        3,
		SpecialityId: 5,
		TakeAlong:    false,
	})

	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       3,
		Level:        3,
		SpecialityId: 9,
		TakeAlong:    true,
	})
	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       3,
		Level:        3,
		SpecialityId: 10,
		TakeAlong:    false,
	})

	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       4,
		Level:        3,
		SpecialityId: 14,
		TakeAlong:    true,
	})
	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       4,
		Level:        3,
		SpecialityId: 15,
		TakeAlong:    false,
	})

	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       5,
		Level:        3,
		SpecialityId: 19,
		TakeAlong:    true,
	})
	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       5,
		Level:        3,
		SpecialityId: 20,
		TakeAlong:    false,
	})

	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       6,
		Level:        3,
		SpecialityId: 29,
		TakeAlong:    true,
	})
	batchUserSpeciality = append(batchUserSpeciality, model.UserSpeciality{
		UserId:       userId,
		HeroId:       6,
		Level:        3,
		SpecialityId: 30,
		TakeAlong:    false,
	})
	_, _ = session.Insert(&batchUserSpeciality)

	var batchUserCards []model.UserHeroCards
	batchUserCards = append(batchUserCards, model.UserHeroCards{
		UserId:         userId,
		CardId:         34,
		UnlockQuantity: 1,
	})
	batchUserCards = append(batchUserCards, model.UserHeroCards{
		UserId:         userId,
		CardId:         1013,
		UnlockQuantity: 2,
	})
	batchUserCards = append(batchUserCards, model.UserHeroCards{
		UserId:         userId,
		CardId:         1046,
		UnlockQuantity: 2,
	})
	batchUserCards = append(batchUserCards, model.UserHeroCards{
		UserId:         userId,
		CardId:         1024,
		UnlockQuantity: 2,
	})
	batchUserCards = append(batchUserCards, model.UserHeroCards{
		UserId:         userId,
		CardId:         1004,
		UnlockQuantity: 2,
	})
	batchUserCards = append(batchUserCards, model.UserHeroCards{
		UserId:         userId,
		CardId:         1053,
		UnlockQuantity: 2,
	})

	batchUserCards = append(batchUserCards, model.UserHeroCards{
		UserId:         userId,
		CardId:         1039,
		UnlockQuantity: 2,
	})
	batchUserCards = append(batchUserCards, model.UserHeroCards{
		UserId:         userId,
		CardId:         1015,
		UnlockQuantity: 2,
	})
	batchUserCards = append(batchUserCards, model.UserHeroCards{
		UserId:         userId,
		CardId:         1047,
		UnlockQuantity: 2,
	})
	batchUserCards = append(batchUserCards, model.UserHeroCards{
		UserId:         userId,
		CardId:         2023,
		UnlockQuantity: 2,
	})
	batchUserCards = append(batchUserCards, model.UserHeroCards{
		UserId:         userId,
		CardId:         2,
		UnlockQuantity: 1,
	})
	batchUserCards = append(batchUserCards, model.UserHeroCards{
		UserId:         userId,
		CardId:         1057,
		UnlockQuantity: 2,
	})

	batchUserCards = append(batchUserCards, model.UserHeroCards{
		UserId:         userId,
		CardId:         1040,
		UnlockQuantity: 2,
	})
	batchUserCards = append(batchUserCards, model.UserHeroCards{
		UserId:         userId,
		CardId:         1014,
		UnlockQuantity: 1,
	})
	batchUserCards = append(batchUserCards, model.UserHeroCards{
		UserId:         userId,
		CardId:         1048,
		UnlockQuantity: 1,
	})
	batchUserCards = append(batchUserCards, model.UserHeroCards{
		UserId:         userId,
		CardId:         2024,
		UnlockQuantity: 1,
	})
	batchUserCards = append(batchUserCards, model.UserHeroCards{
		UserId:         userId,
		CardId:         2003,
		UnlockQuantity: 2,
	})
	batchUserCards = append(batchUserCards, model.UserHeroCards{
		UserId:         userId,
		CardId:         2055,
		UnlockQuantity: 1,
	})
	_, _ = session.Insert(&batchUserCards)

	var batchUserPassive []model.UserPassive
	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    1,
		Level:     1,
		PassiveId: 25,
	})
	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    1,
		Level:     1,
		PassiveId: 26,
	})
	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    1,
		Level:     1,
		PassiveId: 27,
	})

	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    2,
		Level:     1,
		PassiveId: 1,
	})
	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    2,
		Level:     1,
		PassiveId: 2,
	})
	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    2,
		Level:     1,
		PassiveId: 3,
	})

	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    3,
		Level:     1,
		PassiveId: 7,
	})
	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    3,
		Level:     1,
		PassiveId: 8,
	})
	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    3,
		Level:     1,
		PassiveId: 9,
	})

	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    4,
		Level:     1,
		PassiveId: 13,
	})
	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    4,
		Level:     1,
		PassiveId: 14,
	})
	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    4,
		Level:     1,
		PassiveId: 15,
	})

	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    5,
		Level:     1,
		PassiveId: 19,
	})
	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    5,
		Level:     1,
		PassiveId: 20,
	})
	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    5,
		Level:     1,
		PassiveId: 21,
	})

	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    6,
		Level:     1,
		PassiveId: 31,
	})
	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    6,
		Level:     1,
		PassiveId: 32,
	})
	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    6,
		Level:     1,
		PassiveId: 33,
	})

	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    1,
		Level:     2,
		PassiveId: 28,
	})
	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    1,
		Level:     2,
		PassiveId: 29,
	})
	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    1,
		Level:     2,
		PassiveId: 30,
	})

	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    2,
		Level:     2,
		PassiveId: 4,
	})
	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    2,
		Level:     2,
		PassiveId: 5,
	})
	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    2,
		Level:     2,
		PassiveId: 6,
	})

	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    3,
		Level:     2,
		PassiveId: 10,
	})
	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    3,
		Level:     2,
		PassiveId: 11,
	})
	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    3,
		Level:     2,
		PassiveId: 12,
	})

	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    4,
		Level:     2,
		PassiveId: 16,
	})
	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    4,
		Level:     2,
		PassiveId: 17,
	})
	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    4,
		Level:     2,
		PassiveId: 18,
	})

	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    5,
		Level:     2,
		PassiveId: 22,
	})
	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    5,
		Level:     2,
		PassiveId: 23,
	})
	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    5,
		Level:     2,
		PassiveId: 24,
	})

	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    6,
		Level:     2,
		PassiveId: 34,
	})
	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    6,
		Level:     2,
		PassiveId: 35,
	})
	batchUserPassive = append(batchUserPassive, model.UserPassive{
		UserId:    userId,
		HeroId:    6,
		Level:     2,
		PassiveId: 36,
	})
	_, _ = session.Insert(&batchUserPassive)

	var batchLog []model.UserHeroUpgradeRewardLog
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 1,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 2,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 16,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 17,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 31,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 32,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 46,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 47,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 61,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 62,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 76,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 77,
	})

	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 3,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 18,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 33,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 48,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 63,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 78,
	})

	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 4,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 19,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 34,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 49,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 64,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 79,
	})

	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 5,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 6,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 7,
	})

	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 20,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 21,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 22,
	})

	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 35,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 36,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 37,
	})

	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 50,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 51,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 52,
	})

	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 65,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 66,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 67,
	})

	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 80,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 81,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 82,
	})

	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 8,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 23,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 38,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 53,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 68,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 83,
	})

	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 9,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 24,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 39,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 54,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 69,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 84,
	})

	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 10,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 11,
	})

	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 25,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 26,
	})

	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 40,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 41,
	})

	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 55,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 56,
	})

	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 70,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 71,
	})

	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 85,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 86,
	})

	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 12,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 13,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 14,
	})

	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 27,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 28,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 29,
	})

	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 42,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 43,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 44,
	})

	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 57,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 58,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 59,
	})

	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 72,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 73,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 74,
	})

	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 87,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 88,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 89,
	})

	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 15,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 30,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 45,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 60,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 75,
	})
	batchLog = append(batchLog, model.UserHeroUpgradeRewardLog{
		UserId:              userId,
		HeroUpgradeRewardId: 90,
	})

	_, _ = session.Insert(&batchLog)

	_ = session.Commit()
}

func InitReleaseRank() {
	// 获取用户总量
	count, _ := base.Engine.Count(new(model.User))
	pageSize := 100 // 每次读取100个用户id
	pageCount := int(count) / pageSize
	if int(count)%pageSize > 1 {
		pageCount++
	}

	for i := 0; i < pageCount; i++ {
		var users []model.User
		_ = base.Engine.Cols("id", "tmp_achieve_num").
			Limit(pageSize, i*pageSize).
			Find(&users)
		for _, user := range users {
			if user.TmpAchieveNum > 0 {
				exp := float64(200 * user.TmpAchieveNum)
				if exp < 200 {
					Reward0(user.Id, exp)
				} else if exp < 500 {
					Reward1(user.Id, exp-200)
				} else if exp < 900 {
					Reward2(user.Id, exp-500)
				} else if exp < 1400 {
					Reward3(user.Id, exp-900)
				} else if exp < 2000 {
					Reward4(user.Id, exp-1400)
				} else if exp < 2700 {
					Reward5(user.Id, exp-2000)
				} else if exp < 3500 {
					Reward6(user.Id, exp-2700)
				} else if exp < 4400 {
					Reward7(user.Id, exp-3500)
				} else if exp < 5400 {
					Reward8(user.Id, exp-4400)
				} else {
					Reward9(user.Id)
				}
			}
		}
		fmt.Println(i+1, "/", pageCount)
	}
}

func RebuildRankTable() {
	service.DropAllRank()
	var batch []*model.Rank

	batch = append(batch, &model.Rank{
		Name:       "0",
		Rank:       0,
		Experience: 200,
	})
	batch = append(batch, &model.Rank{
		Name:       "1",
		Rank:       1,
		Experience: 300,
	})
	batch = append(batch, &model.Rank{
		Name:       "2",
		Rank:       2,
		Experience: 400,
	})
	batch = append(batch, &model.Rank{
		Name:       "3",
		Rank:       3,
		Experience: 500,
	})
	batch = append(batch, &model.Rank{
		Name:       "4",
		Rank:       4,
		Experience: 600,
	})
	batch = append(batch, &model.Rank{
		Name:       "5",
		Rank:       5,
		Experience: 700,
	})
	batch = append(batch, &model.Rank{
		Name:       "6",
		Rank:       6,
		Experience: 800,
	})
	batch = append(batch, &model.Rank{
		Name:       "7",
		Rank:       7,
		Experience: 900,
	})
	batch = append(batch, &model.Rank{
		Name:       "8",
		Rank:       8,
		Experience: 1000,
	})
	batch = append(batch, &model.Rank{
		Name:       "9",
		Rank:       9,
		Experience: 0,
	})

	service.InsertRank(batch)
}

func RebuildHeroUpgradeRewardTable() {
	service.DropAllHeroUpgradeReward()
	var heroId int64
	var batch []*model.HeroUpgradeReward

	heroId, _ = service.GetHeroId(0)
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        1,
		Category:    1,
		ObjectId:    int(enum.SpecialityBeastTraining),
		ObjectValue: 2,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        1,
		Category:    1,
		ObjectId:    int(enum.SpecialityNature),
		ObjectValue: 2,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        2,
		Category:    4,
		ObjectValue: 1,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        3,
		Category:    2,
		ObjectId:    34,
		ObjectValue: 1,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        4,
		Category:    3,
		ObjectId:    int(enum.HeroPassive1),
		ObjectValue: 1,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        4,
		Category:    3,
		ObjectId:    int(enum.HeroPassive2),
		ObjectValue: 1,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        4,
		Category:    3,
		ObjectId:    int(enum.HeroPassive3),
		ObjectValue: 1,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        5,
		Category:    4,
		ObjectValue: 1,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        6,
		Category:    2,
		ObjectId:    1039,
		ObjectValue: 2,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        7,
		Category:    1,
		ObjectId:    int(enum.SpecialityNegotiate),
		ObjectValue: 3,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        7,
		Category:    1,
		ObjectId:    int(enum.SpecialityMedicine),
		ObjectValue: 3,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        8,
		Category:    3,
		ObjectId:    int(enum.HeroPassive4),
		ObjectValue: 2,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        8,
		Category:    3,
		ObjectId:    int(enum.HeroPassive5),
		ObjectValue: 2,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        8,
		Category:    3,
		ObjectId:    int(enum.HeroPassive6),
		ObjectValue: 2,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        9,
		Category:    2,
		ObjectId:    1040,
		ObjectValue: 2,
	})

	heroId, _ = service.GetHeroId(1)
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        1,
		Category:    1,
		ObjectId:    int(enum.SpecialityCraftyHand),
		ObjectValue: 2,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        1,
		Category:    1,
		ObjectId:    int(enum.SpecialityVerbalTrick),
		ObjectValue: 2,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        2,
		Category:    4,
		ObjectValue: 1,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        3,
		Category:    2,
		ObjectId:    1013,
		ObjectValue: 2,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        4,
		Category:    3,
		ObjectId:    int(enum.HeroPassive1),
		ObjectValue: 1,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        4,
		Category:    3,
		ObjectId:    int(enum.HeroPassive2),
		ObjectValue: 1,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        4,
		Category:    3,
		ObjectId:    int(enum.HeroPassive3),
		ObjectValue: 1,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        5,
		Category:    4,
		ObjectValue: 1,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        6,
		Category:    2,
		ObjectId:    1015,
		ObjectValue: 2,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        7,
		Category:    1,
		ObjectId:    int(enum.SpecialityPerformance),
		ObjectValue: 3,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        7,
		Category:    1,
		ObjectId:    int(enum.SpecialityNegotiate),
		ObjectValue: 3,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        8,
		Category:    3,
		ObjectId:    int(enum.HeroPassive4),
		ObjectValue: 2,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        8,
		Category:    3,
		ObjectId:    int(enum.HeroPassive5),
		ObjectValue: 2,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        8,
		Category:    3,
		ObjectId:    int(enum.HeroPassive6),
		ObjectValue: 2,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        9,
		Category:    2,
		ObjectId:    1014,
		ObjectValue: 1,
	})

	heroId, _ = service.GetHeroId(2)
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        1,
		Category:    1,
		ObjectId:    int(enum.SpecialityDeterrence),
		ObjectValue: 2,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        1,
		Category:    1,
		ObjectId:    int(enum.SpecialityHumanity),
		ObjectValue: 2,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        2,
		Category:    4,
		ObjectValue: 1,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        3,
		Category:    2,
		ObjectId:    1046,
		ObjectValue: 2,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        4,
		Category:    3,
		ObjectId:    int(enum.HeroPassive1),
		ObjectValue: 1,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        4,
		Category:    3,
		ObjectId:    int(enum.HeroPassive2),
		ObjectValue: 1,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        4,
		Category:    3,
		ObjectId:    int(enum.HeroPassive3),
		ObjectValue: 1,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        5,
		Category:    4,
		ObjectValue: 1,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        6,
		Category:    2,
		ObjectId:    1047,
		ObjectValue: 2,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        7,
		Category:    1,
		ObjectId:    int(enum.SpecialityCraftyHand),
		ObjectValue: 3,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        7,
		Category:    1,
		ObjectId:    int(enum.SpecialityMysterious),
		ObjectValue: 3,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        8,
		Category:    3,
		ObjectId:    int(enum.HeroPassive4),
		ObjectValue: 2,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        8,
		Category:    3,
		ObjectId:    int(enum.HeroPassive5),
		ObjectValue: 2,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        8,
		Category:    3,
		ObjectId:    int(enum.HeroPassive6),
		ObjectValue: 2,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        9,
		Category:    2,
		ObjectId:    1048,
		ObjectValue: 1,
	})

	heroId, _ = service.GetHeroId(3)
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        1,
		Category:    1,
		ObjectId:    int(enum.SpecialityComputer),
		ObjectValue: 2,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        1,
		Category:    1,
		ObjectId:    int(enum.SpecialityNegotiate),
		ObjectValue: 2,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        2,
		Category:    4,
		ObjectValue: 1,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        3,
		Category:    2,
		ObjectId:    1024,
		ObjectValue: 2,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        4,
		Category:    3,
		ObjectId:    int(enum.HeroPassive1),
		ObjectValue: 1,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        4,
		Category:    3,
		ObjectId:    int(enum.HeroPassive2),
		ObjectValue: 1,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        4,
		Category:    3,
		ObjectId:    int(enum.HeroPassive3),
		ObjectValue: 1,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        5,
		Category:    4,
		ObjectValue: 1,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        6,
		Category:    2,
		ObjectId:    2023,
		ObjectValue: 2,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        7,
		Category:    1,
		ObjectId:    int(enum.SpecialityHumanity),
		ObjectValue: 3,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        7,
		Category:    1,
		ObjectId:    int(enum.SpecialityVerbalTrick),
		ObjectValue: 3,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        8,
		Category:    3,
		ObjectId:    int(enum.HeroPassive4),
		ObjectValue: 2,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        8,
		Category:    3,
		ObjectId:    int(enum.HeroPassive5),
		ObjectValue: 2,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        8,
		Category:    3,
		ObjectId:    int(enum.HeroPassive6),
		ObjectValue: 2,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        9,
		Category:    2,
		ObjectId:    2024,
		ObjectValue: 1,
	})

	heroId, _ = service.GetHeroId(4)
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        1,
		Category:    1,
		ObjectId:    int(enum.SpecialityMedicine),
		ObjectValue: 2,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        1,
		Category:    1,
		ObjectId:    int(enum.SpecialityBeastTraining),
		ObjectValue: 2,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        2,
		Category:    4,
		ObjectValue: 1,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        3,
		Category:    2,
		ObjectId:    1004,
		ObjectValue: 2,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        4,
		Category:    3,
		ObjectId:    int(enum.HeroPassive1),
		ObjectValue: 1,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        4,
		Category:    3,
		ObjectId:    int(enum.HeroPassive2),
		ObjectValue: 1,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        4,
		Category:    3,
		ObjectId:    int(enum.HeroPassive3),
		ObjectValue: 1,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        5,
		Category:    4,
		ObjectValue: 1,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        6,
		Category:    2,
		ObjectId:    2,
		ObjectValue: 1,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        7,
		Category:    1,
		ObjectId:    int(enum.SpecialityCharm),
		ObjectValue: 3,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        7,
		Category:    1,
		ObjectId:    int(enum.SpecialityComputer),
		ObjectValue: 3,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        8,
		Category:    3,
		ObjectId:    int(enum.HeroPassive4),
		ObjectValue: 2,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        8,
		Category:    3,
		ObjectId:    int(enum.HeroPassive5),
		ObjectValue: 2,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        8,
		Category:    3,
		ObjectId:    int(enum.HeroPassive6),
		ObjectValue: 2,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        9,
		Category:    2,
		ObjectId:    2003,
		ObjectValue: 2,
	})

	heroId, _ = service.GetHeroId(5)
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        1,
		Category:    1,
		ObjectId:    int(enum.SpecialityCharm),
		ObjectValue: 2,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        1,
		Category:    1,
		ObjectId:    int(enum.SpecialityPerformance),
		ObjectValue: 2,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        2,
		Category:    4,
		ObjectValue: 1,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        3,
		Category:    2,
		ObjectId:    1053,
		ObjectValue: 2,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        4,
		Category:    3,
		ObjectId:    int(enum.HeroPassive1),
		ObjectValue: 1,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        4,
		Category:    3,
		ObjectId:    int(enum.HeroPassive2),
		ObjectValue: 1,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        4,
		Category:    3,
		ObjectId:    int(enum.HeroPassive3),
		ObjectValue: 1,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        5,
		Category:    4,
		ObjectValue: 1,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        6,
		Category:    2,
		ObjectId:    1057,
		ObjectValue: 2,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        7,
		Category:    1,
		ObjectId:    int(enum.SpecialityComputer),
		ObjectValue: 3,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        7,
		Category:    1,
		ObjectId:    int(enum.SpecialityNegotiate),
		ObjectValue: 3,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        8,
		Category:    3,
		ObjectId:    int(enum.HeroPassive4),
		ObjectValue: 2,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        8,
		Category:    3,
		ObjectId:    int(enum.HeroPassive5),
		ObjectValue: 2,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        8,
		Category:    3,
		ObjectId:    int(enum.HeroPassive6),
		ObjectValue: 2,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        9,
		Category:    2,
		ObjectId:    2055,
		ObjectValue: 1,
	})

	service.InsertHeroUpgradeReward(batch)
}

func AddCardsReward() {
	var heroId int64
	var batch []*model.HeroUpgradeReward

	heroId, _ = service.GetHeroId(0)
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        3,
		Category:    2,
		ObjectId:    34,
		ObjectValue: 1,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        6,
		Category:    2,
		ObjectId:    1039,
		ObjectValue: 2,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        9,
		Category:    2,
		ObjectId:    1040,
		ObjectValue: 2,
	})

	heroId, _ = service.GetHeroId(1)
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        3,
		Category:    2,
		ObjectId:    1013,
		ObjectValue: 2,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        6,
		Category:    2,
		ObjectId:    1015,
		ObjectValue: 2,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        9,
		Category:    2,
		ObjectId:    1014,
		ObjectValue: 1,
	})

	heroId, _ = service.GetHeroId(2)
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        3,
		Category:    2,
		ObjectId:    1046,
		ObjectValue: 2,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        6,
		Category:    2,
		ObjectId:    1047,
		ObjectValue: 2,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        9,
		Category:    2,
		ObjectId:    1048,
		ObjectValue: 1,
	})

	heroId, _ = service.GetHeroId(3)
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        3,
		Category:    2,
		ObjectId:    1024,
		ObjectValue: 2,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        6,
		Category:    2,
		ObjectId:    2023,
		ObjectValue: 2,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        9,
		Category:    2,
		ObjectId:    2024,
		ObjectValue: 1,
	})

	heroId, _ = service.GetHeroId(4)
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        3,
		Category:    2,
		ObjectId:    1004,
		ObjectValue: 2,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        6,
		Category:    2,
		ObjectId:    2,
		ObjectValue: 1,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        9,
		Category:    2,
		ObjectId:    2003,
		ObjectValue: 2,
	})

	heroId, _ = service.GetHeroId(5)
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        3,
		Category:    2,
		ObjectId:    1053,
		ObjectValue: 2,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        6,
		Category:    2,
		ObjectId:    1057,
		ObjectValue: 2,
	})
	batch = append(batch, &model.HeroUpgradeReward{
		HeroId:      int(heroId),
		Rank:        9,
		Category:    2,
		ObjectId:    2055,
		ObjectValue: 1,
	})

	service.InsertHeroUpgradeReward(batch)
}
