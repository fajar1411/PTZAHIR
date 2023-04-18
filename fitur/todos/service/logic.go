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
			msg = "content todo not found"
		} else {
			msg = "internal server error"
		}
		return todos.TodoEntities{}, errors.New(msg)
	}

	return res, nil
}

// Update implements todos.TodoService
func (tc *todoCase) Update(id int, input todos.TodoEntities) (todos.TodoEntities, error) {
	if id <= 0 {
		log.Println("Activities Tidak Ada")
	}

	errTitle := tc.vld.Var(input.Title, "required")
	if errTitle != nil {
		return todos.TodoEntities{}, errors.New("dont empty")
	}
	errpriority := tc.vld.Var(input.Priority, "required")
	if errpriority != nil {
		return todos.TodoEntities{}, errors.New("dont empty")
	}
	erractive := tc.vld.Var(input.IsActive, "required")
	if erractive != nil {
		return todos.TodoEntities{}, errors.New("dont empty")
	}
	errstatus := tc.vld.Var(input.Status, "required")
	if errstatus != nil {
		return todos.TodoEntities{}, errors.New("dont empty")
	}
	res, err := tc.qry.Update(id, input)
	if err != nil {
		// fmt.Println(err)
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "id not found"
		} else {
			msg = "internal server error"
		}
		return todos.TodoEntities{}, errors.New(msg)
	}

	return res, nil
}

// GetAll implements todos.TodoService
func (tc *todoCase) GetAll(activid int) ([]todos.TodoEntities, error) {
	all, err := tc.qry.GetAll(activid)

	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "Activities not found"
		} else {
			msg = "internal server error"
		}
		return nil, errors.New(msg)
	}
	return all, nil
}
