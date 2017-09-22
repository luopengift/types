package types

import (
	"fmt"
	"math"
)

// 浮点数四舍五入，并取前几位
func Round(f float64, n int) float64 {
	pow10_n := math.Pow10(n)
	return math.Trunc((f+0.5/pow10_n)*pow10_n) / pow10_n
}

// 通过interface{}获取字符串
func ToString(v interface{}) string {
	return fmt.Sprintf("%v", v)
}

// 通过interface{}获获取字节数组
func ToBytes(v interface{}) []byte {
    return []byte(fmt.Sprintf("%v",v))
}


