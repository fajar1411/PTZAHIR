package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"todo/fitur/activities"
	"todo/helper"

	"github.com/labstack/echo/v4"
)

type ActivitiesHandler struct {
	ActivitiesServices activities.ActivitiesService
}

func (ad *ActivitiesHandler) FormData(c echo.Context) error {

	Inputform := ActivitiesRequest{}

	errbind := c.Bind(&Inputform)
	if errbind != nil {
		return c.JSON(http.StatusBadRequest, helper.Responsive{
			Status:  http.StatusText(echo.ErrBadRequest.Code),
			Massage: errbind.Error(),
			Data:    map[string]interface{}{},
		})
	}

	if Inputform.Email == "" {
		return c.JSON(http.StatusBadRequest, helper.Responsive{
			Status:  http.StatusText(echo.ErrBadRequest.Code),
			Massage: "Check email input",
			Data:    map[string]interface{}{},
		})
	} else if Inputform.Title == "" {
		return c.JSON(http.StatusBadRequest, helper.Responsive{
			Status:  http.StatusText(echo.ErrBadRequest.Code),
			Massage: "Check title input",
			Data:    map[string]interface{}{},
		})
	}
	dataCore := ActivitiesRequestToUserCore(Inputform)

	res, err := ad.ActivitiesServices.FormData(dataCore)

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse(err.Error()))
	}

	dataResp := ToFormResponse(res)
	return c.JSON(http.StatusCreated, helper.Responsive{
		Status:  "Success",
		Massage: "Success",
		Data:    dataResp,
	})

}
func (ad *ActivitiesHandler) GetActivity(c echo.Context) error {

	res, err := ad.ActivitiesServices.GetActivity()
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

func (ad *ActivitiesHandler) GetId(c echo.Context) error {
	id, errcnv := strconv.Atoi(c.Param("id"))
	if errcnv != nil {
		return c.JSON(http.StatusBadRequest, helper.Responsive{
			Status:  "Error",
			Massage: "Invalid ID",
			Data:    map[string]interface{}{},
		})
	}

	res, err := ad.ActivitiesServices.GetId(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, helper.FailedResponse(err.Error()))
	}
	dataResp := ToFormResponse(res)
	reid := strconv.Itoa(id)
	if dataResp.ID == 0 {
		return c.JSON(http.StatusNotFound, helper.Responsive{
			Status:  "Error",
			Massage: fmt.Sprintf(" Activity with ID " + reid + " Not Found "),
			Data:    map[string]interface{}{},
		})
	} else {
		return c.JSON(http.StatusOK, helper.Responsive{
			Status:  "Success",
			Massage: "Success",
			Data:    dataResp,
		})
	}

}

func (ad *ActivitiesHandler) Updata(c echo.Context) error {
	id, errcnv := strconv.Atoi(c.Param("id"))

	if errcnv != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponsFail{
			Status:  "Error",
			Massage: "Invalid ID",
		})
	}
	Inputform := ActivitiesRequest{}
	if err := c.Bind(&Inputform); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponsFail{
			Status:  "Error",
			Massage: "Failed to bind data, Check your input",
		})
	}

	res, err := ad.ActivitiesServices.Updata(id, ActivitiesRequestToUserCore(Inputform))
	if err != nil {
		return c.JSON(http.StatusNotFound, helper.FailedResponse(err.Error()))
	}
	dataResp := ToFormResponse(res)
	reid := strconv.Itoa(id)
	if dataResp.ID == 0 {
		return c.JSON(http.StatusNotFound, helper.ResponsFail{
			Status:  "Error",
			Massage: fmt.Sprintf(" Activity with ID " + reid + " Not Found "),
		})
	} else {
		return c.JSON(http.StatusOK, helper.Responsive{
			Status:  "Success",
			Massage: "Success",
			Data:    dataResp,
		})
	}

}

func (ad *ActivitiesHandler) Delete(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	err := ad.ActivitiesServices.Delete(id)
	reid := strconv.Itoa(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.Responsive{
			Status:  "Error",
			Massage: fmt.Sprintf(" Activity with ID " + reid + " Not Found "),
			Data:    map[string]interface{}{},
		})
	}
	resid := strconv.Itoa(id)

	return c.JSON(http.StatusNotFound, helper.ResponsFail{
		Status:  "Not Found",
		Massage: fmt.Sprintf(" Activity with ID " + resid + " Not Found "),
	})

}
