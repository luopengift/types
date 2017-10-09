package types

import (
	"fmt"
	"reflect"
)

/*
  Example:
    type Test struct {
        i   int
        j   int
    }

    func (t *Test) Add(i int) (int,error) {
        return t.i+t.j+i, nil
    }

    var T = &Test{1,2}

    CallFuncName(T,"Add",1) // [4 <nil>] <nil>
*/
func CallFuncName(class interface{}, fn string, args ...interface{}) ([]interface{}, error) {
	value := reflect.ValueOf(class)
	method := value.MethodByName(fn)
	if bool(method.Kind() != reflect.Func) {
		return nil, fmt.Errorf("%s is %v can not callable", fn, method.Kind())
	}
	numIn := method.Type().NumIn()
	argsIn := make([]reflect.Value, numIn)
	for i := 0; i < numIn; i++ {
		argsIn[i] = reflect.ValueOf(args[i])
	}
	numOut := method.Type().NumOut()
	argsOut := method.Call(argsIn)
	rets := make([]interface{}, numOut)
	for i := 0; i < numOut; i++ {
		rets[i] = argsOut[i].Interface()
	}
	return rets, nil
}
