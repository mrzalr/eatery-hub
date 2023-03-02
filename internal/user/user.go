package user

import "github.com/mrzalr/eatery-hub/internal/models"

type Repository interface {
	Create(user models.User) (models.User, error)
}
type Usecase interface {
	CreateUser(user models.User) (models.UserResponse, error)
}
