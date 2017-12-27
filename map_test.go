package types

import (
	"fmt"
	"testing"
)

func Test_HashMap(t *testing.T) {
	fmt.Printf("[" + t.Name() + "]")

	hm := HashMap{}
	hm["1"] = 1
	fmt.Printf("  %#v", hm)
}

func Test_SortMap(t *testing.T) {
	fmt.Printf("[" + t.Name() + "]")
}
