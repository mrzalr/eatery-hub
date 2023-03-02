package http

import (
	"github.com/gin-gonic/gin"
	"github.com/mrzalr/eatery-hub/internal/models"
	"github.com/mrzalr/eatery-hub/internal/user"
	"github.com/mrzalr/eatery-hub/pkg/utils"
)

type handler struct {
	usecase user.Usecase
}

func New(usecase user.Usecase) *handler {
	return &handler{usecase}
}

func (h *handler) CreateUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user := models.User{}
		err := ctx.BindJSON(&user)
		if err != nil {
			errs := []string{err.Error()}
			models.ResponseBadRequest(ctx, errs)
			return
		}

		errs := user.Validate(utils.Validator)
		if len(errs) > 0 {
			models.ResponseBadRequest(ctx, errs)
			return
		}

		userResponse, err := h.usecase.CreateUser(user)
		if err != nil {
			errs := []string{err.Error()}
			models.ResponseBadGateway(ctx, errs)
			return
		}

		models.ResponseCreated(ctx, userResponse)
	}
}
