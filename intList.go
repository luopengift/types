package types

// IntList intList
type IntList []int

// Len len
func (list *IntList) Len() int {
	return len(*list)
}

// Append append to the end
// L.append(object) -- append object to end
func (list *IntList) Append(v int) {
	*list = append(*list, v)
}

// Insert insert
func (list *IntList) Insert(idx, v int) {
	rest := (*list)[idx:]
	*list = append((*list)[0:idx], v)
	*list = append(*list, rest...)
}

// Pop 指定下标，删除指定的元素，如果删除一个不存在的元素会报错，默认Pop()删除最后一个元素
func (list *IntList) Pop(idx ...int) int {
	var index int
	if len(idx) == 0 {
		index = list.Len() - 1
	} else {
		index = idx[0]
	}
	pop := (*list)[index]
	*list = append((*list)[:index], (*list)[index+1:]...)
	return pop
}

// Remove 根据value删除元素
func (list *IntList) Remove(v int) {
	for idx, value := range *list {
		if value == v {
			*list = append((*list)[:idx], (*list)[idx+1:]...)
		}
	}
}

// Index index
func (list *IntList) Index(v int) int {
	for idx, value := range *list {
		if value == v {
			return idx
		}
	}
	return -1
}

// Count L.count(value) -> integer -- return number of occurrences of value
func (list *IntList) Count(v int) int {
	cnt := 0
	for _, value := range *list {
		if value == v {
			cnt++
		}
	}
	return cnt
}

// Contains contains
func (list *IntList) Contains(v int) bool {
	return list.Count(v) != 0
}
