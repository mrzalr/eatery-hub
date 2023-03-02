package models

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ResponseBadRequest(ctx *gin.Context, errs []string) {
	ctx.JSON(http.StatusBadRequest,
		Response{
			Status:  http.StatusBadRequest,
			Message: "bad request",
			Errors:  errs,
			Data:    nil,
		},
	)
}

func ResponseBadGateway(ctx *gin.Context, errs []string) {
	ctx.JSON(http.StatusBadRequest,
		Response{
			Status:  http.StatusBadGateway,
			Message: "bad gateway",
			Errors:  errs,
			Data:    nil,
		},
	)
}

func ResponseCreated(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusBadRequest,
		Response{
			Status:  http.StatusCreated,
			Message: "created",
			Errors:  []string{},
			Data:    data,
		},
	)
}
