package servises

import (
	"fmt"
	"time"

	"github.com/dshns/todo-list/internal/models"
	"github.com/dshns/todo-list/internal/repository"
	"github.com/dshns/todo-list/internal/utils"
)

type TaskServise struct {
	repositoryInst *repository.TaskRepository
}

func NewTaskServise(repo *repository.TaskRepository) *TaskServise {
	return &TaskServise{repositoryInst: repo}
}

func isCorrect(task *models.Task) error {
	now := time.Now()

	if task.Title == "" {
		return fmt.Errorf("the title field should not be empty")
	}

	if task.Date == "" {
		task.Date = now.Format(utils.DateFormat)
	}

	date, err := time.Parse(utils.DateFormat, task.Date)
	if err != nil {
		return fmt.Errorf("invalid date format")
	}

	if date.Format(utils.DateFormat) < now.Format(utils.DateFormat) {
		if task.Repeat == "" {
			task.Date = now.Format(utils.DateFormat)
		} else {
			nextDate, err := utils.NextDate(now, task.Date, task.Repeat)
			if err != nil {
				return err
			}

			task.Date = nextDate
		}
	}

	return nil
}

func (s *TaskServise) AddTask(task *models.Task) (int, error) {
	if err := isCorrect(task); err != nil {
		return 0, fmt.Errorf("failed add task: %w", err)
	}

	return s.repositoryInst.AddTask(task)
}

func (s *TaskServise) GetAllTasks() ([]models.Task, error) {
	tasks, err := s.repositoryInst.GetAllTasks()
	if err != nil {
		return nil, fmt.Errorf("failed get all tasks: %w", err)
	}

	return tasks, nil
}

func (s *TaskServise) EditingTask(task *models.Task) error {
	if task.ID == "" {
		return fmt.Errorf("the id field should not be empty")
	}

	if err := isCorrect(task); err != nil {
		return fmt.Errorf("failed edit task: %w", err)
	}

	return s.repositoryInst.EditingTask(task)
}

func (s *TaskServise) GetTaskByID(id int) (*models.Task, error) {
	if id <= 0 {
		return nil, fmt.Errorf("the id field should be greater than 0")
	}

	return s.repositoryInst.GetTaskByID(id)
}
