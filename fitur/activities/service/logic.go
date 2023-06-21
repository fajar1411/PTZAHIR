package service

import (
	"errors"
	"log"
	"strings"
	"todo/fitur/activities"
	"todo/validasi"

	"github.com/go-playground/validator/v10"
)

type activitiesCase struct {
	qry activities.ActivitiesData
	vld *validator.Validate
}

func NewService(ad activities.ActivitiesData, vld *validator.Validate) activities.ActivitiesService {
	return &activitiesCase{
		qry: ad,
		vld: vld,
	}
}

// FormData implements activities.ActivitiesService
func (ac *activitiesCase) FormData(newActivity activities.ActivitiesEntities) (data activities.ActivitiesEntities, row int, err error) {

	errEmail := ac.vld.Var(newActivity.Email, "required,email")
	if errEmail != nil {
		return data, -1, errors.New("invalid format email")
	}
	errtitle := ac.vld.Var(newActivity.Name, "required")
	if errtitle != nil {
		return data, -1, errors.New("title cannot be null")
	}
	rowUnique, _ := ac.qry.UniqueData(newActivity)
	if rowUnique == 1 {
		return data, -1, errors.New("email already exists")
	}
	res, row, err := ac.qry.FormData(newActivity)

	return res, row, err
}

// GetActivity implements activities.ActivitiesService
func (ac *activitiesCase) GetActivity(name string, gender string, pagination int, limit int) (data []activities.ActivitiesEntities, totalpage int, err error) {
	offset := (pagination - 1) * limit
	data, totalpage, err = ac.qry.GetActivity(name, gender, limit, offset)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "Activities not found"
		} else {
			msg = "internal server error"
		}
		return nil, 0, errors.New(msg)
	}
	return data, totalpage, nil
}

// // GetId implements activities.ActivitiesService
func (ac *activitiesCase) GetId(id int) (data activities.ActivitiesEntities, row int, err error) {
	if id <= 0 {
		log.Println("activities belum terdaftar")
	}
	res, row, err := ac.qry.GetId(id)

	return res, row, err
}

// // Updata implements activities.ActivitiesService
func (ac *activitiesCase) Updata(id int, datup activities.ActivitiesEntities) (activities.ActivitiesEntities, error) {
	if id <= 0 {
		log.Println("activities belum terdaftar")
	}

	email := datup.Email
	if email != "" {
		errEmail := ac.vld.Var(email, "required,email")
		if errEmail != nil {
			log.Println("validation error", errEmail)
			msg := validasi.ValidationErrorHandle(errEmail)
			return activities.ActivitiesEntities{}, errors.New(msg)
		}
	}
	Name := datup.Name
	if Name != "" {
		errName := ac.vld.Var(Name, "required")
		if errName != nil {
			log.Println("validation error", errName)
			msg := validasi.ValidationErrorHandle(errName)
			return activities.ActivitiesEntities{}, errors.New(msg)
		}
	}
	Phone := datup.Phone
	if Phone != "" {
		errphone := ac.vld.Var(Phone, "required")
		if errphone != nil {
			log.Println("validation error", errphone)
			msg := validasi.ValidationErrorHandle(errphone)
			return activities.ActivitiesEntities{}, errors.New(msg)
		}
	}
	res, err := ac.qry.Updata(id, datup)
	if err != nil {
		msg2 := ""
		if strings.Contains(err.Error(), "Duplicate") {
			msg2 = "email sudah terdaftar"
		} else if strings.Contains(err.Error(), "not found") {
			msg2 = "id activities not found"
		} else {
			msg2 = "server error"
		}
		return activities.ActivitiesEntities{}, errors.New(msg2)
	}

	return res, nil
}

// // Delete implements activities.ActivitiesService
func (ac *activitiesCase) Delete(id int) error {
	if id <= 0 {
		log.Println("Activites not found")
	}
	err := ac.qry.Delete(id)

	if err != nil {
		log.Println("query error", err.Error())
		return errors.New("query error, delete account fail")
	}

	return nil
}
