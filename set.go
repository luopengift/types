package types

import (
    "sync"
)

type Setter interface {
	Add(e ...interface{}) *Set  //添加一个元素
	Remove(e interface{}) *Set  //删除一个元素
	Clear() error                //清空所有元素
	Contains(e interface{}) bool //
	Elements() []interface{}     //获取元素集合
	Len() int
	Same(other Setter) bool  //是否和其他set一致
	Union(other Setter) *Set //并集
	Inter(other Setter) *Set //交集
	Diff(other Setter) *Set  //差集
}

type Set struct {
	s map[interface{}]bool
    mux *sync.RWMutex
}

func NewSet(e ...interface{}) *Set {
	set := Set{
        s: make(map[interface{}]bool),
        mux: new(sync.RWMutex),
    }
	set.Add(e...)
	return &set
}

func (self *Set) Add(e ...interface{}) *Set {
    self.mux.Lock()
	for _, v := range e {
		self.s[v] = true
	}
    self.mux.Unlock()
    return self
}

func (self *Set) Remove(v interface{}) *Set {
    self.mux.Lock()
	delete(self.s, v)
    self.mux.Unlock()
	return self
}

func (self *Set) Clear() error {
	self.mux.Lock()
    for k,_ := range self.s {
        self.Remove(k)
    }
    self.mux.Unlock()
	return nil
}

func (self *Set) Contains(v interface{}) bool {
    self.mux.RLock()
    defer self.mux.RUnlock()
	return self.s[v]
}

func (self *Set) Elements() []interface{} {
	elements := []interface{}{}
	for k := range self.s {
		elements = append(elements, k)
	}
	return elements
}

func (self *Set) Len() int {
	return len(self.s)
}

func (self *Set) Same(o Setter) bool {
	if self.Len() != o.Len() {
		return false
	}
	for k := range self.s {
		if !o.Contains(k) {
			return false
		}
	}
	return true
}

func (self *Set) Union(o Setter) *Set {
	union := NewSet(self.Elements()...) //新创建一个集合,以免影响原集合
	for _, v := range o.Elements() {
		union.Add(v)
	}
	return union
}
func (self *Set) Inter(o Setter) *Set {
	inter := NewSet()
	for _, v := range self.Elements() {
		if o.Contains(v) {
			inter.Add(v)
		}
	}
	return inter
}

func (self *Set) Diff(o Setter) *Set {
	diff := NewSet()
	for _, v := range self.Elements() {
		if !o.Contains(v) {
			diff.Add(v)
		}
	}
	return diff
}
