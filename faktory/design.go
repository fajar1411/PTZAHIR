package faktory

import (
	activitiesdata "todo/fitur/activities/data"
	activitiesservice "todo/fitur/activities/service"

	tododata "todo/fitur/todos/data"
	todoservice "todo/fitur/todos/service"
	activitieshandler "todo/routes"

	todohandler "todo/routes"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitFactory(e *echo.Echo, db *gorm.DB) {
	v := validator.New()

	activitiesRepofaktory := activitiesdata.NewActivities(db)
	activitiesFaktory := activitiesservice.NewService(activitiesRepofaktory, v)
	activitieshandler.NewHandlerActivities(activitiesFaktory, e)

	Tododatafaktory := tododata.NewTodo(db)
	todoserviceFaktory := todoservice.NewService(Tododatafaktory, v)
	todohandler.NewHandlerTodo(todoserviceFaktory, e)
}
