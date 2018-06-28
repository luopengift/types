package validate

import (
	"fmt"
	"reflect"
	"strconv"
)

type Max struct {
	FieldName string
	Value     reflect.Value
	MaxValue  float64
}

func MaxInit(valueField reflect.Value, structField reflect.StructField, value string) (Validor, error) {
	maxValue, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return nil, err
	}
	max := &Max{}
	max.FieldName = structField.Name
	max.MaxValue = float64(maxValue)
	max.Value = valueField
	return max, nil
}

func (max *Max) Valid() error {
	var v float64
	var err error
	switch max.Value.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		v = float64(max.Value.Int())
	case reflect.Float32, reflect.Float64:
		v = max.Value.Float()
	case reflect.String, reflect.Slice, reflect.Map, reflect.Array, reflect.Chan:
		v = float64(max.Value.Len())
	default:
		return fmt.Errorf("[%s] max is not support type<%v>", max.FieldName, max.Value.Kind())
	}
	if err != nil {
		return err
	}
	if max.MaxValue < v {
		return fmt.Errorf("[%s] max valid fail,max=%v, value=%v", max.FieldName, max.MaxValue, v)
	}
	return nil
}
