package storage

import (
	"app/entity"
	"testing"
)

func TestUserStorage(t *testing.T) {
	// Тестирование добавления пользователя
	student := entity.Student{Name: "John Doe"} // Удалите поле Age, если оно не определено в entity.Student
	createdStudent := CreateStudent(student)
	if createdStudent.Sid == 0 {
		t.Errorf("Failed to create user: student ID is zero")
	}

	// Тестирование получения пользователя по ID
	retrievedUser, err := GetStudentByID(createdStudent.Sid)
	if err != nil || retrievedUser.Sid != createdStudent.Sid {
		t.Errorf("Failed to retrieve student: %v", err)
	}

	// Тестирование обновления пользователя
	updatedUser := *retrievedUser
	updatedUser.Name = "Jane Doe"
	err = UpdateStudent(updatedUser)
	if err != nil {
		t.Errorf("Failed to update user: %v", err)
	}

	// Проверка обновленных данных
	updatedUserCheck, _ := GetStudentByID(updatedUser.Sid)
	if updatedUserCheck.Name != "Jane Doe" {
		t.Errorf("User update failed: expected name %s, got %s", "Jane Doe", updatedUserCheck.Name)
	}

	// Тестирование удаления пользователя
	err = DeleteStudent(createdStudent.Sid)
	if err != nil {
		t.Errorf("Failed to delete user: %v", err)
	}

	// Проверка наличия удаленного пользователя
	_, err = GetStudentByID(createdStudent.Sid)
	if err == nil {
		t.Errorf("User was not deleted")
	}

	// Тестирование получения всех пользователей
	users, err := GetAllStudents()
	if err != nil || len(users) != 0 {
		t.Errorf("Failed to get all student")
	}
}
