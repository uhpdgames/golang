package controllers

import (
	"backend/models"
	"backend/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TodoController struct {
	todoService *services.TodoService
}

func NewTodoController(todoService *services.TodoService) *TodoController {
	return &TodoController{todoService: todoService}
}

func (c *TodoController) CreateTodo(ctx *gin.Context) {
	var input models.CreateTodoInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := ctx.MustGet("user_id").(float64)
	todo, err := c.todoService.CreateTodo(uint(userID), &input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, todo)
}

func (c *TodoController) GetTodos(ctx *gin.Context) {
	userID := ctx.MustGet("user_id").(float64)
	todos, err := c.todoService.GetUserTodos(uint(userID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, todos)
}

func (c *TodoController) UpdateTodo(ctx *gin.Context) {
	todoID, _ := strconv.Atoi(ctx.Param("id"))
	userID := ctx.MustGet("user_id").(float64)

	var input models.UpdateTodoInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todo, err := c.todoService.UpdateTodo(uint(userID), uint(todoID), &input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, todo)
}

func (c *TodoController) DeleteTodo(ctx *gin.Context) {
	todoID, _ := strconv.Atoi(ctx.Param("id"))
	userID := ctx.MustGet("user_id").(float64)

	if err := c.todoService.DeleteTodo(uint(userID), uint(todoID)); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Todo deleted successfully"})
}

func (u *TodoController) HomePage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "Main website",
	})
}
