package main

type StringIntMap struct{
	data map[string]int
}

func NewStringIntMap() *StringIntMap {
	return &StringIntMap{
		data: make(map[string]int),
	}
}

func (m *StringIntMap) Add(key string, value int) {
	m.data[key] = value
}

func (m *StringIntMap) Remove(key string) {
	delete(m.data, key)
}

func (m *StringIntMap) Copy() map[string]int{
	newMap := make(map[string]int, len(m.data))
	for s, v := range m.data{
		newMap[s] = v
	}
	return newMap
}

func (m *StringIntMap) Exists(key string) bool {
	_, ok := m.data[key] 
	return ok
}

func (m *StringIntMap) Get(key string) (int, bool) {
	v, ok := m.data[key] 
	return v, ok
}