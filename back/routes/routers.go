package routes

import (
	"ped_poject/database"
	"ped_poject/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(
	r *gin.Engine,
	authHandler *handlers.AuthHandler,
) {

	r.POST("/login", authHandler.Login)
	r.POST("/register", authHandler.Register)
}
func SetupRouter(
	userRepo *database.UserRepository,
) *gin.Engine {

	r := gin.Default()

	authHandler := handlers.NewAuthHandler(userRepo)

	SetupRoutes(
		r,
		authHandler,
	)

	return r
}
