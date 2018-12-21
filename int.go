package types

import "strconv"

// Int int field
type Int int

// Int to int
func (i Int) Int() int {
	return int(i)
}

// Bool convert to bool
func (i Int) Bool() bool {
	if i.Int() == 0 {
		return false
	}
	return true
}

// Int64 to int64
func (i Int) Int64() int64 {
	return int64(i)
}

// Float64 to float64
func (i Int) Float64() float64 {
	return float64(i)
}

// String to string
func (i Int) String() string {
	return strconv.Itoa(i.Int())
}
