package routes

import (
	"todo/fitur/activities"
	handleractivities "todo/fitur/activities/handler"
	"todo/fitur/todos"
	handlerTodos "todo/fitur/todos/handler"

	"github.com/labstack/echo/v4"
)

func NewHandlerActivities(Service activities.ActivitiesService, e *echo.Echo) {
	handlers := &handleractivities.ActivitiesHandler{
		ActivitiesServices: Service,
	}

	e.POST("/activities/form", handlers.FormData)
	// e.GET("/activities", handlers.GetActivity)
	// e.GET("/activities/:id", handlers.GetId)
	// e.PATCH("/activities/:id", handlers.Updata)
	// e.DELETE("/activities/:id", handlers.Delete)

}

func NewHandlerTodo(Service todos.TodoService, e *echo.Echo) {
	handlers := &handlerTodos.TodosHandler{
		TodoServices: Service,
	}

	e.POST("/todo-items", handlers.AddTodo)
	e.PATCH("/todo-items/:id", handlers.Update)
	e.GET("/todo-items/:id", handlers.GetAll)

}
