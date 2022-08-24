package models

import "github.com/google/uuid"

// AddTaskData структура запроса на добавление задачи
type AddTaskData struct {
	title, description string
}

// Title заголовок задачи
func (d *AddTaskData) Title() string {
	return d.title
}

// Description описание задачи
func (d *AddTaskData) Description() string {
	return d.description
}

// NewAddTaskData новый запрос на создание задачи
func NewAddTaskData(title, description string) *AddTaskData {
	return &AddTaskData{
		title:       title,
		description: description,
	}
}

// UpdateTaskData структура запроса на обновление задачи
type UpdateTaskData struct {
	id                 *uuid.UUID
	title, description string
}

// ID ID задачи
func (d *UpdateTaskData) ID() *uuid.UUID {
	return d.id
}

// Title заголовок задачи
func (d *UpdateTaskData) Title() string {
	return d.title
}

// Description описание задачи
func (d *UpdateTaskData) Description() string {
	return d.description
}

// NewUpdateTaskData создание запроса на обновление
func NewUpdateTaskData(ID *uuid.UUID, title string, description string) *UpdateTaskData {
	return &UpdateTaskData{
		id:          ID,
		title:       title,
		description: description,
	}
}

// ListTaskData структура запроса на получение списка задач
type ListTaskData struct {
	limit  uint64
	offset uint64
}

// Limit максимальное кол-во выдаваемых задач
func (d *ListTaskData) Limit() uint64 {
	return d.limit
}

// Offset сколько задач пропустить
func (d *ListTaskData) Offset() uint64 {
	return d.offset
}

// NewListTaskData новый запрос на чтение списка
func NewListTaskData(limit uint64, offset uint64) *ListTaskData {
	return &ListTaskData{
		limit:  limit,
		offset: offset,
	}
}

// DeleteTaskData структура запроса на удаления задачи
type DeleteTaskData struct {
	id *uuid.UUID
}

// ID ID задачи
func (d *DeleteTaskData) ID() *uuid.UUID {
	return d.id
}

// NewDeleteTaskData новый запрос на удаление
func NewDeleteTaskData(ID *uuid.UUID) *DeleteTaskData {
	return &DeleteTaskData{
		id: ID,
	}
}

// GetTaskData структура запроса на получение описания задачи
type GetTaskData struct {
	id *uuid.UUID
}

// ID ID задачи
func (d *GetTaskData) ID() *uuid.UUID {
	return d.id
}

// NewGetTaskData новый запрос на чтение задачи
func NewGetTaskData(ID *uuid.UUID) *GetTaskData {
	return &GetTaskData{
		id: ID,
	}
}
