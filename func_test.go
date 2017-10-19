package types

import (
	"testing"
)

type Test struct {
	i int
	j int
}

func (t *Test) Add(i int) (int, error) {
	return t.i + t.j + i, nil
}

func Test_CallMethodName(t *testing.T) {
	log.Debug("[" + t.Name() + "]")
	var T = &Test{1, 2}
	a, err := T.Add(1)
	log.Info("  T.Add(1):%v,%v", a, err)
	b, err := CallMethodName(T, "Add", 1)
	log.Info("  CallMethodName:%v,%v", b, err)
}

func add(i, j int) int {
	return i + j
}

func Test_CallFuncName(t *testing.T) {
	log.Debug("[" + t.Name() + "]")
	a, err := CallFuncName(add, 1, 2)
	log.Info("  CallFuncName:%v,%v", a, err)
}

func Test_FuncWithTimeout(t *testing.T) {
	log.Debug("[" + t.Name() + "]")
	a, err := FuncWithTimeout(0, add, 1, 2)
	log.Info("  CallFuncWithTimeout:%v,%v", a, err)

}
