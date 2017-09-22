package types

import (
    "testing"
    "github.com/luopengift/golibs/logger"
)

func Test_list(t *testing.T) {
    logger.Warn("=> Test list")
    list := NewList()
    logger.Info("%v", list.Len())
    list.Append(1,2,3)
    logger.Info("%v %v", list, list.Len())
    logger.Info("1 contains %v", list.Contains(1))
    logger.Info("4 contains %v", list.Contains(4))
}
