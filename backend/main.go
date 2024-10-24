package main

import (
	"backend/config"
	"backend/controllers"
	"backend/middleware"
	"backend/routes"
	"backend/services"
	"backend/utils"
	"fmt"

	"github.com/gin-gonic/gin"
)

const FRONT_END_URL string = "http://localhost:3000/"
const BACK_END_PORT string = ":8080"

func main() {

	db := config.ConnectDB()
	// Initialize services
	//userService := services.NewUserService(db)
	authService := services.NewAuthService(db)
	todoService := services.NewTodoService(db)
	// Initialize controllers
	//userController := controllers.NewUserController(userService)
	authController := controllers.NewAuthController(authService)
	todoController := controllers.NewTodoController(todoService)
	r := gin.Default() // has middleware by Gin.
	//r := gin.New() || //empty middleware
	// middleware Logger routes
	r.Use(gin.Logger())
	r.Use(func(c *gin.Context) {
		fmt.Printf("Request: %s %s\n", c.Request.Method, c.Request.URL)
		fmt.Printf("Headers: %v\n", c.Request.Header)

		c.Next()
	})

	// r.Use(gin.Recovery())
	//custom
	//r.Use(utils.Logger())
	//r.Use(utils.CORSMiddleware())
	// 1 request per second with a burst of 5
	r.Use(utils.RateLimitMiddleware(1, 5))
	// error hanlder
	r.Use(utils.ErrorHandler())
	r.Use(middleware.CORSMiddleware())

	routes.SetupRoutes(r, todoController, authController)

	// r.Use(cors.New(cors.Config{
	// 	AllowOrigins:     []string{FRONT_END_URL},
	// 	AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
	// 	AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
	// 	AllowCredentials: true,
	// 	AllowOriginFunc: func(origin string) bool {
	// 		return origin == "http://localhost"
	// 	},
	// 	MaxAge: 12 * time.Hour,
	// }))

	r.Run(BACK_END_PORT)
}
