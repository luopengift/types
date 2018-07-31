package types

import "testing"

func TestString(t *testing.T) {
	s := String(1)
	t.Log(s.Int())
	t.Log(s)
}
