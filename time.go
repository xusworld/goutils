package goutils

import "time"

const (
	NumMonthFormat  = "200601"
	NumDateFormat   = "20060102"
	NumHMDFormat    = "150405"
	NumHMFormat     = "1504"
	DateFormat      = "2006-01-02"
	TimeFormat      = "2006-01-02 15:04:05"
	TimeFormatMilli = "2006-01-02 15:04:05.000"
	TimeFormatMicro = "2006-01-02 15:04:05.000000"
	TimeFormatNano  = "2006-01-02 15:04:05.000000000"
)

// GetNow 获取当前时间戳(s)
func GetNow() int64 {
	return time.Now().Unix()
}

// GetNowMs 获取当前时间戳(ms)
func GetNowMs() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

// GetNowUs 获取当前时间戳(us)
func GetNowUs() int64 {
	return time.Now().UnixNano() / int64(time.Microsecond)
}

// GetNowUs 获取当前时间(2006-01-02 15:04:05)
func GetNowString() string {
	return time.Now().Format(TimeFormat)
}

// GetDayBeginTime 获取当天开始时间
func GetDayBeginTime(t time.Time) time.Time {
	y, m, d := t.Date()
	n := time.Date(y, m, d, 0, 0, 0, 0, time.Local)
	return n
}

// GetDayEndTime 获取当天结束时间
func GetDayEndTime(t time.Time) time.Time {
	y, m, d := t.Date()
	n := time.Date(y, m, d, 23, 59, 59, 999999999, time.Local)
	return n
}

// GetDaySecs 获取指定时间是当天的第几秒
func GetDaySecs(t time.Time) int64 {
	return t.Unix() - GetDayBeginTime(t).Unix()
}

// GetWeekDate 获取当周某天
func GetWeekDate(t time.Time, w time.Weekday) time.Time {
	d := w - t.Weekday()
	// 在中国，周日是最后一天，因此针对周日做个修正
	if w == time.Sunday && d != 0 {
		d += 7
	} else if t.Weekday() == time.Sunday && d != 0 {
		d -= 7
	}
	return t.AddDate(0, 0, int(d))
}

// GetMondayDateStr 本周一的年月日， 如20180402
func GetMondayDateStr(tm time.Time) string {
	return GetWeekDate(tm, time.Monday).Format(NumDateFormat)
}

// GetSunDayDateStr 本周日的年月日， 如20180408
func GetSunDayDateStr(tm time.Time) string {
	return GetWeekDate(tm, time.Sunday).Format(NumDateFormat)
}

// IsLeapYear 判断是否是闰年
func IsLeapYear(year int) bool {
	if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
		return true
	}
	return false
}

// IsOneDay 判断是否同一天
func IsOneDay(time1, time2 int64) bool {
	tm1 := time.Unix(time1, 0)
	tm2 := time.Unix(time2, 0)
	if tm1.Format("20060102") != tm2.Format("20060102") {
		return false
	}
	return true
}

// IsOneHour 判断是否同一小时
func IsOneHour(time1, time2 int64) bool {
	tm1 := time.Unix(time1, 0)
	tm2 := time.Unix(time2, 0)
	if tm1.Format("20060102 03 PM") != tm2.Format("20060102 03 PM") {
		return false
	}
	return true
}

// IsOneMinute 判断是否同一分钟
func IsOneMinute(time1, time2 int64) bool {
	tm1 := time.Unix(time1, 0)
	tm2 := time.Unix(time2, 0)
	if tm1.Format("20060102 03:04 PM") != tm2.Format("20060102 03:04 PM") {
		return false
	}
	return true
}

// IsOneWeek 判断是否同一周
func IsOneWeek(time1, time2 int64) bool {
	tm1 := time.Unix(time1, 0)
	tm2 := time.Unix(time2, 0)
	if GetMondayDateStr(tm1) != GetMondayDateStr(tm2) {
		return false
	}
	return true
}

// IsOneMonth 判断是否同一个月
func IsOneMonth(time1, time2 int64) bool {
	tm1 := time.Unix(time1, 0)
	tm2 := time.Unix(time2, 0)
	if tm1.Format("200601") != tm2.Format("200601") {
		return false
	}
	return true
}

// IsOneYear 判断是否同一年
func IsOneYear(time1, time2 int64) bool {
	tm1 := time.Unix(time1, 0)
	tm2 := time.Unix(time2, 0)
	if tm1.Format("2006") != tm2.Format("2006") {
		return false
	}
	return true
}

// TimeCount
// 时间统计类，用于统计时间间隔
type TimeCount struct {
	int64
}

// Set 开始计时
func (self *TimeCount) Set() {
	self.int64 = time.Now().UnixNano()
}

// Get 返回从开始到现在的时间间隔(s)
func (self *TimeCount) Get() uint32 {
	return uint32((time.Now().UnixNano() - self.int64) / int64(time.Second))
}

// GetMs 返回从开始到现在的时间间隔(ms)
func (self *TimeCount) GetMs() uint32 {
	return uint32((time.Now().UnixNano() - self.int64) / int64(time.Millisecond))
}

// GetUs 返回从开始到现在的时间间隔(us)
func (self *TimeCount) GetUs() uint32 {
	return uint32((time.Now().UnixNano() - self.int64) / int64(time.Microsecond))
}

func NewTimeCount() (t *TimeCount) {
	t = new(TimeCount)
	t.Set()
	return t
}

