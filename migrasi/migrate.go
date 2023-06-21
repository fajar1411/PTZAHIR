package migrasi

import (
	_activities "todo/fitur/activities/data"

	"gorm.io/gorm"
)

func MigrateDB(db *gorm.DB) {
	db.AutoMigrate(&_activities.Activities{})
	// db.AutoMigrate(&_todos.Todos{})

}
