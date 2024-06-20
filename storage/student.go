package storage

import (
	"app/entity"
	"errors"
	"sync"
)

type StudentStorage struct {
	studentMap map[uint64]entity.Student
	next       uint64
	mutx       sync.RWMutex
}

func NewStudentStorage() *StudentStorage {
	return &StudentStorage{
		studentMap: make(map[uint64]entity.Student),
	}
}

var st *StudentStorage

func init() {
	st = new(StudentStorage)
}
func CreateStudent(student entity.Student) entity.Student {
	st.mutx.Lock()
	defer st.mutx.Unlock()
	student.Sid = st.next
	st.next++
	st.studentMap[student.Sid] = student
	return student
}

func UpdateStudent(student entity.Student) error {
	st.mutx.Lock()
	defer st.mutx.Unlock()
	if _, exists := st.studentMap[student.Sid]; !exists {
		return errors.New("Студент не найден")
	}
	st.studentMap[student.Sid] = student
	return nil
}

func DeleteStudent(iid uint64) error {
	st.mutx.Lock()
	defer st.mutx.Unlock()
	if _, exists := st.studentMap[iid]; !exists {
		return errors.New("Студент не найден")
	}
	delete(st.studentMap, iid)
	return nil
}

func GetAllStudents() ([]entity.Student, error) {
	st.mutx.RLock()
	defer st.mutx.RUnlock()
	studentList := make([]entity.Student, len(st.studentMap))
	var inc int
	for _, student := range st.studentMap {
		studentList[inc] = student
		inc++
	}
	return studentList, nil
}

func GetStudentByID(iid uint64) (*entity.Student, error) {
	st.mutx.RLock()
	defer st.mutx.RUnlock()
	student, exists := st.studentMap[iid]
	if !exists {
		return nil, errors.New("Студент не найден")
	}
	return &student, nil
}
