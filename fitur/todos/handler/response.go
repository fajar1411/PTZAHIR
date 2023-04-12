package handler

import (
	"time"

	"todo/fitur/todos"
)

type FormResponse struct {
	ID           uint      `json:"id"`
	ActivitiesID uint      `json:"activity_group_id"`
	Title        string    `json:"title"`
	IsActive     bool      `json:"is_active"`
	Priority     string    `json:"priority"`
	Createdat    time.Time `json:"createdAt"`
	Updatedat    time.Time `json:"updatedAt"`
}

func ToFormResponse(data todos.TodoEntities) FormResponse {
	return FormResponse{
		ID:           data.ID,
		Title:        data.Title,
		Priority:     data.Priority,
		IsActive:     data.IsActive,
		Createdat:    data.Createdat,
		Updatedat:    data.Updatedat,
		ActivitiesID: data.ActivitiesID,
	}
}

func ListCoreToRespons(dataentitys []todos.TodoEntities) []FormResponse {
	var activres []FormResponse

	for _, val := range dataentitys {
		activres = append(activres, ToFormResponse(val))
	}
	return activres
}
