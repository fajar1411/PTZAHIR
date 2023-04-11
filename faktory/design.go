package faktory

import (
	activitiesdata "todo/fitur/activities/data"
	activitiesservice "todo/fitur/activities/service"

	activitieshandler "todo/routes"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitFactory(e *echo.Echo, db *gorm.DB) {
	v := validator.New()

	activitiesRepofaktory := activitiesdata.NewActivities(db)
	activitiesFaktory := activitiesservice.NewService(activitiesRepofaktory, v)
	activitieshandler.NewHandlerActivities(activitiesFaktory, e)

}
