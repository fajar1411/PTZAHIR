package handler

import (
	"fmt"
	"log"
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
			Status:  "Error",
			Massage: "Failed to bind data, Check your input",
			Data:    map[string]interface{}{},
		})
	}

	dataCore := ActivitiesRequestToUserCore(Inputform)

	res, row, err := ad.ActivitiesServices.FormData(dataCore)
	if row == -1 {
		return c.JSON(http.StatusBadRequest, helper.ResponsFail{
			Status:  "Bad Request",
			Massage: err.Error(),
		})
	}

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponsFail{
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
	pageStr := c.QueryParam("page")

	page, _ := strconv.Atoi(pageStr)
	nama := c.QueryParam("name")
	gender := c.QueryParam("gender")
	limit := 3

	log.Println("nama:", nama)
	log.Println("geder:", gender)

	res, totalpage, err := ad.ActivitiesServices.GetActivity(nama, gender, page, limit)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse(err.Error()))
	}

	dataResp := ListCoreToRespons(res)

	return c.JSON(http.StatusOK, helper.ResponsivePage{
		Status:  "Success",
		Massage: "Success",
		Data:    dataResp,
		Page:    totalpage,
	})

}

func (ad *ActivitiesHandler) GetId(c echo.Context) error {
	id, errcnv := strconv.Atoi(c.Param("id"))
	if errcnv != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponsFail{
			Status:  "Error",
			Massage: "Invalid ID",
		})
	}

	res, row, err := ad.ActivitiesServices.GetId(id)
	reid := strconv.Itoa(id)
	if row == 0 {
		return c.JSON(http.StatusBadRequest, helper.Responsive{
			Status:  "Error",
			Massage: "Activity with ID " + reid + " Not Found",
			Data:    map[string]interface{}{},
		})
	}
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed to get data"))
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
