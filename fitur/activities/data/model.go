package data

import (
	"todo/fitur/activities"

	"gorm.io/gorm"
)

type Activities struct {
	gorm.Model
	Gender string
	Phone  string
	Name   string
	Email  string `gorm:"unique;not null"`
}

// register
func FromEntities(dataCore activities.ActivitiesEntities) Activities { //fungsi yang mengambil data dari entities usercore dan merubah data ke user gorm(model.go)
	return Activities{
		Model: gorm.Model{ID: dataCore.ID, UpdatedAt: dataCore.Updatedat,
			CreatedAt: dataCore.Createdat,
		},
		Email:  dataCore.Email,
		Name:   dataCore.Name,
		Gender: dataCore.Gender,
		Phone:  dataCore.Phone,
	}

}

// profile user
func (dataModel *Activities) ModelsToCore() activities.ActivitiesEntities { //fungsi yang mengambil data dari  user gorm(model.go)  dan merubah data ke entities usercore
	return activities.ActivitiesEntities{
		ID:        dataModel.ID,
		Name:      dataModel.Name,
		Phone:     dataModel.Phone,
		Gender:    dataModel.Gender,
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

// func ToCore(model Activities) activities.ActivitiesEntities {
// 	return activities.ActivitiesEntities{
// 		ID:        model.ID,
// 		Title:     model.Title,
// 		Email:     model.Email,
// 		Updatedat: model.UpdatedAt,
// 	}

// }
