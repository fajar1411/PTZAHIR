package activities

import (
	"time"
)

type ActivitiesEntities struct {
	ID        uint
	Title     string `validate:"required"`
	Email     string `validate:"required,email"`
	Createdat time.Time
	Updatedat time.Time
}

type ActivitiesService interface {
	FormData(newActivity ActivitiesEntities) (ActivitiesEntities, error)
	GetActivity() ([]ActivitiesEntities, error)
	GetId(id int) (ActivitiesEntities, error)
	Updata(id int, datup ActivitiesEntities) (ActivitiesEntities, error)
	Delete(id int) error
}

type ActivitiesData interface {
	FormData(newActivity ActivitiesEntities) (ActivitiesEntities, error)
	GetActivity() ([]ActivitiesEntities, error)
	GetId(id int) (ActivitiesEntities, error)
	Updata(id int, datup ActivitiesEntities) (ActivitiesEntities, error)
	Delete(id int) error
}
