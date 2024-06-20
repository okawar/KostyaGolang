package main

import (
	"app/config"
	"app/entity"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"net/http"
)

const (
	dbName = "vds"
)

var db *gorm.DB

func initDB(dbName string) (*gorm.DB, error) {
	dsn := fmt.Sprintf("user=postgres password=agysaap38 dbname=%s", dbName)

	// Подключение к базе данных
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})
	if err != nil {
		return nil, err
	}

	// AutoMigrate для таблицы Item
	if err := db.AutoMigrate(&entity.Student{}); err != nil {
		return nil, err
	}

	// AutoMigrate для таблицы Vds
	if err := db.AutoMigrate(&entity.Vds{}); err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&entity.Token{}); err != nil {
		return nil, err
	}
	return db, nil
}

func main() {
	// Инициализация базы данных в функции main
	var err error
	db, err = initDB("vds")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		sqlDB, err := db.DB()
		if err != nil {
			fmt.Println(err)
		} else {
			sqlDB.Close()
		}
	}()

	// Вызов других функций, передавая подключение к базе данных
	WebServ()
}

/*
func begin(user entity.User, vds entity.Vds) {
	WebServ()
	fmt.Println("Пользователь: ", user.Uid)
	fmt.Println("Логин: ", user.Login)
	fmt.Println("Пароль: ", user.Pass)
	fmt.Println("Имя: ", user.Name)
	fmt.Println("Админ доступ: ", user.Access)
	fmt.Println("----------------------------------------------")
	fmt.Println("Вам доступны следующие VDS: ")
	fmt.Println("ID машины: ", vds.Vid)
	fmt.Println("Список пользователей: ", vds.Uid)
	fmt.Println("Имя VDS: ", vds.Name)
	fmt.Println("Доступная оперативная память: ", vds.Ram)
	fmt.Println("Доступное дисковое пространство: ", vds.Hdd)
	fmt.Println("Доступные ядра: ", vds.Core)
	fmt.Println("Статус машины: ", vds.Status)
}*/

func WebServ() {
	cfg := config.Get()

	m := http.NewServeMux()

	m.Handle("/", http.HandlerFunc(handle))

	srv := &http.Server{
		Addr:    ":" + cfg.Server.Port,
		Handler: m,
	}
	//http.HandleFunc("/auth", AuthHandler)
	srv.ListenAndServe()
}
