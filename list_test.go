package types

import (
	"fmt"
	"testing"
)

func Test_List(t *testing.T) {
	fmt.Println("[" + t.Name() + "]")
	list := List{1, 2, 3, "4"}
	fmt.Printf("  1 contains %v\n", list.Contains(1))
	fmt.Printf("  4 contains %v\n", list.Contains(4))
}
