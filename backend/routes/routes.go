package routes

import (
	"backend/controllers"
	"backend/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, todoController *controllers.TodoController, authController *controllers.AuthController) {

	r.LoadHTMLGlob("templates/*")

	r.GET("/index", todoController.HomePage)

	api := r.Group("/api")
	{
		api.POST("/register", authController.Register)
		api.POST("/login", authController.Login)
		api.GET("/users", authController.GetUsers)

		// Protected routes
		protected := api.Group("/")
		protected.Use(middleware.AuthMiddleware())
		{

			// Todo routes
			protected.POST("/todos", todoController.CreateTodo)
			protected.GET("/todos", todoController.GetTodos)
			protected.PUT("/todos/:id", todoController.UpdateTodo)
			protected.DELETE("/todos/:id", todoController.DeleteTodo)
		}

		//api.GET("/users", userController.GetUsers)
		//api.GET("/users/:id", userController.GetUser)
		//api.POST("/users", userController.CreateUser)
		//api.PUT("/users/:id", userController.UpdateUser)
		//api.DELETE("/users/:id", userController.DeleteUser)
	}

	// Serve static files
	// Serve static files with middleware
	// r.Use(middleware.ServeStaticOrAPI())
	// r.Static("/static", "./static")
	// r.StaticFile("/favicon.ico", "./static/favicon.ico")

	// // Handle React Routes
	// r.NoRoute(func(c *gin.Context) {
	// 	c.File("./static/index.html")
	// })

}
