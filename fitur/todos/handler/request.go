package handler

import (
	"todo/fitur/todos"
)

type TodoRequest struct {
	ID           uint
	Priority     string `json:"priority" form:"priority"`
	IsActive     bool   `json:"is_active" form:"is_active"`
	ActivitiesID uint   `json:"activity_group_id" form:"activity_group_id"`
	Title        string `json:"title" form:"title"`
	Status       string `json:"status" form:"status"`
}

func TodoRequestToEnitities(data TodoRequest) todos.TodoEntities {
	return todos.TodoEntities{
		ID:           data.ID,
		Priority:     data.Priority,
		IsActive:     data.IsActive,
		ActivitiesID: data.ActivitiesID,
		Title:        data.Title,
		Status:       data.Status,
	}
}
