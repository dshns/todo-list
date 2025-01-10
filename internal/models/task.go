package models

type Task struct {
	ID      string `json:"id" binding:"required"`
	Date    string `json:"date"`
	Title   string `json:"title" binding:"required"`
	Comment string `json:"comment,omitempty"`
	Repeat  string `json:"repeat"`
}
