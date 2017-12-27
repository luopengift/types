package types

import (
	"fmt"
	"testing"
)

func Test_Set(t *testing.T) {
	fmt.Printf("[" + t.Name() + "]")
	set1 := NewSet(1, "2", "3")
	set1.Add(1)
	fmt.Printf("  %v", set1.Contains(1))
	set2 := NewSet("1", "2", "3", "4")
	fmt.Printf("  %v %v", set1, set2)
	set3 := NewSet()
	fmt.Printf("  2-1 %v", set2.Diff(set1))
	fmt.Printf("  1-2 %v", set1.Diff(set2))
	fmt.Printf("  %v", set1.Diff(set3))
	fmt.Printf("  %v", set3.Len())
}
