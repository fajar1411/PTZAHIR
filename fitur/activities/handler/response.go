package handler

import (
	"time"
	"todo/fitur/activities"
)

type FormResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"Name"`
	Email     string    `json:"email"`
	Gender    string    `json:"gender"`
	Phone     string    `json:"phone"`
	Createdat time.Time `json:"created_At"`
	Updatedat time.Time `json:"updated_At"`
}

func ToFormResponse(data activities.ActivitiesEntities) FormResponse {
	return FormResponse{
		ID:        data.ID,
		Name:      data.Name,
		Email:     data.Email,
		Phone:     data.Phone,
		Gender:    data.Gender,
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
