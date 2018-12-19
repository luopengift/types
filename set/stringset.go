package types

// StringSet unsafe set
type StringSet map[string]bool

// NewStringSet new StringSet
func NewStringSet(e ...string) StringSet {
	set := StringSet{}
	set.Add(e...)
	return set
}

// Add add elements
func (s StringSet) Add(e ...string) StringSet {
	for _, v := range e {
		s[v] = true
	}
	return s
}

// Contains contains
func (s StringSet) Contains(v string) bool {
	return s[v]
}

// Remove 删除一个元素
func (s StringSet) Remove(v string) StringSet {
	delete(s, v)
	return s
}

// Clear 清空所有元素
func (s StringSet) Clear() error {
	for k := range s {
		s.Remove(k)
	}
	return nil
}

// Elements 获取元素集合
func (s StringSet) Elements() []string {
	var elements []string
	for k := range s {
		elements = append(elements, k)
	}
	return elements
}

// Same 是否和其他set一致
func (s StringSet) Same(set StringSet) bool {
	if s.Len() != set.Len() {
		return false
	}
	for _, v := range s.Elements() {
		if !set.Contains(v) {
			return false
		}
	}
	return true
}

// Len len
func (s StringSet) Len() int {
	return len(s)
}
