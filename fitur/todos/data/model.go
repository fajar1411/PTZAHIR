package data

import (
	"todo/fitur/activities/data"
	"todo/fitur/todos"

	"gorm.io/gorm"
)

type Todos struct {
	gorm.Model
	Title        string
	Priority     string
	IsActive     bool
	ActivitiesID uint
	Activities   data.Activities
}

func Todata(data todos.TodoEntities) Todos {
	return Todos{
		Model: gorm.Model{ID: data.ID, UpdatedAt: data.Updatedat,
			CreatedAt: data.Createdat,
		},
		Title:        data.Title,
		Priority:     data.Priority,
		IsActive:     data.IsActive,
		ActivitiesID: data.ActivitiesID,
	}
}

func (data *Todos) ModelsToCore() todos.TodoEntities { //fungsi yang mengambil data dari  user gorm(model.go)  dan merubah data ke entities usercore
	return todos.TodoEntities{
		ID:           data.ID,
		Title:        data.Title,
		Priority:     data.Priority,
		IsActive:     data.IsActive,
		Createdat:    data.CreatedAt,
		Updatedat:    data.UpdatedAt,
		ActivitiesID: data.ActivitiesID,
	}
}

func ToCore(data Todos) todos.TodoEntities {
	return todos.TodoEntities{
		ID:           data.ID,
		Title:        data.Title,
		Priority:     data.Priority,
		IsActive:     data.IsActive,
		Updatedat:    data.UpdatedAt,
		ActivitiesID: data.ActivitiesID,
	}
}

func ListModelTOCore(dataModel []Todos) []todos.TodoEntities { //fungsi yang mengambil data dari  user gorm(model.go)  dan merubah data ke entities usercore
	var dataCore []todos.TodoEntities
	for _, value := range dataModel {
		dataCore = append(dataCore, value.ModelsToCore())
	}
	return dataCore //  untuk menampilkan data ke controller
}
