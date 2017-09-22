package types

import (
	"strconv"
)

type Int64 int64

func (i Int64) Int() int         { return int(i) }
func (i Int64) Int64() int64     { return int64(i) }
func (i Int64) Float64() float64 { return float64(i) }
func (i Int64) String() string   { return strconv.FormatInt(i.Int64(), 10) }
func (i Int64) Bytes() []byte    { return []byte(i.String()) }
