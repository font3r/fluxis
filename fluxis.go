package main

import (
	"fmt"
	"strings"
	"time"
)

type Storage struct {
	entries map[string]*CacheEntry
}

type CacheEntry struct {
	Key   string
	Value string
	TTL   int64
}

func NewStorage() Storage {
	return Storage{entries: make(map[string]*CacheEntry)}
}

func (s *Storage) GetKey(key string, now func() time.Time) *CacheEntry {
	if entry, exists := s.entries[key]; exists {
		if !now().After(time.Unix(entry.TTL, 0)) {
			return entry
		}
	}

	return &CacheEntry{}
}

func (s *Storage) SetKey(key string, value string, ttl int, now func() time.Time) {
	s.entries[key] = &CacheEntry{
		Key:   key,
		Value: value,
		TTL:   now().Add(time.Second * time.Duration(ttl)).Unix(),
	}
}

func (s *Storage) DeleteKey(key string) {
	delete(s.entries, key)
}

func (s *Storage) Debug() string {
	data := strings.Builder{}
	for _, v := range s.entries {
		data.WriteString(fmt.Sprintf("%s = %s (%v)\n", v.Key, v.Value, v.TTL))
	}

	return data.String()
}
