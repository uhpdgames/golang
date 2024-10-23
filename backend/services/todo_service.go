package services

import (
	"backend/models"
	"errors"

	"gorm.io/gorm"
)

type TodoService struct {
	DB *gorm.DB
}

func NewTodoService(db *gorm.DB) *TodoService {
	return &TodoService{DB: db}
}

func (s *TodoService) CreateTodo(userID uint, input *models.CreateTodoInput) (*models.Todo, error) {
	todo := &models.Todo{
		Title:       input.Title,
		Description: input.Description,
		UserID:      userID,
	}

	if err := s.DB.Create(todo).Error; err != nil {
		return nil, err
	}

	return todo, nil
}

func (s *TodoService) GetUserTodos(userID uint) ([]models.Todo, error) {
	var todos []models.Todo
	if err := s.DB.Where("user_id = ?", userID).Find(&todos).Error; err != nil {
		return nil, err
	}
	return todos, nil
}

func (s *TodoService) UpdateTodo(userID uint, todoID uint, input *models.UpdateTodoInput) (*models.Todo, error) {
	var todo models.Todo

	if err := s.DB.Where("id = ? AND user_id = ?", todoID, userID).First(&todo).Error; err != nil {
		return nil, errors.New("todo not found")
	}

	if input.Title != "" {
		todo.Title = input.Title
	}
	if input.Description != "" {
		todo.Description = input.Description
	}
	if input.Status != nil {
		todo.Status = *input.Status
	}

	if err := s.DB.Save(&todo).Error; err != nil {
		return nil, err
	}

	return &todo, nil
}

func (s *TodoService) DeleteTodo(userID uint, todoID uint) error {
	result := s.DB.Where("id = ? AND user_id = ?", todoID, userID).Delete(&models.Todo{})
	if result.RowsAffected == 0 {
		return errors.New("todo not found")
	}
	return result.Error
}
