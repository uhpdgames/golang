package controllers

import (
	"net/http"
	"strconv"

	"backend/models"
	"backend/services"
	"backend/utils"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *services.UserService
}

func (u *UserController) HomePage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "Main website",
	})
}

func NewUserController(userService *services.UserService) *UserController {
	return &UserController{userService: userService}
}

func (c *UserController) GetUsers(ctx *gin.Context) {
	users, err := c.userService.GetAllUsers()
	if err != nil {
		utils.ResponseError(ctx, http.StatusInternalServerError, "Failed to get users")
		return
	}
	utils.ResponseSuccess(ctx, http.StatusOK, users)
}

func (c *UserController) GetUser(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	user, err := c.userService.GetUserByID(uint(id))
	if err != nil {
		utils.ResponseError(ctx, http.StatusNotFound, "User not found")
		return
	}
	utils.ResponseSuccess(ctx, http.StatusOK, user)
}

func (c *UserController) CreateUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		utils.ResponseError(ctx, http.StatusBadRequest, "Invalid input")
		return
	}

	if err := c.userService.CreateUser(&user); err != nil {
		utils.ResponseError(ctx, http.StatusInternalServerError, "Failed to create user")
		return
	}
	utils.ResponseSuccess(ctx, http.StatusCreated, user)
}

func (c *UserController) UpdateUser(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		utils.ResponseError(ctx, http.StatusBadRequest, "Invalid input")
		return
	}
	user.ID = uint(id)

	if err := c.userService.UpdateUser(&user); err != nil {
		utils.ResponseError(ctx, http.StatusInternalServerError, "Failed to update user")
		return
	}
	utils.ResponseSuccess(ctx, http.StatusOK, user)
}

func (c *UserController) DeleteUser(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	if err := c.userService.DeleteUser(uint(id)); err != nil {
		utils.ResponseError(ctx, http.StatusInternalServerError, "Failed to delete user")
		return
	}
	utils.ResponseSuccess(ctx, http.StatusOK, gin.H{"message": "User deleted successfully"})
}
