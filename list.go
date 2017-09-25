package types

type List []interface{}

// 判断一个元素是否在List中
func (l List) Contains(v interface{}) bool {
	return l.Count(v) != 0
}

//count the times of v in list
func (l List) Count(v interface{}) int {
	cnt := 0
	for _, value := range l {
		if value == v {
			cnt += 1
		}
	}
	return cnt
}

func (l *List) String() string {
	s, _ := ToString(*l)
	return s
}
