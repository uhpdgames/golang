package controllers

import (
	"backend/models"
	"backend/services"
	"backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService *services.AuthService
}

func NewAuthController(authService *services.AuthService) *AuthController {
	return &AuthController{authService: authService}
}

func (c *AuthController) Register(ctx *gin.Context) {
	var input models.RegisterInput

	// Validate input
	if err := ctx.ShouldBindJSON(&input); err != nil {
		errors := utils.GetValidationErrors(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"errors": errors,
		})
		return
	}

	user, err := c.authService.Register(&input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	// password response
	user.Password = ""
	ctx.JSON(http.StatusCreated, gin.H{
		"status": "success",
		"data":   user,
	})
}

func (c *AuthController) Login(ctx *gin.Context) {
	var input models.LoginInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	response, err := c.authService.Login(&input)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (c *AuthController) GetUsers(ctx *gin.Context) {
	users, err := c.authService.GetUsers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, users)
}
