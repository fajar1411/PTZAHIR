package todos

import "time"

type TodoEntities struct {
	ID           uint
	Title        string `validate:"required,min=3,required"`
	Priority     string
	IsActive     bool
	Createdat    time.Time
	Updatedat    time.Time
	ActivitiesID uint
}

type TodoService interface {
	AddTodo(newTodo TodoEntities) (TodoEntities, error)
}

type TodoData interface {
	AddTodo(newTodo TodoEntities) (TodoEntities, error)
}
