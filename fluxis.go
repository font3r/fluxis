package main

import "fmt"

type Storage struct {
	keys map[string]string
}

func NewStorage() Storage {
	return Storage{keys: make(map[string]string)}
}

func (s *Storage) GetKey(key string) string {
	if key, exists := s.keys[key]; exists {
		return key
	}

	return ""
}

func (s *Storage) SetKey(key string, value string) error {
	s.keys[key] = value

	return nil
}

func (s *Storage) Debug() string { return fmt.Sprintf("%v", s.keys) }
