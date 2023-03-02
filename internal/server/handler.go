package server

import (
	userHttp "github.com/mrzalr/eatery-hub/internal/user/delivery/http"
	userRepo "github.com/mrzalr/eatery-hub/internal/user/repository/mysql"
	userUcase "github.com/mrzalr/eatery-hub/internal/user/usecase"
)

func (s *server) MapHandlers() {
	userRepository := userRepo.New(s.DB)

	userUsecase := userUcase.New(userRepository)

	userHandler := userHttp.New(userUsecase)

	v1 := s.App.Group("/api/v1")

	user := v1.Group("/users")
	userHandler.MapRoutes(user)
}
