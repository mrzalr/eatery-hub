package usecase

import (
	"github.com/gin-gonic/gin"
	"github.com/mrzalr/eatery-hub/internal/models"
	"github.com/mrzalr/eatery-hub/internal/user"
	"github.com/mrzalr/eatery-hub/pkg/utils"
)

type usecase struct {
	repository user.Repository
}

func New(repository user.Repository) user.Usecase {
	return &usecase{repository}
}

func (u *usecase) CreateUser(user models.User) (models.UserWithToken, error) {
	user, err := u.repository.Create(user)
	if err != nil {
		return models.UserWithToken{}, err
	}

	token, err := utils.GenerateJwtToken(gin.H{
		"userID":    user.ID,
		"firstname": user.FirstName,
		"lastname":  user.LastName,
		"email":     user.Email,
	})

	if err != nil {
		return models.UserWithToken{}, err
	}

	userResponse := models.UserResponse{}
	userResponse.ParseFromUser(user)

	userWithToken := models.UserWithToken{
		UserResponse: userResponse,
		Token:        token,
	}

	return userWithToken, nil
}
