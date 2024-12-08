package main

import (
	"main/internal/database"
	"main/internal/handlers"
	"main/internal/taskService"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Инициализация базы данных
	database.InitDB()
	database.DB.AutoMigrate(&taskService.Task{})

	// Создание репозитория, сервиса и хендлеров
	repo := taskService.NewTaskRepository(database.DB)
	service := taskService.NewService(repo)
	handler := handlers.NewHandler(service)

	// Инициализация роутера и маршрутов
	router := mux.NewRouter()
	router.HandleFunc("/api/get", handler.GetTasksHandler).Methods("GET")
	router.HandleFunc("/api/post", handler.PostTaskHandler).Methods("POST")
	router.HandleFunc("/api/patch/{id}", handler.UpdateTaskHandler).Methods("PATCH")
	router.HandleFunc("/api/delete/{id}", handler.DeleteTaskHandler).Methods("DELETE")

	// Запуск сервера
	http.ListenAndServe(":8080", router)
}
