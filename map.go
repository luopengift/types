package types

type Map = map[string]interface{}

type HashMap Map

// Implement Configer interface.
func (hm HashMap) Parse(v interface{}) error {
	return Format(hm, v)
}

type SortMap struct {
	keys   []string
	values List
}

func NewSortMap() *SortMap {
	return new(SortMap)
}

func (sm *SortMap) String() string {
	m := Map{}
	for index, key := range sm.keys {
		m[key] = sm.values[index]
	}
	s, _ := ToString(m)
	return s
}

func (sm *SortMap) Index(index int) (string, interface{}) {
	return sm.keys[index], sm.values[index]
}
