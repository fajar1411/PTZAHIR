package data

import (
	"todo/fitur/todos"

	"gorm.io/gorm"
)

type Todos struct {
	gorm.Model
	ActivityGroupID uint
	Title           string
	IsActive        bool
	Priority        string
}

func Todata(data todos.TodoEntities) Todos {
	return Todos{
		Model: gorm.Model{ID: data.ID, UpdatedAt: data.Updatedat,
			CreatedAt: data.Createdat,
		},
		Title:           data.Title,
		Priority:        data.Priority,
		IsActive:        data.IsActive,
		ActivityGroupID: data.ActivitiesID,
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
		ActivitiesID: data.ActivityGroupID,
	}
}

func ToCore(data Todos) todos.TodoEntities {
	return todos.TodoEntities{
		ID:           data.ID,
		Title:        data.Title,
		Priority:     data.Priority,
		IsActive:     data.IsActive,
		Updatedat:    data.UpdatedAt,
		ActivitiesID: data.ActivityGroupID,
	}
}

func ListModelTOCore(dataModel []Todos) []todos.TodoEntities { //fungsi yang mengambil data dari  user gorm(model.go)  dan merubah data ke entities usercore
	var dataCore []todos.TodoEntities
	for _, value := range dataModel {
		dataCore = append(dataCore, value.ModelsToCore())
	}
	return dataCore //  untuk menampilkan data ke controller
}
