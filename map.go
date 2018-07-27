package types

// Map map
type Map = map[string]interface{}

// HashMap map
type HashMap Map

// Parse Implement Configer interface.
func (m HashMap) Parse(v interface{}) error {
	return Format(m, v)
}

// SortMap sort Map
type SortMap struct {
	keys   []string
	values List
}

// NewSortMap sortmap
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

// Index index
func (sm *SortMap) Index(index int) (string, interface{}) {
	return sm.keys[index], sm.values[index]
}
