package util

import "time"

// GetOneDayTime 获取一天 0 点和 24 点的时间戳
func GetOneDayTime(t int64) (int64,int64) {
	tt:=time.Unix(t,0)
	stime:=time.Date(tt.Year(), tt.Month(), tt.Day(), 0, 0, 0, 0, tt.Location()).Unix()
	return stime,stime+86400
}

// FormatTimeYMD 格式化时间 2006/04/02
func FormatTimeYMD(t int64) string  {
	tt:=time.Unix(t,0)
	return tt.Format("2006/04/02")
}
