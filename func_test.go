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

func Test_CallFuncName(t *testing.T) {
	log.Debug("[" + t.Name() + "]")
	var T = &Test{1, 2}
	a, err := T.Add(1)
	log.Info("  T.Add(1):%v,%v", a, err)
	b, err := CallFuncName(T, "Add", 1)
	log.Info("  CallFuncName:%v,%v", b, err)
}
