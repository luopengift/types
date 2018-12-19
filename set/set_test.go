package types

import (
	"fmt"
	"testing"
)

func Test_Set(t *testing.T) {
	fmt.Println("[" + t.Name() + "]")
	set1 := NewSafeSet(1, "2", "3")
	set1.Add(1)
	fmt.Printf("  %v\n", set1.Contains(1))
	set2 := NewSafeSet("1", "2", "3", "4")
	fmt.Printf("  %v %v\n", set1, set2)
	set3 := NewSafeSet()
	fmt.Printf("  2-1 %v\n", set2.Diff(set1))
	fmt.Printf("  1-2 %v\n", set1.Diff(set2))
	fmt.Printf("  %v\n", set1.Diff(set3))
	fmt.Printf("  %v\n", set3.Len())
}
