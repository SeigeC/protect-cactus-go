package model

import (
	"singo/util"

	"gorm.io/gorm"
)

// UserHistory 用户成绩
type UserHistory struct {
	ID        int64
	UserID    int64
	Type      int   // 运动类型
	Score     int64 // 用户得分
	SpotsTime int64 // 用户运动时间
	Time      int64 // 记录时间
}

// GetOneDayMaxScore 获取一天最大得分
func (user *User) GetOneDayMaxScore(t int, tt int64) (*UserHistory, error) {
	var history UserHistory
	st, et := util.GetOneDayTime(tt)
	err := DB.Table("user_score").Where("user_id = ? AND type = ? AND time >= ? AND time < ?", user, t, st, et).Order("score DESC").First(&history).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &history, nil
}

// GetAllHistoryByTime 通过启止时间获取 history
func (user *User) GetAllHistoryByTime(st, et int64, t ...int) ([]*UserHistory, error) {
	var historys []*UserHistory

	db := DB.Table("user_score").Where("user_id = ?", user)
	if len(t) == 1 {
		db.Where("type = ?", t[0])
	}
	err := db.Where("time >= ? AND time < ?", st, et).Find(&historys).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return historys, nil
}
