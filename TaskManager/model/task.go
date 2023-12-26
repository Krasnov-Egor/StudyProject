package model

import "github.com/google/uuid"

// Структура для представления задачи
type Task struct {
	ID        uuid.UUID
	Desc      string
	Completed bool
}
