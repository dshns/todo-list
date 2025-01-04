package handlers

import (
	"time"

	"github.com/dshns/todo-list/internal/models"
	"github.com/dshns/todo-list/internal/servises"
	"github.com/dshns/todo-list/internal/utils"
	"github.com/gofiber/fiber/v2"
)

func NewTasksHandler(serv *servises.TaskServise) *tasksHandler {
	return &tasksHandler{serviceInst: serv}
}

type tasksHandler struct {
	serviceInst *servises.TaskServise
}

func (handler *tasksHandler) NextDate(c *fiber.Ctx) error {
	now := c.Query("now")
	date := c.Query("date")
	repeat := c.Query("repeat")

	if now == "" || date == "" || repeat == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Missing required query parameters"})
	}

	parsedTime, err := time.Parse(utils.DateFormat, now)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid 'now' format: " + err.Error()})
	}

	nextDate, err := utils.NextDate(parsedTime, date, repeat)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendString(nextDate)
}

func (handler *tasksHandler) AddTask(c *fiber.Ctx) error {
	var task models.Task
	if err := c.BodyParser(&task); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	id, err := handler.serviceInst.AddTask(&task)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"id": id})
}
