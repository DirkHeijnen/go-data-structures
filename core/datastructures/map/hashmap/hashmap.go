package advent_of_code

import "fmt"

type HashMap struct {
	data map[interface{}]interface{}
}

func NewMap() *HashMap {
	return &HashMap{data: make(map[interface{}]interface{})}
}

// Map Methods
func (m *HashMap) Put(key interface{}, value interface{}) {
	m.data[key] = value
}

func (m *HashMap) Get(key interface{}) (value interface{}, exists bool) {
	value, exists = m.data[key]
	return
}

func (m *HashMap) Has(key interface{}) bool {
	_, exists := m.data[key]
	return exists
}

func (m *HashMap) Delete(key interface{}) {
	delete(m.data, key)
}

func (m *HashMap) Keys() []interface{} {
	keys := make([]interface{}, 0, len(m.data))
	for key := range m.data {
		keys = append(keys, key)
	}
	return keys
}

// Container Methods
func (m *HashMap) IsEmpty() bool {
	return m.Size() == 0
}

func (m *HashMap) Size() int {
	return len(m.data)
}

func (m *HashMap) Clear() {
	for k := range m.data {
		delete(m.data, k)
	}
}

func (m *HashMap) ToString() string {
	str := "Hashmap\n"
	str += fmt.Sprintf("%v", m.data)
	return str
}
