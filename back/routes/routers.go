package routes

import (
	"ped_poject/database"
	"ped_poject/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(
	r *gin.Engine,
	authHandler *handlers.AuthHandler,
	taskHandler *handlers.TaskHandler,
) {

	r.POST("/login", authHandler.Login)
	r.POST("/register", authHandler.Register)
	r.POST("/create_task", taskHandler.CreateTask)
	r.POST("/search_task", taskHandler.SearchTask)
}
func SetupRouter(
	userRepo *database.UserRepository,
	taskRepo *database.TaskRepository,
) *gin.Engine {

	r := gin.Default()

	authHandler := handlers.NewAuthHandler(userRepo)
	taskHandler := handlers.NewTaskHandler(taskRepo)
	SetupRoutes(
		r,
		authHandler,
		taskHandler,
	)

	return r
}
