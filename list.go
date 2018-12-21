package types

// List list
type List []interface{}

// Len len
func (i List) Len() int {
	return len(i)
}

// Append L.append(object) -- append object to end
func (i List) Append(v ...interface{}) {
	i = append(i, v...)
}

// Insert insert
func (i List) Insert(id int, v interface{}) {
	rest := i[id:]
	i = append(i[0:id], v)
	i = append(i, rest...)
}

// Pop 指定下标，删除指定的元素，如果删除一个不存在的元素会报错，默认Pop()删除最后一个元素
func (i List) Pop(id ...int) interface{} {
	var index int
	if len(id) == 0 {
		index = i.Len() - 1
	} else {
		index = id[0]
	}
	pop := i[index]
	i = append(i[:index], i[index+1:]...)
	return pop
}

// // Remove 根据value删除元素
// func (i List) Remove(v interface{}) {
// 	for idx, value := range i {
// 		if value == v {
// 			i = append(i[:idx], i[idx+1:i.Len()]...)
// 		}
// 	}
// }

// Index index
func (i List) Index(v interface{}) int {
	for idx, value := range i {
		if value == v {
			return idx
		}
	}
	return -1
}

// Count L.count(value) -> integer -- return number of occurrences of value
func (i List) Count(v interface{}) int {
	cnt := 0
	for _, value := range i {
		if value == v {
			cnt++
		}
	}
	return cnt
}

// Contains contains
func (i List) Contains(v interface{}) bool {
	return i.Count(v) != 0
}

// List to []interface{}
func (i List) List() []interface{} {
	return i
}
