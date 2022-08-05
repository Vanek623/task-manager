package models

import "time"

// Task краткая информация о задаче
type Task struct {
	ID    uint
	Title string
}

// DetailedTask подробная информаци о задаче
type DetailedTask struct {
	Title       string
	Description string
	Edited      time.Time
}
