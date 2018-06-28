package validate

import (
	"fmt"
	"reflect"
	"regexp"
)

type Regexp struct {
	FieldName   string
	Value       reflect.Value
	RegexpValue string
}

func RegexpInit(valueField reflect.Value, structField reflect.StructField, value string) (Validor, error) {
	re := &Regexp{}
	re.FieldName = structField.Name
	re.RegexpValue = value
	re.Value = valueField
	return re, nil
}

func (re *Regexp) Valid() error {
	var v string
	switch re.Value.Kind() {
	case reflect.String:
		v = re.Value.String()
	default:
		return fmt.Errorf("[%s] re is not support type<%v>", re.FieldName, re.Value.Kind())
	}
	reg, err := regexp.Compile(re.RegexpValue)
	if err != nil {
		return err
	}
	if ok := reg.MatchString(v); !ok {
		return fmt.Errorf("[%s] re valid fail,re=%v, value=%v", re.FieldName, re.RegexpValue, v)
	}
	return nil
}
