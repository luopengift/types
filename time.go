package types

import (
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
