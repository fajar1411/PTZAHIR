package handler

import (
	"log"
	"net/http"
	"todo/fitur/todos"
	"todo/helper"

	"github.com/labstack/echo/v4"
)

type TodosHandler struct {
	TodoServices todos.TodoService
}

func (th *TodosHandler) AddTodo(c echo.Context) error {

	Inputform := TodoRequest{}
	Inputform.Priority = "very-high"
	errbind := c.Bind(&Inputform)
	if errbind != nil {
		return c.JSON(http.StatusBadRequest, "format inputan salah")
	}

	datacore := TodoRequestToEnitities(Inputform)
	res, err2 := th.TodoServices.AddTodo(datacore)
	log.Print("ini handler", res)

	if err2 != nil {
		return c.JSON(helper.PesanGagalHelper(err2.Error()))
	}
	dataResp := ToFormResponse(res)
	return c.JSON(http.StatusCreated, helper.Responsive{
		Status:  "Success",
		Massage: "Success",
		Data:    dataResp,
	})

}
