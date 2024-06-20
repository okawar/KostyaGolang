package main

import (
	"app/entity"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestAuthHandler(t *testing.T) {
	// Настройка подключения к тестовой базе данных PostgreSQL
	dsn := fmt.Sprintf("user=postgres password=postgre dbname=%s", dbName)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}

	// Миграция схемы таблицы пользователей
	db.AutoMigrate(&entity.Student{})

	// Создание тестового пользователя
	user := entity.Student{
		Login:  "testuser",
		Pass:   "password123",
		Name:   "John Doe",
		Sid:    1,
		Access: true,
	}
	db.Create(&user)

	// Создание запроса для тестирования AuthHandler
	testAuth := SAuth{
		Login: "testuser",
		Pass:  "password123",
	}
	requestBody, _ := json.Marshal(testAuth)
	request := httptest.NewRequest(http.MethodPost, "/auth", bytes.NewBuffer(requestBody))
	request.Header.Set("Content-Type", "application/json")

	// Создание записи ответа
	responseRecorder := httptest.NewRecorder()

	// Вызов функции AuthHandler
	AuthHandler(responseRecorder, request)

	// Проверка ожидаемого кода состояния
	if status := responseRecorder.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Проверка тела ответа
	var response struct {
		Token  string `json:"token"`
		Sid    uint64 `json:"sid"`
		Name   string `json:"name"`
		Access bool   `json:"access"`
	}
	json.NewDecoder(responseRecorder.Body).Decode(&response)

	if response.Sid != user.Sid || response.Name != user.Name || response.Access != user.Access {
		t.Errorf("Handler returned unexpected body: got %v want %v", response, user)
	}
}
