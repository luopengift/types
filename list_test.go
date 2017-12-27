package types

import (
	"testing"
	"fmt"
)

func Test_List(t *testing.T) {
	fmt.Printf("[" + t.Name() + "]")
	list := List{1, 2, 3, "4"}
	fmt.Printf("  %v %v", list, list.String())
	fmt.Printf("  1 contains %v", list.Contains(1))
	fmt.Printf("  4 contains %v", list.Contains(4))
}
