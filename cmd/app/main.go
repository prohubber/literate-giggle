package main

import (
	"log"
	"main/project/internal/database"
	"main/project/internal/handlers"
	"main/project/internal/taskService"
	"main/project/internal/userService"
	"main/project/internal/web/tasks"
	"main/project/internal/web/users"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Инициализация базы данных
	database.InitDB()
	if err := database.DB.AutoMigrate(&taskService.Task{}, &userService.User{}); err != nil {
		log.Fatalf("Failed to auto-migrate: %v", err)
	}

	// Настройка TaskService
	tasksRepo := taskService.NewTaskRepository(database.DB)
	tasksService := taskService.NewService(tasksRepo)
	taskHandlers := handlers.NewHandler(tasksService)

	// Настройка UserService
	userRepo := userService.NewRepository(database.DB)
	userService := userService.NewUserService(userRepo)
	userHandlers := handlers.NewUserHandler(userService)

	// Инициализируем Echo
	e := echo.New()

	// Используем Logger и Recover
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Регистрируем хендлеры для задач
	tasksStrictHandler := tasks.NewStrictHandler(taskHandlers, nil)
	tasks.RegisterHandlers(e, tasksStrictHandler)

	// Регистрируем хендлеры для пользователей
	usersStrictHandler := users.NewStrictHandler(userHandlers, nil)
	users.RegisterHandlers(e, usersStrictHandler)

	// Запуск сервера
	if err := e.Start(":8080"); err != nil {
		log.Fatalf("Failed to start with error: %v", err)
	}
}
