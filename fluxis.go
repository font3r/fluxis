package main

import (
	"fmt"
	"strings"
)

type Storage struct {
	keys map[string]*Key
}

type Key struct {
	Key   string
	Value string
	TTL   int
}

func NewStorage() Storage {
	return Storage{keys: make(map[string]*Key)}
}

func (s *Storage) GetKey(key string) *Key {
	if key, exists := s.keys[key]; exists {
		return key
	}

	return &Key{}
}

func (s *Storage) SetKey(key string, value string) {
	s.keys[key] = &Key{Key: key, Value: value, TTL: 0}
}

func (s *Storage) Debug() string {
	data := strings.Builder{}
	for _, v := range s.keys {
		data.WriteString(fmt.Sprintf("%s = %s (%d)\n", v.Key, v.Value, v.TTL))
	}

	return data.String()
}
