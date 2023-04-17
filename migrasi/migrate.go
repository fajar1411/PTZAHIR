package migrasi

import (
	_activities "todo/fitur/activities/data"
	_todos "todo/fitur/todos/data"

	"gorm.io/gorm"
)

func MigrateDB(db *gorm.DB) {
	db.AutoMigrate(&_activities.Activities{})
	db.AutoMigrate(&_todos.Todos{})

}
