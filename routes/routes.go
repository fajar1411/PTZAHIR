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

	e.POST("/activity-groups", handlers.FormData)
	e.GET("/activity-groups", handlers.GetActivity)
	e.GET("/activity-groups/:id", handlers.GetId)
	e.PATCH("/activity-groups/:id", handlers.Updata)
	e.DELETE("/activity-groups/:id", handlers.Delete)

}
