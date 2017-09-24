package types

type Map map[string]interface{}

func NewMap() *Map {
	return new(Map)
}

// Store a <key, value> pair into map
func (m *Map) Put(key string, value interface{}) error {
	(*m)[key] = value
	return nil
}

// Load value according to key
func (m *Map) Get(key string) (interface{}, bool) {
	value, ok := (*m)[key]
	return value, ok
}

// Delete value according to key
func (m *Map) Delete(key string) error {
	delete(*m, key)
	return nil
}

func (m *Map) Keys() []string {
	keyList := []string{}
	for key, _ := range *m {
		keyList = append(keyList, key)
	}
	return keyList
}

func (m *Map) String(key string) string {
	value, ok := m.Get(key)
	if !ok {
		return ""
	}
	s, err := ToString(value)
    if err != nil {
        println(err)
    }
    return s
}

type SortMap struct {
	keys   []string
	values List
}

func NewSortMap() *SortMap {
	return new(SortMap)
}

func (sm *SortMap) String() string {
	m := map[string]interface{}{}
	for index, key := range sm.keys {
		m[key] = sm.values[index]
	}
	s, _ := ToString(m)
    return s
}

func (sm *SortMap) Index(index int) (string, interface{}) {
	return sm.keys[index], sm.values.Index(index)
}
