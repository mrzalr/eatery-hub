package mysql

import (
	"github.com/mrzalr/eatery-hub/internal/models"
	"github.com/mrzalr/eatery-hub/internal/user"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) user.Repository {
	return &repository{db}
}

func (r *repository) Create(user models.User) (models.User, error) {
	tx := r.db.Create(&user)
	return user, tx.Error
}
