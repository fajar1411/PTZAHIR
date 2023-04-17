package data

import (
	"errors"
	"log"
	"strings"
	"todo/fitur/activities"

	"gorm.io/gorm"
)

type activitiesData struct {
	db *gorm.DB
}

func NewActivities(db *gorm.DB) activities.ActivitiesData {
	return &activitiesData{
		db: db,
	}
}

// FormData implements activities.ActivitiesData
func (ad *activitiesData) FormData(newActivity activities.ActivitiesEntities) (activities.ActivitiesEntities, error) {
	activitiesGorm := FromEntities(newActivity)
	tx := ad.db.Create(&activitiesGorm) // proses insert data

	if tx.Error != nil {
		log.Println("register query error", tx.Error.Error())
		msg := ""
		if strings.Contains(tx.Error.Error(), "Duplicate") {
			msg = "data is duplicated"
		} else {
			msg = "server error"
		}
		return activities.ActivitiesEntities{}, errors.New(msg)
	}
	newActivity.ID = activitiesGorm.ID
	newActivity.Createdat = activitiesGorm.CreatedAt
	newActivity.Updatedat = activitiesGorm.UpdatedAt
	return newActivity, nil
}

// GetActivity implements activities.ActivitiesData
func (ac *activitiesData) GetActivity() ([]activities.ActivitiesEntities, error) {
	var activ []Activities

	tx := ac.db.Raw("SELECT activities.id, activities.title, activities.email, activities.created_at, activities.updated_at From activities WHERE activities.deleted_at IS NULL").Find(&activ)

	if tx.Error != nil {
		log.Println("All Activities error", tx.Error.Error())
		return []activities.ActivitiesEntities{}, tx.Error
	}
	var activcore = ListModelEntities(activ)
	return activcore, nil
}

// GetId implements activities.ActivitiesData
func (ac *activitiesData) GetId(id int) (activities.ActivitiesEntities, error) {
	var activ Activities

	tx := ac.db.Raw("SELECT activities.id, activities.title, activities.email, activities.created_at, activities.updated_at From activities WHERE activities.ID= ? AND activities.deleted_at IS NULL", id).Find(&activ)

	if tx.Error != nil {
		log.Println("All Activities error", tx.Error.Error())
		return activities.ActivitiesEntities{}, tx.Error
	}
	var activcore = activ.ModelsToCore()
	return activcore, nil
}

// Updata implements activities.ActivitiesData
func (ad *activitiesData) Updata(id int, datup activities.ActivitiesEntities) (activities.ActivitiesEntities, error) {
	activ := Activities{}

	acgorm := FromEntities(datup)
	qry := ad.db.Model(&activ).Where("id = ?", id).Updates(&acgorm)

	affrows := qry.RowsAffected
	if affrows == 0 {
		log.Println("no rows affected")
		return activities.ActivitiesEntities{}, errors.New("no data updated")
	}
	err := qry.Error
	if err != nil {
		log.Println("update activities query error", err.Error())
		return activities.ActivitiesEntities{}, err
	}
	tx := ad.db.Raw("SELECT activities.id, activities.title, activities.email, activities.created_at, activities.updated_at From activities Where activities.id= ?", id).Find(&activ)

	if tx.Error != nil {
		log.Println("All Activities error", tx.Error.Error())
		return activities.ActivitiesEntities{}, tx.Error
	}
	var activcore = activ.ModelsToCore()
	return activcore, nil

}

// Delete implements activities.ActivitiesData
func (ad *activitiesData) Delete(id int) error {
	activ := Activities{}
	qry := ad.db.Delete(&activ, id)

	rowAffect := qry.RowsAffected
	if rowAffect <= 0 {
		log.Println("no data processed")
		return errors.New("no user has delete")
	}

	err := qry.Error
	if err != nil {
		log.Println("delete query error", err.Error())
		return errors.New("delete account fail")
	}

	return nil

}
