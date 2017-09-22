package types

import (
	"strconv"
)

type Bytes []byte

func (b Bytes) Int() int {
	i, err := strconv.Atoi(b.String())
	if err != nil {
		return 0
	}
	return i

}

func (b Bytes) Int64() int64 {
	i, err := strconv.ParseInt(b.String(), 10, 64)
	if err != nil {
		return 0
	}
	return i

}

func (b Bytes) String() string { return string(b) }
func (b Bytes) Bytes() []byte  { return []byte(b) }
