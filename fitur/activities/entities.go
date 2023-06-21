package activities

import (
	"time"
)

type ActivitiesEntities struct {
	ID        uint
	Name      string
	Email     string
	Phone     string
	Gender    string
	Createdat time.Time
	Updatedat time.Time
}

type ActivitiesService interface {
	FormData(newActivity ActivitiesEntities) (data ActivitiesEntities, row int, err error)
	GetActivity(name, gender string, pagination, limit int) (data []ActivitiesEntities, totalpage int, err error)
	GetId(id int) (data ActivitiesEntities, row int, err error)
	Updata(id int, datup ActivitiesEntities) (ActivitiesEntities, error)
	Delete(id int) error
}

type ActivitiesData interface {
	FormData(newActivity ActivitiesEntities) (data ActivitiesEntities, row int, err error)
	GetActivity(name, gender string, limit, offset int) (data []ActivitiesEntities, totalpage int, err error)
	UniqueData(insert ActivitiesEntities) (row int, err error)
	GetId(id int) (data ActivitiesEntities, row int, err error)
	Updata(id int, datup ActivitiesEntities) (ActivitiesEntities, error)
	Delete(id int) error
}
