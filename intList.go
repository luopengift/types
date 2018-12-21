package types

// IntList intList
type IntList []int

// Len len
func (i IntList) Len() int {
	return len(i)
}

// Append append to the end
// L.append(object) -- append object to end
func (i IntList) Append(v ...int) {
	i = append(i, v...)
}

// Insert insert
func (i IntList) Insert(idx, v int) {
	rest := i[idx:]
	i = append(i[0:idx], v)
	i = append(i, rest...)
}

// Pop 指定下标，删除指定的元素，如果删除一个不存在的元素会报错，默认Pop()删除最后一个元素
func (i IntList) Pop(id ...int) int {
	var index int
	if len(id) == 0 {
		index = i.Len() - 1
	} else {
		index = id[0]
	}
	pop := (i)[index]
	i = append((i)[:index], (i)[index+1:]...)
	return pop
}

// // Remove 根据value删除元素
// func (i IntList) Remove(v int) {
// 	for idx, value := range i {
// 		if value == v {
// 			i = append(i[:idx], i[idx+1:]...)
// 		}
// 	}
// }

// Index index
func (i IntList) Index(v int) int {
	for idx, value := range i {
		if value == v {
			return idx
		}
	}
	return -1
}

// Count L.count(value) -> integer -- return number of occurrences of value
func (i IntList) Count(v int) int {
	cnt := 0
	for _, value := range i {
		if value == v {
			cnt++
		}
	}
	return cnt
}

// Contains contains
func (i IntList) Contains(v int) bool {
	return i.Count(v) != 0
}

// List to []int
func (i IntList) List() []int {
	return i
}
