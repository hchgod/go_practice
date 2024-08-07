package main

import (
	"fmt"
	"sync"
)

type SMap struct {
	mu   sync.Mutex
	data map[string]string
}

func NewSMap() *SMap {
	return &SMap{
		data: make(map[string]string),
	}
}

func (m *SMap) Get(k string) string {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.data[k]
}

func (m *SMap) Set(k, v string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.data[k] = v
}

func (m *SMap) Delete(k string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.data, k)
}

func main() {
	// Create a new SMap instance
	m := NewSMap()

	// Set key-value pairs
	m.Set("key1", "value1")
	m.Set("key2", "value2")
	m.Set("key3", "value3")

	// Retrieve values
	fmt.Println("Value for key1:", m.Get("key1"))
	fmt.Println("Value for key2:", m.Get("key2"))
	fmt.Println("Value for key3:", m.Get("key3"))

	// Delete a key
	m.Delete("key2")

	// Check if a key exists
	if value := m.Get("key2"); value == "" {
		fmt.Println("Key 'key2' not found")
	} else {
		fmt.Println("Value for key2 after deletion:", value)
	}
}
