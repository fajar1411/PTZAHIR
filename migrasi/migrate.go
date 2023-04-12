package migrasi

import (
	activities "todo/fitur/activities/data"
	todos "todo/fitur/todos/data"

	"gorm.io/gorm"
)

func MigrateDB(db *gorm.DB) {
	db.AutoMigrate(&activities.Activities{})
	db.AutoMigrate(&todos.Todo{})

}
