package types

// Interface interface field
type Interface struct {
	v interface{}
}

// Int to int
func (v *Interface) Int() int {
	i, _ := ToInt(v.v)
	return i
}

// Float64 to float64
func (v *Interface) Float64() float64 {
	i, _ := ToFloat64(v.v)
	return i
}

// String to string
func (v *Interface) String() string {
	i, _ := ToString(v.v)
	return i
}

// List to list
func (v *Interface) List() []interface{} {
	return v.v.([]interface{})
}

// Map to map
func (v *Interface) Map() map[string]interface{} {
	return v.v.(map[string]interface{})
}
