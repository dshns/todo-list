package repository

import (
	"database/sql"
	"fmt"
	"strconv"

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

func (r *TaskRepository) GetAllTasks() ([]models.Task, error) {
	rows, err := r.storage.DB.Query("SELECT id, date, title, comment, repeat FROM scheduler ORDER BY date LIMIT 50")
	if err != nil {
		return nil, fmt.Errorf("failed to get all tasks: %w", err)
	}

	defer rows.Close()

	var res []models.Task
	for rows.Next() {
		task := models.Task{}
		err := rows.Scan(&task.ID, &task.Date, &task.Title, &task.Comment, &task.Repeat)
		if err != nil {
			return nil, fmt.Errorf("failed to scan task: %w", err)
		}
		res = append(res, task)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to iterate over rows: %w", err)
	}

	return res, nil
}

func (r *TaskRepository) EditingTask(task *models.Task) error {
	id, err := strconv.Atoi(task.ID)
	if err != nil {
		return fmt.Errorf("failed to convert id to int: %w", err)
	}
	res, err := r.storage.DB.Exec("UPDATE scheduler SET date = :date, title = :title, comment = :comment, repeat = :repeat WHERE id = :id ",
		sql.Named("date", task.Date),
		sql.Named("title", task.Title),
		sql.Named("comment", task.Comment),
		sql.Named("repeat", task.Repeat),
		sql.Named("id", id))
	if err != nil {
		return fmt.Errorf("failed to edit task: %w", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get affected rows: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no changes: %w", err)
	}

	return nil
}

func (r *TaskRepository) GetTaskByID(id int) (*models.Task, error) {
	row := r.storage.DB.QueryRow("SELECT * FROM scheduler WHERE id = :id", sql.Named("id", id))

	task := models.Task{}
	err := row.Scan(&task.ID, &task.Date, &task.Title, &task.Comment, &task.Repeat)
	if err != nil {
		return nil, fmt.Errorf("failed to get task by id: %w", err)
	}

	return &task, nil
}
