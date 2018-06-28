package validate

import (
	"fmt"
	"reflect"
	"strings"
)

type Validor interface {
	Valid() error
}

var iface = reflect.TypeOf(new(Validor)).Elem()

func Validate(v interface{}) error {
	rt := reflect.TypeOf(v)
	rv := reflect.ValueOf(v)
	if rt.Implements(iface) {
		res := rv.MethodByName("Valid").Call(nil)
		if ok := res[0].IsNil(); ok {
			return nil
		}
		return res[0].Interface().(error)
	}
	return valid(rv.Elem(), rt.Elem())

}
func valid(rv reflect.Value, rt reflect.Type) error {

	for i := 0; i < rt.NumField(); i++ {
		structField := rt.Field(i)
		valueField := rv.Field(i)
		validTags := structField.Tag.Get("valid")
		if valueField.Kind() == reflect.Struct {
			//TODO
			//if structField.Type.Implements(iface) {
			//	res := valueField.MethodByName("Valid").Call(nil)
			//	return res[0].Interface().(error)
			//}
			if err := valid(valueField, valueField.Type()); err != nil {
				return err
			}
		}
		if validTags == "" || validTags == "-" {
			continue
		}
		for _, validItem := range strings.Split(validTags, ",") {
			valid, err := Init(valueField, structField, validItem)
			if err != nil {
				return err
			}
			err = valid.Valid()
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func Init(valueField reflect.Value, structField reflect.StructField, validItem string) (Validor, error) {
	v := strings.Split(validItem, "=")
	var value string
	if len(v) == 2 {
		value = v[1]
	}
	switch v[0] {
	case "max":
		return MaxInit(valueField, structField, value)
	case "min":
		return MinInit(valueField, structField, value)
	case "regexp":
		return RegexpInit(valueField, structField, value)
	default:
		return nil, fmt.Errorf("valid key<%s> is not support!", v[0])
	}
}
