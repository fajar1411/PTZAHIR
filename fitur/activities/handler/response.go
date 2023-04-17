package handler

import (
	"time"
	"todo/fitur/activities"
)

type FormResponse struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Email     string    `json:"email"`
	Createdat time.Time `json:"created_At"`
	Updatedat time.Time `json:"updated_At"`
}

func ToFormResponse(data activities.ActivitiesEntities) FormResponse {
	return FormResponse{
		ID:        data.ID,
		Title:     data.Title,
		Email:     data.Email,
		Createdat: data.Createdat,
		Updatedat: data.Updatedat,
	}
}
func ListCoreToRespons(dataentitys []activities.ActivitiesEntities) []FormResponse {
	var activres []FormResponse

	for _, val := range dataentitys {
		activres = append(activres, ToFormResponse(val))
	}
	return activres
}
