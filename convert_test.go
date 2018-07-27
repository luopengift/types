package types

import (
	"fmt"
	"testing"
)

var (
	tIntNum     = int(1)
	tInt64Num   = int64(1)
	tFloat64Num = float64(1.12345)
	tNumString  = "1"
	tStr        = "A"
	tByte       = []byte("A")
	//	tMap        = map[string]string{"A": "b", "a": "B"}
	tMap  = map[string]interface{}{"a": "a", "b": 1}
	tList = []interface{}{1, "2", 3.1415}
)

func Test_Round(t *testing.T) {
	fmt.Println("[" + t.Name() + "]")
	fmt.Printf("  %T => %#v\n", tFloat64Num, Round(tFloat64Num, 2))
}

func Test_ToInt(t *testing.T) {
	fmt.Println("[" + t.Name() + "]")
	i, err := ToInt(tIntNum)
	fmt.Printf("  %T => %#v, %v\n", tIntNum, i, err)
	i, err = ToInt(tInt64Num)
	fmt.Printf("  %T => %#v, %v\n", tInt64Num, i, err)
	i, err = ToInt(tFloat64Num)
	fmt.Printf("  %T => %#v, %v\n", tFloat64Num, i, err)
	i, err = ToInt(tNumString)
	fmt.Printf("  %T => %#v, %v\n", tNumString, i, err)
	i, err = ToInt(tByte)
	fmt.Printf("  %T => %#v, %v\n", tByte, i, err)
	i, err = ToInt(tMap)
	fmt.Printf("  %T => %#v, %v\n", tMap, i, err)
	i, err = ToInt(tList)
	fmt.Printf("  %T => %#v, %v\n", tList, i, err)
}

func Test_ToString(t *testing.T) {
	fmt.Println("[" + t.Name() + "]")
	i, err := ToString(tIntNum)
	fmt.Printf("  %T => %#v, %v\n", tIntNum, i, err)
	i, err = ToString(tInt64Num)
	fmt.Printf("  %T => %#v, %v\n", tInt64Num, i, err)
	i, err = ToString(tFloat64Num)
	fmt.Printf("  %T => %#v, %v\n", tFloat64Num, i, err)
	i, err = ToString(tNumString)
	fmt.Printf("  %T => %#v, %v\n", tNumString, i, err)
	i, err = ToString(tByte)
	fmt.Printf("  %T => %#v, %v\n", tByte, i, err)
	i, err = ToString(tMap)
	fmt.Printf("  %T => %#v, %v\n", tMap, i, err)
	i, err = ToString(tList)
	fmt.Printf("  %T => %#v, %v\n", tList, i, err)
}

func Test_ToMap(t *testing.T) {
	fmt.Printf("[" + t.Name() + "]")
	i, err := ToMap(tIntNum)
	fmt.Printf("  %T => %#v, %v\n", tIntNum, i, err)
	i, err = ToMap(tInt64Num)
	fmt.Printf("  %T => %#v, %v\n", tInt64Num, i, err)
	i, err = ToMap(tFloat64Num)
	fmt.Printf("  %T => %#v, %v\n", tFloat64Num, i, err)
	i, err = ToMap(tNumString)
	fmt.Printf("  %T => %#v, %v\n", tNumString, i, err)
	i, err = ToMap(tByte)
	fmt.Printf("  %T => %#v, %v\n", tByte, i, err)
	i, err = ToMap(tMap)
	fmt.Printf("  %T => %#v, %v\n", tMap, i, err)
	i, err = ToMap(tList)
	fmt.Printf("  %T => %#v, %v\n", tList, i, err)
}

func Test_ToBytes(t *testing.T) {
	fmt.Println("[" + t.Name() + "]")
	i, err := ToBytes(tIntNum)
	fmt.Printf("  %T => %#v, %v\n", tIntNum, i, err)
	i, err = ToBytes(tInt64Num)
	fmt.Printf("  %T => %#v, %v\n", tInt64Num, i, err)
	i, err = ToBytes(tFloat64Num)
	fmt.Printf("  %T => %#v, %v\n", tFloat64Num, i, err)
	i, err = ToBytes(tNumString)
	fmt.Printf("  %T => %#v, %v\n", tNumString, i, err)
	i, err = ToBytes(tByte)
	fmt.Printf("  %T => %#v, %v\n", tByte, i, err)
	i, err = ToBytes(tMap)
	fmt.Printf("  %T => %#v, %v\n", tMap, i, err)
	i, err = ToBytes(tList)
	fmt.Printf("  %T => %#v, %v\n", tList, i, err)
}

type test struct {
	A string
	B int
}

func Test_CmpFormatANDToString(t *testing.T) {

	fmt.Println("[" + t.Name() + "]")
	tt := test{}
	err := Format(tMap, &tt)
	fmt.Printf("  Format:%T => %#v, %v\n", tMap, tt, err)
	tMap := `{"A":"a", "B":1}`
	tm, err := ToMap(tMap)
	fmt.Printf("  Format:%T => %#v, %v\n", tMap, tm, err)

}
