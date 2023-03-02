package http

import "github.com/gin-gonic/gin"

func (h *handler) MapRoutes(app *gin.RouterGroup) {
	app.POST("/register", h.CreateUser())
}
