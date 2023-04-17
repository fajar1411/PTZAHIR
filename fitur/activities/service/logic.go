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
func (ac *activitiesCase) FormData(newActivity activities.ActivitiesEntities) (activities.ActivitiesEntities, error) {

	errEmail := ac.vld.Var(newActivity.Email, "required,email")
	if errEmail != nil {
		return activities.ActivitiesEntities{}, errors.New("invalid format email")
	}
	errtitle := ac.vld.Var(newActivity.Title, "required")
	if errtitle != nil {
		return activities.ActivitiesEntities{}, errors.New("invalid format email")
	}

	res, err := ac.qry.FormData(newActivity)
	if err != nil {
		msg2 := ""
		if strings.Contains(err.Error(), "duplicated") {
			msg2 = "email already exists"
		} else if strings.Contains(err.Error(), "empty") {
			msg2 = "title not allowed empty"
		} else {
			msg2 = "server error"
		}
		return activities.ActivitiesEntities{}, errors.New(msg2)
	}

	return res, nil
}

// GetActivity implements activities.ActivitiesService
func (ac *activitiesCase) GetActivity() ([]activities.ActivitiesEntities, error) {
	all, err := ac.qry.GetActivity()

	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "Activities not found"
		} else {
			msg = "internal server error"
		}
		return nil, errors.New(msg)
	}
	return all, nil
}

// GetId implements activities.ActivitiesService
func (ac *activitiesCase) GetId(id int) (activities.ActivitiesEntities, error) {
	if id <= 0 {
		log.Println("activities belum terdaftar")
	}
	res, err := ac.qry.GetId(id)
	if err != nil {
		log.Println(err)
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "id activities not found"
		} else {
			msg = "terdapat masalah pada server"
		}
		return activities.ActivitiesEntities{}, errors.New(msg)
	}
	return res, nil
}

// Updata implements activities.ActivitiesService
func (ac *activitiesCase) Updata(id int, datup activities.ActivitiesEntities) (activities.ActivitiesEntities, error) {
	if id <= 0 {
		log.Println("Activities Tidak Ada")
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
	title := datup.Title
	if title != "" {
		errTitle := ac.vld.Var(title, "required")
		if errTitle != nil {
			log.Println("validation error", errTitle)
			msg := validasi.ValidationErrorHandle(errTitle)
			return activities.ActivitiesEntities{}, errors.New(msg)
		}
	}
	res, err := ac.qry.Updata(id, datup)
	if err != nil {
		msg2 := ""
		if strings.Contains(err.Error(), "duplicated") {
			msg2 = "email sudah terdaftar"
		} else {
			msg2 = "server error"
		}
		return activities.ActivitiesEntities{}, errors.New(msg2)
	}

	return res, nil
}

// Delete implements activities.ActivitiesService
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
