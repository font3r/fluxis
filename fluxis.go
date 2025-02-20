package main

import (
	"fmt"
	"strings"
)

type Storage struct {
	entries map[string]*CacheEntry
}

type CacheEntry struct {
	Key   string
	Value string
	TTL   int
}

func NewStorage() Storage {
	return Storage{entries: make(map[string]*CacheEntry)}
}

func (s *Storage) GetKey(key string) *CacheEntry {
	if key, exists := s.entries[key]; exists {
		return key
	}

	return &CacheEntry{}
}

func (s *Storage) SetKey(key string, value string) {
	s.entries[key] = &CacheEntry{Key: key, Value: value, TTL: 0}
}

func (s *Storage) DeleteKey(key string) {
	delete(s.entries, key)
}

func (s *Storage) Debug() string {
	data := strings.Builder{}
	for _, v := range s.entries {
		data.WriteString(fmt.Sprintf("%s = %s (%d)\n", v.Key, v.Value, v.TTL))
	}

	return data.String()
}
