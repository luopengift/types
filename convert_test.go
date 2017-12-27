package types

import (
	"fmt"
	"testing"
)

var (
	T_IntNum     = int(1)
	T_Int64Num   = int64(1)
	T_Float64Num = float64(1.12345)
	T_NumString  = "1"
	T_Str        = "A"
	T_Byte       = []byte("A")
	//	T_Map        = map[string]string{"A": "b", "a": "B"}
	T_Map  = map[string]interface{}{"a": "a", "b": 1}
	T_List = []interface{}{1, "2", 3.1415}
)

func Test_Round(t *testing.T) {
	fmt.Printf("[" + t.Name() + "]")
	fmt.Printf("  %#T => %#v", T_Float64Num, Round(T_Float64Num, 2))
}

func Test_ToInt(t *testing.T) {
	fmt.Printf("[" + t.Name() + "]")
	i, err := ToInt(T_IntNum)
	fmt.Printf("  %#T => %#v, %v", T_IntNum, i, err)
	i, err = ToInt(T_Int64Num)
	fmt.Printf("  %#T => %#v, %v", T_Int64Num, i, err)
	i, err = ToInt(T_Float64Num)
	fmt.Printf("  %#T => %#v, %v", T_Float64Num, i, err)
	i, err = ToInt(T_NumString)
	fmt.Printf("  %#T => %#v, %v", T_NumString, i, err)
	i, err = ToInt(T_Byte)
	fmt.Printf("  %#T => %#v, %v", T_Byte, i, err)
	i, err = ToInt(T_Map)
	fmt.Printf("  %#T => %#v, %v", T_Map, i, err)
	i, err = ToInt(T_List)
	fmt.Printf("  %#T => %#v, %v", T_List, i, err)
}

func Test_ToString(t *testing.T) {
	fmt.Printf("[" + t.Name() + "]")
	i, err := ToString(T_IntNum)
	fmt.Printf("  %#T => %#v, %v", T_IntNum, i, err)
	i, err = ToString(T_Int64Num)
	fmt.Printf("  %#T => %#v, %v", T_Int64Num, i, err)
	i, err = ToString(T_Float64Num)
	fmt.Printf("  %#T => %#v, %v", T_Float64Num, i, err)
	i, err = ToString(T_NumString)
	fmt.Printf("  %#T => %#v, %v", T_NumString, i, err)
	i, err = ToString(T_Byte)
	fmt.Printf("  %#T => %#v, %v", T_Byte, i, err)
	i, err = ToString(T_Map)
	fmt.Printf("  %#T => %#v, %v", T_Map, i, err)
	i, err = ToString(T_List)
	fmt.Printf("  %#T => %#v, %v", T_List, i, err)
}

func Test_ToMap(t *testing.T) {
	fmt.Printf("[" + t.Name() + "]")
	i, err := ToMap(T_IntNum)
	fmt.Printf("  %#T => %#v, %v", T_IntNum, i, err)
	i, err = ToMap(T_Int64Num)
	fmt.Printf("  %#T => %#v, %v", T_Int64Num, i, err)
	i, err = ToMap(T_Float64Num)
	fmt.Printf("  %#T => %#v, %v", T_Float64Num, i, err)
	i, err = ToMap(T_NumString)
	fmt.Printf("  %#T => %#v, %v", T_NumString, i, err)
	i, err = ToMap(T_Byte)
	fmt.Printf("  %#T => %#v, %v", T_Byte, i, err)
	i, err = ToMap(T_Map)
	fmt.Printf("  %#T => %#v, %v", T_Map, i, err)
	i, err = ToMap(T_List)
	fmt.Printf("  %#T => %#v, %v", T_List, i, err)
}

func Test_ToBytes(t *testing.T) {
	fmt.Printf("[" + t.Name() + "]")
	i, err := ToBytes(T_IntNum)
	fmt.Printf("  %#T => %#v, %v", T_IntNum, i, err)
	i, err = ToBytes(T_Int64Num)
	fmt.Printf("  %#T => %#v, %v", T_Int64Num, i, err)
	i, err = ToBytes(T_Float64Num)
	fmt.Printf("  %#T => %#v, %v", T_Float64Num, i, err)
	i, err = ToBytes(T_NumString)
	fmt.Printf("  %#T => %#v, %v", T_NumString, i, err)
	i, err = ToBytes(T_Byte)
	fmt.Printf("  %#T => %#v, %v", T_Byte, i, err)
	i, err = ToBytes(T_Map)
	fmt.Printf("  %#T => %#v, %v", T_Map, i, err)
	i, err = ToBytes(T_List)
	fmt.Printf("  %#T => %#v, %v", T_List, i, err)
}

type test struct {
	A string
	B int
}

func Test_CmpFormatANDToString(t *testing.T) {

	fmt.Printf("[" + t.Name() + "]")
	tt := test{}
	err := Format(T_Map, &tt)
	fmt.Printf("  Format:%#T => %#v, %v", T_Map, tt, err)
	t_map := `{"A":"a", "B":1}`
	tm, err := ToMap(t_map)
	fmt.Printf("  Format:%#T => %#v, %v", t_map, tm, err)

}
