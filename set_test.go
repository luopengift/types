package types

import (
	"github.com/luopengift/golibs/logger"
	"testing"
)

func Test_Set(t *testing.T) {
    logger.Warn("=> Testing set...")
	set1 := NewSet(1, "2", "3")
	set1.Add(1)
	logger.Info("%v", set1.Contains(1))
	set2 := NewSet("1", "2", "3", "4")
	logger.Info("%v %v", set1, set2)
	set3 := NewSet()
	logger.Info("2-1 %v", set2.Diff(set1))
	logger.Info("1-2 %v", set1.Diff(set2))
	logger.Info("%v", set1.Diff(set3))
	logger.Info("%v", set3.Len())
}
