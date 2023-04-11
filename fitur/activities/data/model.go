package data

import (
	"todo/fitur/activities"

	"gorm.io/gorm"
)

type Activities struct {
	gorm.Model
	Title string `gorm:"type:char(50);not null"`
	Email string `gorm:"type:varchar(50);unique;not null"`
}

// register
func FromEntities(dataCore activities.ActivitiesEntities) Activities { //fungsi yang mengambil data dari entities usercore dan merubah data ke user gorm(model.go)
	return Activities{
		Model: gorm.Model{ID: dataCore.ID, UpdatedAt: dataCore.Updatedat,
			CreatedAt: dataCore.Createdat,
		},
		Email: dataCore.Email,
		Title: dataCore.Title,
	}

}

// profile user
func (dataModel *Activities) ModelsToCore() activities.ActivitiesEntities { //fungsi yang mengambil data dari  user gorm(model.go)  dan merubah data ke entities usercore
	return activities.ActivitiesEntities{
		ID:        dataModel.ID,
		Title:     dataModel.Title,
		Email:     dataModel.Email,
		Updatedat: dataModel.UpdatedAt,
		Createdat: dataModel.CreatedAt,
	}
}

func ListModelEntities(datamodel []Activities) []activities.ActivitiesEntities {
	var entities []activities.ActivitiesEntities

	for _, val := range datamodel {
		entities = append(entities, val.ModelsToCore())
	}
	return entities

}

func ToCore(model Activities) activities.ActivitiesEntities {
	return activities.ActivitiesEntities{
		ID:        model.ID,
		Title:     model.Title,
		Email:     model.Email,
		Updatedat: model.UpdatedAt,
	}

}
