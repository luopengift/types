package types

type List []interface{}

func NewList() *List {
    return new(List)
}
// Return the length of list
func (l *List) Len() int {
    return len(*l)
}
// add value into the last of list
func (l *List) Append(value ...interface{}) error {
    *l = append(*l,value...)
    return nil
}

// 根据索引删除元素
func (l *List) Delete(index int) error {
    *l = append((*l)[:index], (*l)[index+1:]...)
    return nil
}

// 索引
func (l *List) Index(index int) interface{} {
    return (*l)[index]
}

// 判断一个元素是否在List中
func (l *List) Contains(v interface{}) bool {
    return l.Count(v) != 0
}

//count the times of v in list 
func (l *List) Count(v interface{}) int {
    cnt := 0
    for _, value := range *l {
        if value == v {
            cnt += 1
        }
    }
    return cnt
}
