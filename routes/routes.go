package routes

import (
	"todo/fitur/activities"
	handleractivities "todo/fitur/activities/handler"

	"github.com/labstack/echo/v4"
)

func NewHandlerActivities(Service activities.ActivitiesService, e *echo.Echo) {
	handlers := &handleractivities.ActivitiesHandler{
		ActivitiesServices: Service,
	}

	e.POST("/activities/form", handlers.FormData)
	e.GET("/activities", handlers.GetActivity)
	e.GET("/activities/:id", handlers.GetId)
	e.PATCH("/activities/:id", handlers.Updata)
	e.DELETE("/activities/:id", handlers.Delete)

}
