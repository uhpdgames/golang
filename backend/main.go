package main

import (
	"backend/config"
	"backend/controllers"
	"backend/routes"
	"backend/services"
	"backend/utils"
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

const FRONT_END_URL string = "http://localhost:3000/"

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

	// Setup CORS

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000/"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * 60 * 600,
	}))
	// config := cors.DefaultConfig()
	// config.AllowOrigins = []string{FRONT_END_URL}
	// config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}
	// config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	// config.MaxAge = 12 * 60 * 60
	// r.Use(cors.New(config))

	routes.SetupRoutes(r, todoController, authController)

	//r.LoadHTMLFiles("templates/template1.html", "templates/template2.html")

	r.Run(":8080")
}
