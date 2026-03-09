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
	r.POST("/check_email", authHandler.CheckEmail)
	r.POST("/create_task", taskHandler.CreateTask)
	r.POST("/search_task", taskHandler.SearchTask)
	r.PUT("/update_task", taskHandler.Update_Task)
	r.PUT("/update_user", authHandler.UpdateProfile)
}
func SetupRouter(
	userRepo *database.UserRepository,
	taskRepo *database.TaskRepository,
) *gin.Engine {

	r := gin.Default()

	authHandler := handlers.NewAuthHandler(userRepo)
	taskHandler := handlers.NewTaskHandler(taskRepo, userRepo)
	SetupRoutes(
		r,
		authHandler,
		taskHandler,
	)

	return r
}
