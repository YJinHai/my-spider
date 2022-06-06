package util

import "time"

const TimeZone = "Asia/Shanghai"

const Layout = "2006-01-02 15:04:05" //时间常量

type Times struct {
}

var TimeLoc *time.Location

func init() {
	TimeLoc, _ = time.LoadLocation(TimeZone)
}

//设置时区
func (t *Times) Location() *time.Location {
	loc, _ := time.LoadLocation(TimeZone)
	return loc
}

func (t *Times) StringToTimestamp(timeStr string) int64 {
	loc := t.Location()
	t1, _ := time.ParseInLocation("2006-01-02 15:04:05", timeStr, loc)
	return t1.In(loc).Unix()
}

func (t *Times) StringToTime(str string) time.Time {
	loc := t.Location()
	ti, _ := time.ParseInLocation(Layout, str, loc)
	return ti
}
