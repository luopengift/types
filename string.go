package types

import (
	"strconv"
)

type String string

func (s String) Int() int {
	i, err := strconv.Atoi(s.String())
	if err != nil {
		return 0
	}
	return i
}

func (s String) Int64() int64 {
	i, err := strconv.ParseInt(s.String(), 10, 64)
	if err != nil {
		return 0
	}
	return i
}
func (s String) Float64() float64 {
	f, err := strconv.ParseFloat(s.String(), 64)
	if err != nil {
		return 0
	}
	return f
}
func (s String) String() string { return string(s) }
func (s String) Bytes() []byte  { return []byte(s) }
