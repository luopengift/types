package validate

import (
	"fmt"
	"reflect"
	"strconv"
)

type Min struct {
	FieldName string
	Value     reflect.Value
	MinValue  float64
}

func MinInit(valueField reflect.Value, structField reflect.StructField, value string) (Validor, error) {
	minValue, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return nil, err
	}
	min := &Min{}
	min.FieldName = structField.Name
	min.MinValue = float64(minValue)
	min.Value = valueField
	return min, nil
}

func (min *Min) Valid() error {
	var v float64
	var err error
	switch min.Value.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		v = float64(min.Value.Int())
	case reflect.Float32, reflect.Float64:
		v = min.Value.Float()
	case reflect.String, reflect.Slice, reflect.Map, reflect.Array, reflect.Chan:
		v = float64(min.Value.Len())
	default:
		return fmt.Errorf("[%s] min is not support type<%v>", min.FieldName, min.Value.Kind())
	}
	if err != nil {
		return err
	}
	if min.MinValue > v {
		return fmt.Errorf("[%s] min valid fail,min=%v, value=%v", min.FieldName, min.MinValue, v)
	}
	return nil
}
