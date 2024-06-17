package forms

import (
	"database/sql"
	"errors"
	"log"
	"net/mail"
	"strings"

	"models"
)

type UserForm struct {
	FullName     string `json:"full_name"`
	Email        string `json:"email"`
	PhoneNumber  string `json:"phone_number"`
	BirthDate    string `json:"date_of_birth"`
	Address      string `json:"address"`
	NIK          string `json:"nik"`
	Gender       string `json:"gender"`
	DepartmentId int    `json:"department_id"`
	Photo        string `json:"photo"`
	NewPassword  string `json:"new_password"`
}

var (
	ErrInvalidEmail       = errors.New("invalid email address")
	ErrEmailExist         = errors.New("email address exists")
	ErrInvalidPhoneNumber = errors.New("invalid phone number")
	ErrInvalidNIK         = errors.New("invalid NIK")
)

func (user_form *UserForm) Sanitize() {
	user_form.FullName = strings.TrimSpace(user_form.FullName)
	user_form.Email = strings.TrimSpace(user_form.Email)
	user_form.PhoneNumber = strings.TrimSpace(user_form.PhoneNumber)
	user_form.BirthDate = strings.TrimSpace(user_form.BirthDate)
	user_form.Address = strings.TrimSpace(user_form.Address)
	user_form.NIK = strings.TrimSpace(user_form.NIK)
}

func (user_form UserForm) ValidateCreate() (bool, error) {
	// Email validator
	_, err := mail.ParseAddress(user_form.Email)

	// Ensure no error parsing email address
	// but, parsing error potentialy caused by
	// invalid email address
	if err != nil {
		return false, ErrEmailExist
	}

	// Validate uniqueness
	_, err =
		models.User{}.GetUsingEmail(user_form.Email)

	// Ensure no error get user using email
	if err != nil {
		// Ignore no row found error, because
		// the objective is finding an existing email
		if err != sql.ErrNoRows {
			log.Panic("Error get user", err.Error())

			return false, err
		}
	} else {

		return false, ErrEmailExist
	}

	// Phone number validator
	if len(user_form.PhoneNumber) < 11 || len(user_form.PhoneNumber) > 13 {
		return false, ErrInvalidPhoneNumber
	}

	// NIK validator
	if len(user_form.NIK) != 16 {
		return false, ErrInvalidNIK
	}

	return true, nil
}

func (user_form UserForm) ValidateUpdate(user_id int) (bool, error) {
    // Email validator
	_, err := mail.ParseAddress(user_form.Email)

	// Ensure no error parsing email address
	// but, parsing error potentialy caused by
	// invalid email address
	if err != nil {
		return false, ErrEmailExist
	}

	// Validate uniqueness
	mail_user, err := models.User{}.GetUsingEmail(user_form.Email)

	// Ensure no error get user using email
	if err != nil {
		// Ignore no row found error, because
		// the objective is finding an existing email
		if err != sql.ErrNoRows {
			log.Panic("Error get user", err.Error())

			return false, err
		}
	}

	if user_id != mail_user.Id {
		log.Print("Email exist")
		return false, ErrEmailExist
	}

    // Phone number validator
    if len(user_form.PhoneNumber) < 11 || len(user_form.PhoneNumber) > 13 {
        return false, ErrInvalidPhoneNumber
    }

    // NIK validator
    if len(user_form.NIK) != 16 {
        return false, ErrInvalidNIK
    }

    return true, nil
}

