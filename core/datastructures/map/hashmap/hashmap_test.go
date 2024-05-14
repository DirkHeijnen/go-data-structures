package advent_of_code

import "testing"

func TestMap_SetGet(t *testing.T) {
	m := NewMap()
	m.Put("key1", 10)
	m.Put("key2", 20)

	if val, exists := m.Get("key1"); !exists || val != 10 {
		t.Errorf("Expected key1 to be 10, got %d", val)
	}

	if val, exists := m.Get("key2"); !exists || val != 20 {
		t.Errorf("Expected key2 to be 20, got %d", val)
	}

	if _, exists := m.Get("key3"); exists {
		t.Error("Expected key3 to not exist")
	}
}

func TestMap_Delete(t *testing.T) {
	m := NewMap()
	m.Put("key1", 10)
	m.Delete("key1")

	if _, exists := m.Get("key1"); exists {
		t.Error("Expected key1 to be deleted")
	}
}

func TestMap_Size(t *testing.T) {
	m := NewMap()
	m.Put("key1", 10)
	m.Put("key2", 20)

	if size := m.Size(); size != 2 {
		t.Errorf("Expected size to be 2, got %d", size)
	}
}

func TestMap_HasKey(t *testing.T) {
	m := NewMap()
	m.Put("key1", 10)

	if !m.Has("key1") {
		t.Errorf("Expected key1 to exist")
	}

	if m.Has("key2") {
		t.Errorf("Did not expect key2 to exist")
	}
}

func BenchmarkMap_Set(b *testing.B) {
	m := NewMap()
	for i := 0; i < b.N; i++ {
		m.Put(i, i*10)
	}
}

func BenchmarkMap_Get(b *testing.B) {
	m := NewMap()
	for i := 0; i < b.N; i++ {
		m.Put(i, i*10)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Get(i)
	}
}

func BenchmarkMap_HasKey(b *testing.B) {
	m := NewMap()
	for i := 0; i < b.N; i++ {
		m.Put(i, i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Has(i)
	}
}
