package types

// Int64 int64 field
type Int64 struct {
	i int64
}

// Int to int
func (i *Int64) Int() int {
	return int(i.i)
}

// Int64 to int64
func (i *Int64) Int64() int64 {
	return i.i
}

// Float64 to float64
func (i *Int64) Float64() float64 {
	return float64(i.i)
}

// String to string
func (i *Int64) String() string {
	return string(i.i)
}
