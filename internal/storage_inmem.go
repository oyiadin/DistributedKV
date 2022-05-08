package internal

import "errors"

type InMemoryStorage map[string]string

func (s *InMemoryStorage) Init() {

}

func (s InMemoryStorage) Get(key string) (value string, err error) {
	value, ok := s[key]
	if ok {
		err = nil
	} else {
		err = errors.New("invalid key")
	}
	return
}

func (s *InMemoryStorage) Set(key, value string) (err error) {
	(*s)[key] = value
	err = nil
	return
}

func NewInMemoryStorage() *InMemoryStorage {
	return &InMemoryStorage{}
}
