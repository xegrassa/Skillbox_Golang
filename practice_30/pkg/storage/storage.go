package storage

import (
	"fmt"
	h "skillbox/practice_30/pkg/helpers"
	m "skillbox/practice_30/pkg/models"

	"strings"
)

var UserStorage = newStorage()

type Storage struct {
	users      map[int]*m.User
	lastUserId int
}

func (s *Storage) AddNewUser(u *m.User) int {
	s.users[s.lastUserId] = u
	s.lastUserId++
	return s.lastUserId - 1
}

func (s *Storage) GetUser(id int) (*m.User, bool) {
	u, ok := s.users[id]
	return u, ok
}

func (s *Storage) DeleteUser(id int) {
	delete(s.users, id)

	for _, u := range s.users {
		index := h.GetElemIndexFromSlice(u.Friends, id)
		if index != -1 {
			u.Friends = h.DeleteElemFromSlice(u.Friends, index)
		}
	}
}

func (s *Storage) ToString() string {
	result := []string{}
	for k, u := range s.users {
		str := fmt.Sprintf("id: %v  name: %v  age: %d  friend: %v", k, u.Name, u.Age, u.Friends)
		result = append(result, str)
	}
	return strings.Join(result, "\n")
}

func newStorage() Storage {
	return Storage{
		users:      make(map[int]*m.User),
		lastUserId: 0,
	}
}
