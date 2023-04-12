package service

import (
	"errors"
	"log"
	"strings"
	"todo/fitur/todos"
	"todo/validasi"

	"github.com/go-playground/validator/v10"
)

type todoCase struct {
	qry todos.TodoData
	vld *validator.Validate
}

func NewService(td todos.TodoData, vld *validator.Validate) todos.TodoService {
	return &todoCase{
		qry: td,
		vld: vld,
	}
}

// AddTodo implements todos.TodoService
func (tc *todoCase) AddTodo(newTodo todos.TodoEntities) (todos.TodoEntities, error) {
	valerr := tc.vld.Struct(&newTodo)
	if valerr != nil {
		log.Println("validation error", valerr)
		msgvalid := validasi.ValidationErrorHandle(valerr)
		return todos.TodoEntities{}, errors.New(msgvalid)
	}
	res, err := tc.qry.AddTodo(newTodo)
	if err != nil {
		// fmt.Println(err)
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "content family not found"
		} else {
			msg = "internal server error"
		}
		return todos.TodoEntities{}, errors.New(msg)
	}

	return res, nil
}
