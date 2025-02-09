package main

import (
	"log"
	"simple_api/internal/database"
	"simple_api/internal/handlers"
	"simple_api/internal/taskService"
	"simple_api/internal/userService"
	"simple_api/internal/web/tasks"
	"simple_api/internal/web/users"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	database.InitDB()
	if err := database.DB.AutoMigrate(&taskService.Task{}, &userService.User{}); err != nil {
		log.Fatalf("failed to start with err: %v", err)
	}

	tasksRepo := taskService.NewTaskRepository(database.DB)
	tasksService := taskService.NewService(tasksRepo)

	tasksHandler := handlers.NewTasksHandler(tasksService)

	usersRepo := userService.NewUserRepository(database.DB)
	usersService := userService.NewUserService(usersRepo)

	userHandler := handlers.NewUsersHandler(usersService)

	// Инициализируем echo
	e := echo.New()

	// используем Logger и Recover
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Прикол для работы в echo. Передаем и регистрируем хендлер в echo
	strictTasksHandler := tasks.NewStrictHandler(tasksHandler, nil) // тут будет ошибка
	tasks.RegisterHandlers(e, strictTasksHandler)

	strictUsersHandler := users.NewStrictHandler(userHandler, nil) // тут будет ошибка
	users.RegisterHandlers(e, strictUsersHandler)

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("failed to start with err: %v", err)
	}
}
