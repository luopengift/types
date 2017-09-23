package types

import (
	"testing"
)

func Test_Set(t *testing.T) {
	log.Debug("[" + t.Name() + "]")
	set1 := NewSet(1, "2", "3")
	set1.Add(1)
	log.Info("  %v", set1.Contains(1))
	set2 := NewSet("1", "2", "3", "4")
	log.Info("  %v %v", set1, set2)
	set3 := NewSet()
	log.Info("  2-1 %v", set2.Diff(set1))
	log.Info("  1-2 %v", set1.Diff(set2))
	log.Info("  %v", set1.Diff(set3))
	log.Info("  %v", set3.Len())
}
