package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dshns/todo-list/internal/utils"
)

type tasksHandler struct {
}

func NewTasksHandler() *tasksHandler {
	return &tasksHandler{}
}

func (handler *tasksHandler) NextDate(w http.ResponseWriter, r *http.Request) {
	now := r.FormValue("now")
	date := r.FormValue("date")
	repeat := r.FormValue("repeat")

	parsedTime, err := time.Parse(utils.DateFormat, now)
	if err != nil {
		fmt.Fprintf(w, `{"error": "%s"}`, err.Error())
		return
	}

	nextDate, err := utils.NextDate(parsedTime, date, repeat)
	if err != nil {
		fmt.Fprintf(w, `{"error": "%s"}`, err.Error())
		return
	}

	fmt.Fprintln(w, nextDate)
}
