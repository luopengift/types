package types

import (
	"github.com/luopengift/golibs/logger"
	"os"
	"testing"
)

var (
	T_IntNum     = int(1)
	T_Int64Num   = int64(1)
	T_Float64Num = float64(1.12345)
	T_NumString  = "1"
	T_Str        = "A"
	T_Byte       = []byte("A")
	T_Map        = map[string]string{"A": "b", "a": "B"}
	T_List       = []interface{}{1, "2", 3.1415}
	log          = logger.NewLogger("", logger.NULL, os.Stdout)
)

func Test_Round(t *testing.T) {
	log.Debug("[" + t.Name() + "]")
	log.Info("  %#T => %#v", T_Float64Num, Round(T_Float64Num, 2))
}

func Test_ToInt(t *testing.T) {
	log.Debug("[" + t.Name() + "]")
	i, err := ToInt(T_IntNum)
	log.Info("  %#T => %#v, %v", T_IntNum, i, err)
	i, err = ToInt(T_Int64Num)
	log.Info("  %#T => %#v, %v", T_Int64Num, i, err)
	i, err = ToInt(T_Float64Num)
	log.Info("  %#T => %#v, %v", T_Float64Num, i, err)
	i, err = ToInt(T_NumString)
	log.Info("  %#T => %#v, %v", T_NumString, i, err)
	i, err = ToInt(T_Byte)
	log.Info("  %#T => %#v, %v", T_Byte, i, err)
	i, err = ToInt(T_Map)
	log.Info("  %#T => %#v, %v", T_Map, i, err)
	i, err = ToInt(T_List)
	log.Info("  %#T => %#v, %v", T_List, i, err)
}

func Test_ToString(t *testing.T) {
	log.Debug("[" + t.Name() + "]")
	i, err := ToString(T_IntNum)
	log.Info("  %#T => %#v, %v", T_IntNum, i, err)
	i, err = ToString(T_Int64Num)
	log.Info("  %#T => %#v, %v", T_Int64Num, i, err)
	i, err = ToString(T_Float64Num)
	log.Info("  %#T => %#v, %v", T_Float64Num, i, err)
	i, err = ToString(T_NumString)
	log.Info("  %#T => %#v, %v", T_NumString, i, err)
	i, err = ToString(T_Byte)
	log.Info("  %#T => %#v, %v", T_Byte, i, err)
	i, err = ToString(T_Map)
	log.Info("  %#T => %#v, %v", T_Map, i, err)
	i, err = ToString(T_List)
	log.Info("  %#T => %#v, %v", T_List, i, err)
}

func Test_ToMap(t *testing.T) {
	log.Debug("[" + t.Name() + "]")
	i, err := ToMap(T_IntNum)
	log.Info("  %#T => %#v, %v", T_IntNum, i, err)
	i, err = ToMap(T_Int64Num)
	log.Info("  %#T => %#v, %v", T_Int64Num, i, err)
	i, err = ToMap(T_Float64Num)
	log.Info("  %#T => %#v, %v", T_Float64Num, i, err)
	i, err = ToMap(T_NumString)
	log.Info("  %#T => %#v, %v", T_NumString, i, err)
	i, err = ToMap(T_Byte)
	log.Info("  %#T => %#v, %v", T_Byte, i, err)
	i, err = ToMap(T_Map)
	log.Info("  %#T => %#v, %v", T_Map, i, err)
	i, err = ToMap(T_List)
	log.Info("  %#T => %#v, %v", T_List, i, err)
}

func Test_ToBytes(t *testing.T) {
	log.Debug("[" + t.Name() + "]")
	i, err := ToBytes(T_IntNum)
	log.Info("  %#T => %#v, %v", T_IntNum, i, err)
	i, err = ToBytes(T_Int64Num)
	log.Info("  %#T => %#v, %v", T_Int64Num, i, err)
	i, err = ToBytes(T_Float64Num)
	log.Info("  %#T => %#v, %v", T_Float64Num, i, err)
	i, err = ToBytes(T_NumString)
	log.Info("  %#T => %#v, %v", T_NumString, i, err)
	i, err = ToBytes(T_Byte)
	log.Info("  %#T => %#v, %v", T_Byte, i, err)
	i, err = ToBytes(T_Map)
	log.Info("  %#T => %#v, %v", T_Map, i, err)
	i, err = ToBytes(T_List)
	log.Info("  %#T => %#v, %v", T_List, i, err)
}
