package storage

import (
	"errors"
)

type Employee struct {
	Id     int
	name   string
	age    string
	salary int
}

type Storage interface {
	Insert(e Employee) error
	Delete(id int) error
	Get(id int) (Employee, error)
}

type memoryStorage struct {
	data map[int]Employee
}

func NewMemoryStorage() *memoryStorage {
	return &memoryStorage{
		data: make(map[int]Employee),
	}
}

func (s *memoryStorage) Insert(e Employee) error {
	s.data[e.Id] = e

	return nil
}

func (s *memoryStorage) Get(id int) (Employee, error) {
	e, exists := s.data[id]
	if !exists {
		return Employee{}, errors.New("employee with such id doesn't exist")
	}

	return e, nil
}

func (s *memoryStorage) Delete(id int) error {
	delete(s.data, id)
	return nil
}
