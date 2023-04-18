package handler

import (
	"net/http"
	"strconv"
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

	if err2 != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("fail add data"))
	}
	dataResp := ToFormResponse(res)
	return c.JSON(http.StatusCreated, helper.Responsive{
		Status:  "Success",
		Massage: "Success",
		Data:    dataResp,
	})

}
func (th *TodosHandler) Update(c echo.Context) error {

	Inputform := TodoRequest{}
	errbind := c.Bind(&Inputform)
	if errbind != nil {
		return c.JSON(http.StatusBadRequest, "format inputan salah")
	}
	todoID, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error Convert"))
	}
	datacore := TodoRequestToEnitities(Inputform)
	res, err2 := th.TodoServices.Update(todoID, datacore)

	if err2 != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("fail update data"))
	}
	dataResp := ToFormResponse(res)
	return c.JSON(http.StatusOK, helper.Responsive{
		Status:  "Success",
		Massage: "Success",
		Data:    dataResp,
	})

}

func (ad *TodosHandler) GetAll(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	res, err := ad.TodoServices.GetAll(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse(err.Error()))
	}
	dataResp := ListCoreToRespons(res)
	return c.JSON(http.StatusOK, helper.Responsive{
		Status:  "Success",
		Massage: "Success",
		Data:    dataResp,
	})

}
func (th *TodosHandler) DeleteData(c echo.Context) error {
	id, ercnv := strconv.Atoi(c.Param("id"))
	strIdTodo := strconv.Itoa(id)
	if ercnv != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponsFail{
			Status:  "Error",
			Massage: "Invalid ID",
		})
	}
	row, err := th.TodoServices.DeleteData(id)
	if row == 0 {
		return c.JSON(http.StatusNotFound, helper.Responsive{
			Status:  "Not Found",
			Massage: "Todo with ID " + strIdTodo + " Not Found",
			Data:    map[string]interface{}{},
		})

	}
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponsFail{
			Status:  "Error",
			Massage: "failed to get data",
		})
	}
	return c.JSON(http.StatusOK, helper.Responsive{
		Status:  "Success",
		Massage: "Success",
		Data:    map[string]interface{}{},
	})
}
