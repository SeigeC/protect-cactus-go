package service

import (
	"singo/serializer"
	"singo/serializer/handler"
	"singo/util"
	"time"
)

// OneDayMaxScoreParam 获取一天最大得分
type OneDayMaxScoreParam struct {
	Type int   `json:"type" binding:"required"`
	Time int64 `json:"time"`
}

// GetOneDayMaxScore 获取一天最大得分，默认为今天
func (p *OneDayMaxScoreParam) GetOneDayMaxScore(c *handler.Context) (handler.ActionResponse, error) {
	if p.Time == 0 {
		p.Time = time.Now().Unix()
	}
	if c.User == nil {
		return nil, serializer.ErrParamsMsg("用户未登录")
	}

	h, err := c.User.GetOneDayMaxScore(p.Type, p.Time)
	if err != nil {
		return nil, serializer.ErrDatabase.New(err)
	}
	return h, nil
}

// SpotsTimeParam 获取指定时间段的运动时间
type SpotsTimeParam struct {
	Type  int   `json:"type"`
	STime int64 `json:"start_time" binding:"required"`
	ETime int64 `json:"end_time" binding:"required"`
}

func (p *SpotsTimeParam) GetSpotsTime(c *handler.Context) (handler.ActionResponse, error) {
	historyList, err := c.User.GetAllHistoryByTime(p.STime, p.ETime)
	if err != nil {
		return nil, serializer.ErrDatabase.New(err)
	}
	m := make(map[string]int64, len(historyList))
	for _, history := range historyList {
		key := util.FormatTimeYMD(history.Time)
		m[key] = history.SpotsTime
	}
	return m, nil
}
