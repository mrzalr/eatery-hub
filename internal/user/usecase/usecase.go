package usecase

import (
	"github.com/mrzalr/eatery-hub/internal/models"
	"github.com/mrzalr/eatery-hub/internal/user"
)

type usecase struct {
	repository user.Repository
}

func New(repository user.Repository) user.Usecase {
	return &usecase{repository}
}

func (u *usecase) CreateUser(user models.User) (models.UserResponse, error) {
	user, err := u.repository.Create(user)
	if err != nil {
		return models.UserResponse{}, err
	}

	userResponse := models.UserResponse{}
	userResponse.ParseFromUser(user)

	return userResponse, nil
}
