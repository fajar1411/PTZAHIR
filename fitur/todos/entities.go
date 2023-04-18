package todos

import "time"

type TodoEntities struct {
	ID           uint
	Title        string
	Priority     string
	IsActive     bool
	Status       string
	Createdat    time.Time
	Updatedat    time.Time
	ActivitiesID uint
}

type TodoService interface {
	AddTodo(newTodo TodoEntities) (TodoEntities, error)
	Update(id int, input TodoEntities) (TodoEntities, error)
	GetAll(activid int) ([]TodoEntities, error)
	DeleteData(id int) (row int, err error)
}

type TodoData interface {
	AddTodo(newTodo TodoEntities) (TodoEntities, error)
	Update(id int, input TodoEntities) (TodoEntities, error)
	GetAll(activid int) ([]TodoEntities, error)
	DeleteData(id int) (row int, err error)
}
