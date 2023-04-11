package handler

import (
	"todo/fitur/activities"
)

type ActivitiesRequest struct {
	ID    uint
	Title string `json:"title" form:"title"`
	Email string `json:"email" form:"email"`
}

func ActivitiesRequestToUserCore(data ActivitiesRequest) activities.ActivitiesEntities {
	return activities.ActivitiesEntities{
		ID:    data.ID,
		Title: data.Title,
		Email: data.Email,
	}
}
