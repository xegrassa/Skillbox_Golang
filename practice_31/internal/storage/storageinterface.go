package storage

import (
	m "skillbox/practice_30/internal/models"
)

type Storage interface {
	AddNewUser(*m.User) int
	GetUser(int) (*m.User, bool)
	DeleteUser(int)
	UpdateUserAge(int, int) error
	ToString() string
}
