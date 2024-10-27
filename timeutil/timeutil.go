package timeutil

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"

	"github.com/golang-module/carbon/v2"
)

// Carbon 获取Carbon对象
func Carbon() carbon.Carbon {
	return carbon.NewCarbon().SetTimezone(carbon.PRC)
}

// TimeToCarbon 时间转Carbon
func TimeToCarbon(t time.Time) carbon.Carbon {
	return carbon.CreateFromStdTime(t, carbon.PRC)
}

// StrToCarbon 字符串转Carbon
func StrToCarbon(str string) carbon.Carbon {
	return carbon.NewCarbon().Parse(str, carbon.PRC)
}

// NowCarbon 获取当前时间的Carbon对象
func NowCarbon() carbon.Carbon {
	return carbon.Now(carbon.PRC)
}

// NowTime 获取当前时间
func NowTime() time.Time {
	return carbon.Now().StdTime()
}

// NowUnix 获取当前时间戳
func NowUnix() int64 {
	return carbon.Now().Timestamp()
}

// NowString 转换为当前时间 2021-06-29 23:53:32
func NowString() string {
	return carbon.Now().ToDateTimeString()
}

// NowMillisecondString 转换为当前时间 2021-06-29 23:53:32.010
func NowMillisecondString() string {
	return carbon.Now().ToDateTimeMilliString()
}

// NowMicrosecondString 转换为当前时间 2021-06-29 23:53:32.100000
func NowMicrosecondString() string {
	return carbon.Now().ToDateTimeMicroString()
}

// NowSQLNullTime 获取当前时间的sql.NullTime
func NowSQLNullTime() sql.NullTime {
	return sql.NullTime{
		Time:  time.Now(),
		Valid: true,
	}
}

// TimeToDateTimeString 时间转日期 eg:2021-06-29 23:53:32
func TimeToDateTimeString(t time.Time) string {
	return carbon.CreateFromStdTime(t).ToDateTimeString()
}

// TimeToSQLNullTime 将time.Time转换为sql.NullTime
func TimeToSQLNullTime(t time.Time) sql.NullTime {
	return sql.NullTime{
		Time:  t,
		Valid: true,
	}
}

// TimeToHuman 格式化人类友好时间
func TimeToHuman(t time.Time) string {
	var res = ""
	if t.IsZero() {
		return res
	}
	tt := time.Now().Unix() - t.Unix()
	data := []map[string]any{
		{"key": 31536000, "value": "年"},
		{"key": 2592000, "value": "个月"},
		{"key": 604800, "value": "星期"},
		{"key": 86400, "value": "天"},
		{"key": 3600, "value": "小时"},
		{"key": 60, "value": "分钟"},
		{"key": 1, "value": "秒"},
	}
	for _, v := range data {
		var c = tt / int64(v["key"].(int))
		if c != 0 {
			suffix := "前"
			if c < 0 {
				suffix = "后"
				c = -c
			}
			res = strconv.Itoa(int(c)) + v["value"].(string) + suffix
			break
		}
	}
	return res
}

// TimeToDiffForHumans 时间差转换为人类友好时间
func TimeToDiffForHumans(t time.Time) string {
	return carbon.CreateFromStdTime(t, carbon.PRC).SetLocale("zh-CN").DiffForHumans()
}

// TimeToHumanShow 格式化人类友好展示时间
func TimeToHumanShow(t time.Time) string {
	duration := time.Now().Unix() - t.Unix()
	timeStr := ""
	if duration < 60 {
		timeStr = "刚刚发布"
	} else if duration < 3600 {
		timeStr = fmt.Sprintf("%d分钟前更新", duration/60)
	} else if duration < 86400 {
		timeStr = fmt.Sprintf("%d小时前更新", duration/3600)
	} else if duration < 86400*2 {
		timeStr = "昨天更新"
	} else {
		timeStr = time.Unix(t.Unix(), 00).Format("2006.01.02") + "日更新"
	}
	return timeStr
}

// StrToTime 字符串转时间
func StrToTime(str string) time.Time {
	return carbon.NewCarbon().Parse(str, carbon.PRC).StdTime()
}

// RFC3339 格式化时间 RFC3339
func RFC3339(tt time.Time) string {
	if tt.IsZero() {
		return ""
	}
	return tt.Format(time.RFC3339)
}
