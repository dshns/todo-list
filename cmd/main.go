package main

import (
	"fmt"
	"log"
	"os"

	"github.com/dshns/todo-list/internal/database"
	"github.com/dshns/todo-list/internal/handlers"
	"github.com/dshns/todo-list/internal/repository"
	"github.com/dshns/todo-list/internal/servises"
	"github.com/gofiber/fiber/v2"
)

func main() {
	connecter, err := database.OpenOrCreate("scheduler.db")
	if err != nil {
		log.Printf("Failed to connect to database: %v", err)
		return
	}
	defer connecter.DB.Close()

	port := ":7540"

	if envPort, exists := os.LookupEnv("TODO_PORT"); exists {
		port = fmt.Sprintf(":%s", envPort)
	}

	webDir := "./web"
	repo := repository.NewTaskRepository(connecter)
	serv := servises.NewTaskServise(repo)
	h := handlers.NewTasksHandler(serv)

	app := fiber.New()

	app.Get("/api/nextdate", h.NextDate)
	app.Get("/api/task", h.GetTaskByID)
	app.Get("/api/tasks", h.GetAllTasks)
	app.Post("/api/task", h.AddTask)
	app.Put("/api/task", h.EditingTask)

	app.Static("/", webDir)

	log.Printf("Server started on http://localhost%v", port)
	err = app.Listen(port)
	if err != nil {
		log.Fatal(err)
	}
}
