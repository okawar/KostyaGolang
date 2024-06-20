package service

import (
	"app/engine"
	"app/entity"
	"app/storage"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func (e *Service) CreateStudent(ctx *engine.Context) error { //POST
	decoder := json.NewDecoder(ctx.Request.Body)
	var student entity.Student
	err := decoder.Decode(&student)
	if err != nil {
		return err
	}
	student = storage.CreateStudent(student)
	coder, _ := json.Marshal(student)
	ctx.Response.Write(coder)
	return nil

}

func (e *Service) UpdateStudent(ctx *engine.Context) error { //PUT
	decoder := json.NewDecoder(ctx.Request.Body)
	var student entity.Student
	err := decoder.Decode(&student)
	if err != nil {
		return err
	}

	if err := storage.UpdateStudent(student); err != nil {
		return err
	}

	return nil
}

func (e *Service) DeleteStudent(ctx *engine.Context) error { //DELETE
	path := strings.Split(ctx.Request.URL.Path, "/")

	sid, err := strconv.ParseUint(path[len(path)-1], 10, 64)
	if err != nil {
		return err
	}

	err = storage.DeleteStudent(sid)
	if err != nil {
		return err
	}

	ctx.Response.WriteHeader(http.StatusOK)
	return nil
}

func (e *Service) GetStudent(ctx *engine.Context) error { //GET
	studentList, err := storage.GetAllStudents()
	if err != nil {
		return err
	}

	coder, _ := json.Marshal(studentList)
	ctx.Response.Write(coder)

	return nil

}

func (e *Service) GetByStudent(ctx *engine.Context) error { //GET
	sidStr := ctx.Request.URL.Query().Get("sid")
	fmt.Println(sidStr)
	if sidStr == "" {
		return errors.New("Параметр 'sid' не указан")
	}

	sid, err := strconv.ParseUint(sidStr, 10, 64)
	if err != nil {
		return err
	}
	fmt.Println(sid)
	student, err := storage.GetStudentByID(sid)
	if err != nil {
		return err
	}

	coder, err := json.Marshal(student)
	if err != nil {
		return err
	}

	ctx.Response.Write(coder)
	return nil
}
