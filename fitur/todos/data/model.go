package data

import (
	"todo/fitur/activities/data"
	"todo/fitur/todos"

	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Title        string
	Priority     string
	IsActive     bool
	ActivitiesID uint
	Activities   data.Activities
}

// type TodosEx struct {
// 	ID           uint
// 	Title        string
// 	Priority     string
// 	IsActive     bool
// 	ActivitiesID uint
// }

func Todata(data todos.TodoEntities) Todo {
	return Todo{
		Model: gorm.Model{ID: data.ID, UpdatedAt: data.Updatedat,
			CreatedAt: data.Createdat,
		},
		Title:        data.Title,
		Priority:     data.Priority,
		IsActive:     data.IsActive,
		ActivitiesID: data.ActivitiesID,
	}
}

func (data *Todo) ModelsToCore() todos.TodoEntities { //fungsi yang mengambil data dari  user gorm(model.go)  dan merubah data ke entities usercore
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

// func ToCore(data Owner) owner.OwnerEntities {
// 	return owner.OwnerEntities{
// 		ID:        data.ID,
// 		Nama_Toko: data.Nama_Toko,
// 		Status:    data.Status,
// 		Alamat:    data.Alamat,
// 		Ktp:       data.Ktp,
// 		UserID:    data.UserID,
// 	}
// }
// func ListModelTOCore(dataModel []OwnerUser) []owner.OwnerEntities { //fungsi yang mengambil data dari  user gorm(model.go)  dan merubah data ke entities usercore
// 	var dataCore []owner.OwnerEntities
// 	for _, value := range dataModel {
// 		dataCore = append(dataCore, value.ModelsToCore())
// 	}
// 	return dataCore //  untuk menampilkan data ke controller
// }
