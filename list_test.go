package types

import (
	"testing"
)

func Test_List(t *testing.T) {
	log.Debug("[" + t.Name() + "]")
	list := NewList()
	list.Append(1, 2, 3, "4")
	log.Info("  %v %v", list, list.Len())
	log.Info("  1 contains %v", list.Contains(1))
	log.Info("  4 contains %v", list.Contains(4))
}
