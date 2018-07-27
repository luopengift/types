package types

import (
	"sync"
)

// Set set
type Set struct {
	s   map[interface{}]bool
	mux *sync.RWMutex
}

// NewSet new set
func NewSet(e ...interface{}) *Set {
	set := Set{
		s:   make(map[interface{}]bool),
		mux: new(sync.RWMutex),
	}
	set.Add(e...)
	return &set
}

// Add 添加一个元素
func (s *Set) Add(e ...interface{}) *Set {
	s.mux.Lock()
	for _, v := range e {
		s.s[v] = true
	}
	s.mux.Unlock()
	return s
}

// Remove 删除一个元素
func (s *Set) Remove(v interface{}) *Set {
	s.mux.Lock()
	delete(s.s, v)
	s.mux.Unlock()
	return s
}

// Clear 清空所有元素
func (s *Set) Clear() error {
	s.mux.Lock()
	for k := range s.s {
		s.Remove(k)
	}
	s.mux.Unlock()
	return nil
}

// Contains contains
func (s *Set) Contains(v interface{}) bool {
	s.mux.RLock()
	defer s.mux.RUnlock()
	return s.s[v]
}

// Elements 获取元素集合
func (s *Set) Elements() []interface{} {
	elements := []interface{}{}
	for k := range s.s {
		elements = append(elements, k)
	}
	return elements
}

// Len len
func (s *Set) Len() int {
	return len(s.s)
}

// Same 是否和其他set一致
func (s *Set) Same(o *Set) bool {
	if s.Len() != o.Len() {
		return false
	}
	for k := range s.s {
		if !o.Contains(k) {
			return false
		}
	}
	return true
}

// Union 并集
func (s *Set) Union(o *Set) *Set {
	union := NewSet(s.Elements()...) //新创建一个集合,以免影响原集合
	for _, v := range o.Elements() {
		union.Add(v)
	}
	return union
}

// Inter 交集
func (s *Set) Inter(o *Set) *Set {
	inter := NewSet()
	for _, v := range s.Elements() {
		if o.Contains(v) {
			inter.Add(v)
		}
	}
	return inter
}

// Diff 差集
func (s *Set) Diff(o *Set) *Set {
	diff := NewSet()
	for _, v := range s.Elements() {
		if !o.Contains(v) {
			diff.Add(v)
		}
	}
	return diff
}
