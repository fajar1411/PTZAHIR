package handler

import (
	"time"
	"todo/fitur/todos"
)

type FormResponse struct {
	ID           uint      `json:"id"`
	Title        string    `json:"title"`
	Priority     string    `json:"priority"`
	IsActive     bool      `json:"is_active"`
	Updatedat    time.Time `json:"updatedAt"`
	Createdat    time.Time `json:"createdAt"`
	ActivitiesID uint      `json:"activity_group_id"`
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

// func ListCoreToRespons(dataentitys []activities.ActivitiesEntities) []FormResponse {
// 	var activres []FormResponse

// 	for _, val := range dataentitys {
// 		activres = append(activres, ToFormResponse(val))
// 	}
// 	return activres
// }
