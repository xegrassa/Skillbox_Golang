package storage

import (
	"fmt"
	h "skillbox/practice_30/internal/helpers"
	m "skillbox/practice_30/internal/models"

	"strings"
)

func NewMemStorage() *MemStorage {
	return &MemStorage{
		users:      make(map[int]*m.User),
		lastUserId: 0,
	}
}

type MemStorage struct {
	users      map[int]*m.User
	lastUserId int
}

func (s *MemStorage) AddNewUser(u *m.User) int {
	s.users[s.lastUserId] = u
	s.lastUserId++
	return s.lastUserId - 1
}

func (s *MemStorage) GetUser(id int) (*m.User, bool) {
	u, ok := s.users[id]
	return u, ok
}

func (s *MemStorage) DeleteUser(id int) {
	delete(s.users, id)

	for _, u := range s.users {
		index := h.GetElemIndexFromSlice(u.Friends, id)
		if index != -1 {
			u.Friends = h.DeleteElemFromSlice(u.Friends, index)
		}
	}
}

func (s *MemStorage) UpdateUserAge(id int, newAge int) error {
	u, ok := s.users[id]
	if !ok {
		return fmt.Errorf("Не найден пользователь с id:%v", id)
	}
	u.Age = newAge
	return nil
}

func (s *MemStorage) ToString() string {
	result := []string{}
	for k, u := range s.users {
		str := fmt.Sprintf("id: %v  name: %v  age: %d  friend: %v", k, u.Name, u.Age, u.Friends)
		result = append(result, str)
	}
	return strings.Join(result, "\n")
}
