package migrasi

import (
	activities "todo/fitur/activities/data"

	"gorm.io/gorm"
)

func MigrateDB(db *gorm.DB) {
	db.AutoMigrate(&activities.Activities{})
	// db.AutoMigrate(&owner.Owner{})

}
