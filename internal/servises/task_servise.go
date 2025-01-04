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
		return 0, err
	}

	return s.repositoryInst.AddTask(task)
}
