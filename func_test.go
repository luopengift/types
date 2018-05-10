package types

import (
	"fmt"
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
	fmt.Println("[" + t.Name() + "]")
	var T = &Test{1, 2}
	a, err := T.Add(1)
	fmt.Printf("  T.Add(1):%v,%v\n", a, err)
	b, err := CallMethodName(T, "Add", 1)
	fmt.Printf("  CallMethodName:%v,%v\n", b, err)
}

func add(i, j int) int {
	return i + j
}

func Test_CallFuncName(t *testing.T) {
	fmt.Println("[" + t.Name() + "]")
	a, err := CallFuncName(add, 1, 2)
	fmt.Printf("  CallFuncName:%v, %v\n", a, err)
}

func Test_FuncWithTimeout(t *testing.T) {
	fmt.Println("[" + t.Name() + "]")
	a, err := FuncWithTimeout(0, add, 1, 2)
	fmt.Printf("  CallFuncWithTimeout:%v, %v\n", a, err)

}
