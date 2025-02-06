package utils

import (
	"time"
)

// TimeZh 中国习惯的时间格式
const TimeZh = "2006-01-02 15:04:05"

const TimeFormat = "2006-01-02 15:04:05-07"

// GetGmt8Timestamp 获取时区为GMT+8的时间，中国时区
func GetGmt8Timestamp(t ...*time.Time) time.Time {
	now := time.Now()
	if len(t) > 0 && t[0] != nil {
		now = *t[0]
	}

	location, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		return now
	}

	return now.In(location)
}
