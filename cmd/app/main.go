package main

import (
	"log"
	"main/project/internal/database"
	"main/project/internal/handlers"
	"main/project/internal/taskService"
	"main/project/internal/web/tasks"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	database.InitDB()
	if err := database.DB.AutoMigrate(&taskService.Task{}); err != nil {
		log.Fatalf("Failed to auto-migrate: %v", err)
	}

	repo := taskService.NewTaskRepository(database.DB)
	service := taskService.NewService(repo)

	handler := handlers.NewHandler(service)

	// Инициализируем echo
	e := echo.New()

	// используем Logger и Recover
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Прикол для работы в echo. Передаем и регистрируем хендлер в echo
	strictHandler := tasks.NewStrictHandler(handler, nil) // тут будет ошибка
	tasks.RegisterHandlers(e, strictHandler)

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("failed to start with err: %v", err)
	}
}
