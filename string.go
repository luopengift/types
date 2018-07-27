package types

import "strconv"

// String string
type String struct {
	s string
}

// Int to int
func (s *String) Int() int {
	i, _ := strconv.Atoi(s.s)
	return i
}

// Int64 to int64
func (s *String) Int64() int64 {
	i, _ := strconv.ParseInt(s.s, 10, 64)
	return i
}

// Float64 to float64
func (s *String) Float64() float64 {
	num, _ := strconv.ParseFloat(s.s, 64)
	return num
}

func (s *String) String() string {
	return s.s
}
