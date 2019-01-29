package types

import (
	"fmt"
	"strings"
	"time"
)

var m = map[string]string{
	"%Y": "2006",
	"%M": "01",
	"%D": "02",
	"%h": "15",
	"%m": "04",
	"%s": "05",
}

// Now now
func Now() Time {
	return Time(time.Now())
}

// TimeFormat time string
func TimeFormat(str string) string {
	for k, v := range m {
		str = strings.Replace(str, k, time.Now().Format(v), -1)
	}
	return str
}

var timeFormart = "2006-01-02 15:04:05"

// Time time
type Time time.Time

// UnmarshalJSON xxx
func (t *Time) UnmarshalJSON(data []byte) (err error) {
	d := data[1:20]
	d[10] = ' '
	now, err := time.ParseInLocation("2006-01-02 15:04:05", string(d), time.Local)
	*t = Time(now)
	return
}

// MarshalJSON xxx
func (t Time) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(timeFormart)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, timeFormart)
	b = append(b, '"')
	return b, nil
}

func (t Time) String() string {
	return time.Time(t).Format("2006-01-02 15:04:05")
}

// Int64 int64
func (t Time) Int64() int64 {
	return time.Time(t).UnixNano()
}

// Float64 flot64
func (t Time) Float64() float64 {
	return float64(time.Time(t).UnixNano())
}

// Int int
func (t Time) Int() int {
	return int(time.Time(t).UnixNano())
}

// time duration int64
const (
	Second = int64(time.Second)
	Minute = int64(time.Minute)
	Hour   = int64(time.Hour)
	Day    = Hour * 24
	Month  = Day * 30
	Year   = Month * 12
)

// DurationHuman duration human
func DurationHuman(d time.Duration) string {
	duration := int64(d)
	var end, ret string
	var timePoints = []int64{Year, Month, Day, Hour, Minute, Second}
	var zhUnit = []string{"年", "个月", "天", "小时", "分钟", "秒"}
	if duration < 0 {
		end = "之后"
		duration = duration * -1
	} else {
		end = "以前"
	}
	for i, point := range timePoints {
		if duration > point {
			ret += fmt.Sprintf("%d%s", duration/point, zhUnit[i])
			duration = duration % point
		}
	}
	return ret + end
}
