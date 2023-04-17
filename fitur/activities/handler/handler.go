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
		return c.JSON(http.StatusBadRequest, helper.ResponsFail{
			Status:  "Error",
			Massage: "Failed to bind data, Check your input",
		})
	}

	dataCore := ActivitiesRequestToUserCore(Inputform)

	res, err := ad.ActivitiesServices.FormData(dataCore)
	if err != nil {
		return c.JSON(http.StatusNotFound, helper.ResponsFail{
			Status:  "Error",
			Massage: "Data failed to save",
		})
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
	id, _ := strconv.Atoi(c.Param("id"))

	res, err := ad.ActivitiesServices.GetId(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse(err.Error()))
	}
	dataResp := ToFormResponse(res)
	return c.JSON(http.StatusOK, helper.Responsive{
		Status:  "Success",
		Massage: "Success",
		Data:    dataResp,
	})

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
	reid := strconv.Itoa(id)
	res, err := ad.ActivitiesServices.Updata(id, ActivitiesRequestToUserCore(Inputform))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponsFail{
			Status:  "Error",
			Massage: err.Error(),
		})
	}
	if res.ID == 0 {
		return c.JSON(http.StatusNotFound, helper.ResponsFail{
			Status:  "Error",
			Massage: fmt.Sprintf(" Activity with ID " + reid + " Not Found "),
		})
	}
	dataResp := ToFormResponse(res)
	return c.JSON(http.StatusOK, helper.Responsive{
		Status:  "Success",
		Massage: "Success",
		Data:    dataResp,
	})

}

func (ad *ActivitiesHandler) Delete(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	err := ad.ActivitiesServices.Delete(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "internal server error",
		})
	}
	resid := strconv.Itoa(id)

	return c.JSON(http.StatusNotFound, helper.ResponsFail{
		Status:  "Not Found",
		Massage: fmt.Sprintf(" Activity with ID " + resid + " Not Found "),
	})

}
