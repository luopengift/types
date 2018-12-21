package types

// Int64 int64 field
type Int64 int64

// Int to int
func (i Int64) Int() int {
	return int(i)
}

// Bool convert to bool
func (i Int64) Bool() bool {
	if i.Int64() == 0 {
		return false
	}
	return true
}

// Int64 to int64
func (i Int64) Int64() int64 {
	return int64(i)
}

// Float64 to float64
func (i Int64) Float64() float64 {
	return float64(i)
}

// String to string
func (i Int64) String() string {
	return string(i)
}
