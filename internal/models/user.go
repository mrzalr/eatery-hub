package models

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/mrzalr/eatery-hub/pkg/utils"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID          uuid.UUID `json:"id"`
	FirstName   string    `json:"firstname" validate:"required" gorm:"type:varchar(255);not null"`
	LastName    string    `json:"lastname" validate:"required" gorm:"type:varchar(255);not null"`
	Email       string    `json:"email" validate:"required,email" gorm:"type:varchar(255);unique;not null"`
	Password    string    `json:"password" validate:"required,min=6" gorm:"type:varchar(255);not null"`
	PhoneNumber string    `json:"phone_number" validate:"required,e164" gorm:"type:varchar(255);unique;not null"`
	Address     string    `json:"address"`
	DOB         time.Time `json:"date_of_birth"`
	PhotoUrl    string    `json:"photo_url"`
	IsAdmin     string    `json:"is_admin"`
}

func (u *User) BeforeCreate(db *gorm.DB) error {
	u.ID = uuid.New()
	return u.HashPassword()
}

func (u *User) Validate(v *validator.Validate) []string {
	errs := []string{}

	err := v.Struct(u)
	if err != nil {
		errs = utils.TranslateErrors(err)
	}

	return errs
}

func (u *User) HashPassword() error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(hashed)
	return nil
}

func (u *User) CompareHashPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}

type UserResponse struct {
	ID          uuid.UUID `json:"id"`
	FirstName   string    `json:"firstname" validate:"required"`
	LastName    string    `json:"lastname" validate:"required"`
	Email       string    `json:"email" validate:"required"`
	PhoneNumber string    `json:"phone_number" validate:"required"`
	Address     string    `json:"address"`
	DOB         time.Time `json:"date_of_birth"`
	PhotoUrl    string    `json:"photo_url"`
	CreatedAt   time.Time `json:"created_at"`
}

func (u *UserResponse) ParseFromUser(user User) {
	u.ID = user.ID
	u.FirstName = user.FirstName
	u.LastName = user.LastName
	u.Email = user.Email
	u.PhoneNumber = user.PhoneNumber
	u.Address = user.Address
	u.DOB = user.DOB
	u.PhotoUrl = user.PhotoUrl
	u.CreatedAt = user.CreatedAt
}
