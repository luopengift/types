## Types Package

warp basic type for golang and offer some extra type transform functions

## 集合

#### 类似于python的set，使用map的key唯一来确保集合元素的唯一。实现的方法有交集，并集，差集等。

### 使用例子:
```
package main

import (
    "github.com/luopengift/golibs/set"
    "fmt"
)

func main() {
    set1 := set.NewSet(1, "2", "3")
    set1.Add(1)
    fmt.Println(set1.Contains(1))
    set2 := set.NewSet("1", "2", "3", "4")
    fmt.Println(set1, set2)
    set3 := set.NewSet()
    fmt.Println("2-1", set2.Diff(set1))
    fmt.Println("1-2", set1.Diff(set2))
    fmt.Println(set1.Diff(set3))
    fmt.Println(set3.Len())
}
```
