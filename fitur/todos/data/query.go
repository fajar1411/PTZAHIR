package data

import (
	"errors"
	"log"
	"strings"
	"todo/fitur/todos"

	"gorm.io/gorm"
)

type todoData struct {
	db *gorm.DB
}

func NewTodo(db *gorm.DB) todos.TodoData {
	return &todoData{
		db: db,
	}
}

// / AddTodo implements todos.TodoData
func (td *todoData) AddTodo(newTodo todos.TodoEntities) (todos.TodoEntities, error) {

	data := Todata(newTodo)

	err := td.db.Create(&data)
	if err.Error != nil {
		log.Println("add Todo query error", err.Error.Error())
		msg := ""
		if strings.Contains(err.Error.Error(), "not valid") {
			msg = "wrong input"

		} else {
			msg = "server error"
		}
		return todos.TodoEntities{}, errors.New(msg)
	}
	newTodo.ID = data.ID
	newTodo.Createdat = data.CreatedAt
	newTodo.Updatedat = data.UpdatedAt
	return newTodo, nil
}

// Update implements todos.TodoData
func (td *todoData) Update(id int, input todos.TodoEntities) (todos.TodoEntities, error) {
	todo := Todos{}
	data := Todata(input)
	tx := td.db.Model(&todo).Where("id = ?", id).Updates(&data)

	if tx.Error != nil {
		log.Println("update todo error :", tx.Error)
		return todos.TodoEntities{}, tx.Error

	}
	if tx.RowsAffected <= 0 {
		log.Println("update todo query error : data not found")
		return todos.TodoEntities{}, errors.New("not found")
	}
	tx2 := td.db.Raw("SELECT todos.id, todos.title, todos.priority, todos.is_active, todos.created_at, todos.updated_at, todos.activities_id From todos Where todos.id= ?", id).Find(&todo)
	if tx2.Error != nil {
		log.Println("All Activities error", tx.Error.Error())
		return todos.TodoEntities{}, tx2.Error
	}
	var todocore = todo.ModelsToCore()
	return todocore, nil
}

// GetAll implements todos.TodoData
func (td *todoData) GetAll(activid int) ([]todos.TodoEntities, error) {
	var todo []Todos

	tx := td.db.Raw("SELECT todos.id, todos.title, todos.priority, todos.is_active, todos.created_at, todos.updated_at, todos.activities_id From todos WHERE todos.activities_id= ?", activid).Find(&todo)

	if tx.Error != nil {
		log.Println("All Activities error", tx.Error.Error())
		return []todos.TodoEntities{}, tx.Error
	}
	var activcore = ListModelTOCore(todo)
	return activcore, nil
}
