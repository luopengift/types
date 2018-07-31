package types

import "strconv"

// Float64 float64
type Float64 float64

// Int to int
func (f Float64) Int() int {
	return int(f)
}

// Int64 to int64
func (f Float64) Int64() int64 {
	return int64(f)
}

// Float64 to float64
func (f Float64) Float64() float64 {
	return float64(f)
}

// String to string
func (f Float64) String() string {
	return strconv.FormatFloat(f.Float64(), 'E', -1, 64)
}
