package types

import (
	"fmt"
	"testing"
)

func Test_HashMap(t *testing.T) {
	fmt.Println("[" + t.Name() + "]")

	hm := HashMap{}
	hm["1"] = 1
	fmt.Printf("  %#v\n", hm)
	//fmt.Printf("  %T\n", hm.Get("2"))
}

func Test_SortMap(t *testing.T) {
	fmt.Println("[" + t.Name() + "]")
}
