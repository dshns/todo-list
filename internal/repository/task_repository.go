package repository

import (
	"database/sql"
	"fmt"

	"github.com/dshns/todo-list/internal/database"
	"github.com/dshns/todo-list/internal/models"
)

func NewTaskRepository(storage *database.AccessDatabase) *TaskRepository {
	return &TaskRepository{storage: storage}
}

type TaskRepository struct {
	storage *database.AccessDatabase
}

func (r *TaskRepository) AddTask(task *models.Task) (int, error) {
	res, err := r.storage.DB.Exec("INSERT INTO scheduler (date, title, comment, repeat) VALUES (:date, :title, :comment, :repeat)",
		sql.Named("date", task.Date),
		sql.Named("title", task.Title),
		sql.Named("comment", task.Comment),
		sql.Named("repeat", task.Repeat))

	if err != nil {
		return 0, fmt.Errorf("failed to add task: %w", err)
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("failed to get last id: %w", err)
	}

	return int(lastID), nil
}
