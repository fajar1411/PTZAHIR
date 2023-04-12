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
