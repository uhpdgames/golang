package services

import (
	"backend/models"

	"gorm.io/gorm"
)

type UserService struct {
	DB *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{DB: db}
}

func (s *UserService) GetAllUsers() ([]models.User, error) {
	var users []models.User
	result := s.DB.Find(&users)
	return users, result.Error
}

func (s *UserService) GetUserByID(id uint) (models.User, error) {
	var user models.User
	result := s.DB.First(&user, id)
	return user, result.Error
}

func (s *UserService) CreateUser(user *models.User) error {
	return s.DB.Create(user).Error
}

func (s *UserService) UpdateUser(user *models.User) error {
	return s.DB.Save(user).Error
}

func (s *UserService) DeleteUser(id uint) error {
	return s.DB.Delete(&models.User{}, id).Error
}
