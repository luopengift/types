package types

import (
	"testing"
)

func Test_HashMap(t *testing.T) {
    log.Debug("[" + t.Name() + "]")

    hm := HashMap{}
    hm["1"] = 1
    log.Debug("  %#v", hm)
}

func Test_SortMap(t *testing.T) {
	log.Debug("[" + t.Name() + "]")
}
