package handler

import (
	"todo/fitur/activities"
)

type ActivitiesRequest struct {
	ID     uint
	Name   string `json:"name" form:"name"`
	Phone  string `json:"phone" form:"phone"`
	Email  string `json:"email" form:"email"`
	Gender string `json:"gender" form:"gender"`
}

func ActivitiesRequestToUserCore(data ActivitiesRequest) activities.ActivitiesEntities {
	return activities.ActivitiesEntities{
		ID:     data.ID,
		Name:   data.Name,
		Email:  data.Email,
		Phone:  data.Phone,
		Gender: data.Gender,
	}
}
