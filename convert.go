package types

import (
	"encoding/json"
	"github.com/luopengift/golibs/logger"
	"math"
	"reflect"
	"strconv"
	"unsafe"
)

// byte => string
// 直接转换底层指针，两者指向的相同的内存，改一个另外一个也会变。
// 效率是string([]byte{})的百倍以上，且转换量越大效率优势越明显。
func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// string => bool , if fail return false
func StringToBoolean(s string) bool {
	b, err := strconv.ParseBool(s)
	if err != nil {
		logger.Warn("[%#v] string to bool error, return false => %v", s, err)
		return false
	}
	return b
}

// string => []byte
// 直接转换底层指针，两者指向的相同的内存，改一个另外一个也会变。
// 效率是string([]byte{})的百倍以上，且转换量越大效率优势越明显。
// 转换之后若没做其他操作直接改变里面的字符，则程序会崩溃。
// 如 b:=String2bytes("xxx"); b[1]='d'; 程序将panic。
func StringToBytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

//string => int, if fail return 0
func StringToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		logger.Warn("[%#v] string to int error, return 0 => %v", s, err)
		return 0
	}
	return i
}

//string => int64, if fail return 0
func StringToInt64(s string) int64 {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		logger.Warn("[%#v] string to int64 error, return 0 => %v", s, err)
		return 0
	}
	return i
}

// string => float64, if fail return 0
func StringToFloat64(s string) float64 {
	i, err := strconv.ParseFloat(s, 64)
	if err != nil {
		logger.Warn("[%#v] string to float64 error, return 0 => %v", s, err)
		return 0
	}
	return i
}

// interface => []byte
func ToBytes(v interface{}) []byte {
	switch value := reflect.ValueOf(v); v.(type) {
	case string:
		return StringToBytes(value.String())
	case []byte:
		return value.Bytes()
	default:
		b, err := json.Marshal(v)
		if err != nil {
			logger.Warn("[%#v] ToBytes Marshal error, return []byte{} => %v", v, err)
			return []byte{}
		}
		return b
	}
}

// interface => string
func ToString(v interface{}) string {
	switch value := reflect.ValueOf(v); v.(type) {
	case string:
		return value.String()
	case []byte:
		return BytesToString(value.Bytes())
	default:
		b, err := json.Marshal(v)
		if err != nil {
			logger.Warn("[%#v] ToString Marshal error, return \"\" => %v", v, err)
			return ""
		}
		return string(b)
	}
}

// interface => int
func ToInt(v interface{}) int {
	switch value := reflect.ValueOf(v); v.(type) {
	case string:
		return StringToInt(value.String())
	case int, int8, int16, int32, int64:
		return int(value.Int())
	case uint, uint8, uint16, uint32, uint64:
		return int(value.Uint())
	case float32, float64:
		return int(value.Float())
	default:
		logger.Warn("[%#v] ToInt unknow type, return 0", v)
		return 0
	}
}

// interface => int64
func ToInt64(v interface{}) int64 {
	switch value := reflect.ValueOf(v); v.(type) {
	case string:
		return StringToInt64(value.String())
	case int, int8, int16, int32, int64:
		return value.Int()
	case uint, uint8, uint16, uint32, uint64:
		return int64(value.Uint())
	case float32, float64:
		return int64(value.Float())
	default:
		logger.Warn("[%#v] ToInt64 unknow type", v)
		return 0
	}
}

// interface => float64
func ToFloat64(v interface{}) float64 {
	switch value := reflect.ValueOf(v); v.(type) {
	case string:
		return StringToFloat64(value.String())
	case int, int8, int16, int32, int64:
		return float64(value.Int())
	case uint, uint8, uint16, uint32, uint64:
		return float64(value.Uint())
	case float32, float64:
		return value.Float()
	default:
		logger.Warn("[%#v] ToFloat64  unknow type", v)
		return 0
	}
}

func ToJSON(v interface{}) map[string]interface{} {
	m := make(map[string]interface{})
	err := Format(v, &m)
	if err != nil {
		logger.Warn("[%#v] ToJSON format error => %v", v, err)
	}
	return m
}

// in format out
// eg: Format(map[string]interface{...}, &Struct{})
func Format(in, out interface{}) error {
	var err error
	if b, err := json.Marshal(in); err == nil {
		err = json.Unmarshal(b, out)
	}
	return err
}

// float64四舍五入，并取前几位
func Round(f float64, n int) float64 {
	pow10_n := math.Pow10(n)
	return math.Trunc((f+0.5/pow10_n)*pow10_n) / pow10_n
}
