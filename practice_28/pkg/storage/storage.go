package storage

import (
	"fmt"
	"skillbox/practice_28/pkg/student"
)

type Storage struct {
	students map[string]*student.Student
}

func NewStorage() Storage {
	return Storage{
		students: make(map[string]*student.Student),
	}
}

func (st *Storage) Get() []string {
	result := make([]string, 0)
	var str string
	for _, s := range st.students {
		str = fmt.Sprintf("%v %v %v", s.Name, s.Age, s.Grade)
		result = append(result, str)
	}
	return result
}

func (st *Storage) Put(s *student.Student) {
	st.students[s.Name] = s
}
