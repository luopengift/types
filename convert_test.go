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
	log.Info("  %#T => %#v", T_IntNum, ToInt(T_IntNum))
	log.Info("  %#T => %#v", T_Int64Num, ToInt(T_Int64Num))
	log.Info("  %#T => %#v", T_Float64Num, ToInt(T_Float64Num))
	log.Info("  %#T => %#v", T_NumString, ToInt(T_NumString))
	log.Info("  %#T => %#v", T_Byte, ToInt(T_Byte))
	log.Info("  %#T => %#v", T_Map, ToInt(T_Map))
	log.Info("  %#T => %#v", T_List, ToInt(T_List))

}

func Test_ToString(t *testing.T) {
	log.Debug("[" + t.Name() + "]")
	log.Info("  %#T => %#v", T_IntNum, ToString(T_IntNum))
	log.Info("  %#T => %#v", T_Int64Num, ToString(T_Int64Num))
	log.Info("  %#T => %#v", T_Float64Num, ToString(T_Float64Num))
	log.Info("  %#T => %#v", T_NumString, ToString(T_NumString))
	log.Info("  %#T => %#v", T_Byte, ToString(T_Byte))
	log.Info("  %#T => %#v", T_Map, ToString(T_Map))
	log.Info("  %#T => %#v", T_List, ToString(T_List))
}

func Test_ToJSON(t *testing.T) {
	log.Debug("[" + t.Name() + "]")
	log.Info("  %#T => %#v", T_IntNum, ToJSON(T_IntNum))
	log.Info("  %#T => %#v", T_Int64Num, ToJSON(T_Int64Num))
	log.Info("  %#T => %#v", T_Float64Num, ToJSON(T_Float64Num))
	log.Info("  %#T => %#v", T_NumString, ToJSON(T_NumString))
	log.Info("  %#T => %#v", T_Byte, ToJSON(T_Byte))
	log.Info("  %#T => %#v", T_Map, ToJSON(T_Map))
	log.Info("  %#T => %#v", T_List, ToJSON(T_List))
}

func Test_ToBytes(t *testing.T) {
	log.Debug("[" + t.Name() + "]")
	log.Info("  %#T => %#v", T_IntNum, ToBytes(T_IntNum))
	log.Info("  %#T => %#v", T_Int64Num, ToBytes(T_Int64Num))
	log.Info("  %#T => %#v", T_Float64Num, ToBytes(T_Float64Num))
	log.Info("  %#T => %#v", T_NumString, ToBytes(T_NumString))
	log.Info("  %#T => %#v", T_Byte, ToBytes(T_Byte))
	log.Info("  %#T => %#v", T_Map, ToBytes(T_Map))
	log.Info("  %#T => %#v", T_List, ToBytes(T_List))
}
