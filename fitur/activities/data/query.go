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
func (ad *activitiesData) FormData(newActivity activities.ActivitiesEntities) (data activities.ActivitiesEntities, row int, err error) {
	activitiesGorm := FromEntities(newActivity)
	tx := ad.db.Create(&activitiesGorm) // proses insert data

	if tx.Error != nil {
		log.Println("activities query error", tx.Error.Error())
		msg := ""
		if strings.Contains(tx.Error.Error(), "Duplicated") {
			msg = "email already exists"
		} else {
			msg = "server error"
		}
		return data, 0, errors.New(msg)
	}
	newActivity.ID = activitiesGorm.ID
	newActivity.Createdat = activitiesGorm.CreatedAt
	newActivity.Updatedat = activitiesGorm.UpdatedAt
	return newActivity, int(tx.RowsAffected), nil
}

// GetActivity implements activities.ActivitiesData
func (ac *activitiesData) GetActivity(name string, gender string, limit int, offset int) (data []activities.ActivitiesEntities, totalpage int, err error) {
	var activ []Activities

	var count int64
	tx1 := ac.db.Model(&activ).Where("name LIKE ? AND gender LIKE ? ", "%"+name+"%", "%"+gender+"%").Count(&count)
	if tx1.Error != nil {
		return nil, 0, tx1.Error
	}
	if tx1.RowsAffected == 0 {
		return nil, 0, errors.New("error query count")
	}
	if count < 10 {
		totalpage = 1
	} else if int(count)%limit == 0 {
		totalpage = int(count) / limit
	} else {
		totalpage = (int(count) / limit) + 1
	}
	tx := ac.db.Raw("SELECT activities.id, activities.name, activities.email, activities.gender, activities.phone,activities.created_at, activities.updated_at From activities WHERE activities.name= ? AND activities.gender= ? AND activities.deleted_at IS NULL", name, gender).Limit(limit).Offset(offset).Find(&activ)
	if tx.Error != nil {
		return nil, 0, tx.Error
	}
	var activcore = ListModelEntities(activ)
	return activcore, totalpage, nil
}

// // GetId implements activities.ActivitiesData
func (ac *activitiesData) GetId(id int) (data activities.ActivitiesEntities, row int, err error) {
	var activ Activities

	tx := ac.db.Raw("SELECT activities.id, activities.name, activities.email, activities.gender, activities.phone, activities.created_at, activities.updated_at From activities WHERE activities.id= ? AND activities.deleted_at IS NULL", id).Find(&activ)

	if tx.Error != nil {
		log.Println("All Activities error", tx.Error.Error())
		return data, 0, tx.Error
	}
	var activcore = activ.ModelsToCore()
	return activcore, int(tx.RowsAffected), nil
}

// // Updata implements activities.ActivitiesData
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
	tx := ad.db.Raw("SELECT activities.id, activities.name,activities.email, activities.gender, activities.phone, activities.created_at, activities.updated_at From activities Where activities.id= ? AND activities.deleted_at IS NULL", id).Find(&activ)

	if tx.Error != nil {
		log.Println("All Activities error", tx.Error.Error())
		return activities.ActivitiesEntities{}, tx.Error
	}
	var activcore = activ.ModelsToCore()
	return activcore, nil

}

// // Delete implements activities.ActivitiesData
func (ad *activitiesData) Delete(id int) error {
	activ := Activities{}
	qry := ad.db.Delete(&activ, id)

	rowAffect := qry.RowsAffected
	if rowAffect == 0 {
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

// // UniqueData implements activities.ActivitiesData
func (ad *activitiesData) UniqueData(insert activities.ActivitiesEntities) (row int, err error) {
	var datas Activities

	insertdata := FromEntities(insert)
	tx := ad.db.Raw("SELECT activities.id, activities.title, activities.email, activities.created_at, activities.updated_at From activities WHERE activities.email= ? AND activities.deleted_at IS NULL", insertdata.Email).Find(&datas)

	if tx.Error != nil {
		return 0, tx.Error
	}
	return int(tx.RowsAffected), nil
}
