package server

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type server struct {
	DB  *gorm.DB
	App *gin.Engine
}

func New(db *gorm.DB) *server {
	return &server{
		DB:  db,
		App: gin.New(),
	}
}

func (s *server) Run() error {
	port := fmt.Sprintf(":%s", os.Getenv("APP_PORT"))
	return s.App.Run(port)
}
