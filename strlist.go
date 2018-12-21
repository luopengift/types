package types

// StringList StringList
type StringList []string

// Len len
func (s StringList) Len() int {
	return len(s)
}

// Append append to the end
// L.append(object) -- append object to end
func (s StringList) Append(v ...string) {
	s = append(s, v...)
}

// Insert insert
func (s StringList) Insert(idx int, v string) {
	rest := s[idx:]
	s = append(s[0:idx], v)
	s = append(s, rest...)
}

// Pop 指定下标，删除指定的元素，如果删除一个不存在的元素会报错，默认Pop()删除最后一个元素
func (s StringList) Pop(idx ...int) string {
	var index int
	if len(idx) == 0 {
		index = s.Len() - 1
	} else {
		index = idx[0]
	}
	pop := s[index]
	s = append(s[:index], s[index+1:]...)
	return pop
}

// // Remove 根据value删除元素
// func (s StringList) Remove(v string) {
// 	for idx, value := range s {
// 		if value == v {
// 			s = append(s[:idx], s[idx+1:]...)
// 		}
// 	}
// }

// Index index
func (s StringList) Index(v string) int {
	for id, value := range s {
		if value == v {
			return id
		}
	}
	return -1
}

// Count L.count(value) -> integer -- return number of occurrences of value
func (s StringList) Count(v string) int {
	cnt := 0
	for _, value := range s {
		if value == v {
			cnt++
		}
	}
	return cnt
}

// Contains contains
func (s StringList) Contains(v string) bool {
	return s.Count(v) != 0
}

// List to []string
func (s StringList) List() []string {
	return s
}
