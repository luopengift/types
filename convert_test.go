package types

import (
	"fmt"
	"github.com/luopengift/golibs/logger"
	"testing"
)

func Test_ToInt(t *testing.T) {
	logger.SetTimeFormat("")
	logger.Error("ToInt test")
	fmt.Println(ToInt("1"), ToInt("A"))
}
