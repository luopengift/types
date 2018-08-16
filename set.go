package types

import (
	"sync"
)

// Setter set interface
type Setter interface {
	Add(...interface{}) Setter
	Remove(interface{}) Setter
	Contains(interface{}) bool
	Elements() []interface{}
	//Equals(Set) bool
	Len() int
	//Copy() Set
	//Sub(Set) Set
	Diff(Setter) Setter
	Inter(Setter) Setter
	Union(Setter) Setter
	Same(Setter) bool
	Clear() error
}

// UnsafeSet unsafe set
type UnsafeSet map[interface{}]bool

// Add add elements
func (s UnsafeSet) Add(e ...interface{}) Setter {
	for _, v := range e {
		s[v] = true
	}
	return s
}

// Contains contains
func (s UnsafeSet) Contains(v interface{}) bool {
	return s[v]
}

// Remove 删除一个元素
func (s UnsafeSet) Remove(v interface{}) Setter {
	delete(s, v)
	return s
}

// Clear 清空所有元素
func (s UnsafeSet) Clear() error {
	for k := range s {
		s.Remove(k)
	}
	return nil
}

// Elements 获取元素集合
func (s UnsafeSet) Elements() []interface{} {
	elements := []interface{}{}
	for k := range s {
		elements = append(elements, k)
	}
	return elements
}

// Same 是否和其他set一致
func (s UnsafeSet) Same(set Setter) bool {
	if s.Len() != set.Len() {
		return false
	}
	for k := range s.Elements() {
		if !set.Contains(k) {
			return false
		}
	}
	return true
}

// Len len
func (s UnsafeSet) Len() int {
	return len(s)
}

// Union 并集
func (s UnsafeSet) Union(o Setter) Setter {
	union := NewSafeSet(s.Elements()...) //新创建一个集合,以免影响原集合
	for _, v := range o.Elements() {
		union.Add(v)
	}
	return union
}

// Inter 交集
func (s UnsafeSet) Inter(o Setter) Setter {
	inter := NewUnsafeSet()
	for _, v := range s.Elements() {
		if o.Contains(v) {
			inter.Add(v)
		}
	}
	return inter
}

// Diff 差集
func (s UnsafeSet) Diff(o Setter) Setter {
	diff := NewUnsafeSet()
	for _, v := range s.Elements() {
		if !o.Contains(v) {
			diff.Add(v)
		}
	}
	return diff
}

// NewUnsafeSet unsafe set
func NewUnsafeSet(e ...interface{}) Setter {
	set := new(UnsafeSet)
	set.Add(e...)
	return set
}

// SafeSet set
type SafeSet struct {
	s   Setter
	mux *sync.RWMutex
}

// NewSafeSet new set
func NewSafeSet(e ...interface{}) Setter {
	set := SafeSet{
		s:   NewUnsafeSet(e...),
		mux: new(sync.RWMutex),
	}
	return &set
}

// Add 添加一个元素
func (s *SafeSet) Add(e ...interface{}) Setter {
	s.mux.Lock()
	s.s.Add(e...)
	s.mux.Unlock()
	return s
}

// Remove 删除一个元素
func (s *SafeSet) Remove(v interface{}) Setter {
	s.mux.Lock()
	s.s.Remove(v)
	s.mux.Unlock()
	return s
}

// Clear 清空所有元素
func (s *SafeSet) Clear() error {
	s.mux.Lock()
	defer s.mux.Unlock()
	return s.s.Clear()
}

// Contains contains
func (s *SafeSet) Contains(v interface{}) bool {
	s.mux.RLock()
	defer s.mux.RUnlock()
	return s.s.Contains(v)
}

// Elements 获取元素集合
func (s *SafeSet) Elements() []interface{} {
	s.mux.RLock()
	defer s.mux.RUnlock()
	return s.s.Elements()

}

// Len len
func (s *SafeSet) Len() int {
	return s.s.Len()
}

// Same 是否和其他set一致
func (s *SafeSet) Same(set Setter) bool {
	s.mux.RLock()
	defer s.mux.RUnlock()
	return s.s.Same(set)
}

// Union 并集
func (s *SafeSet) Union(o Setter) Setter {
	s.mux.RLock()
	defer s.mux.RUnlock()
	return s.s.Union(o)
}

// Inter 交集
func (s *SafeSet) Inter(set Setter) Setter {
	s.mux.RLock()
	defer s.mux.RUnlock()
	return s.s.Inter(set)
}

// Diff 差集
func (s *SafeSet) Diff(set Setter) Setter {
	s.mux.RLock()
	defer s.mux.RUnlock()
	return s.s.Diff(set)
}
