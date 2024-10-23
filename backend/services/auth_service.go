package services

import (
	"backend/models"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService struct {
	db *gorm.DB
}

func NewAuthService(db *gorm.DB) *AuthService {
	return &AuthService{db: db}
}

func (s *AuthService) Register(input *models.RegisterInput) (*models.User, error) {
	var existingUser models.User
	if result := s.db.Where("email = ?", input.Email).First(&existingUser); result.Error == nil {
		return nil, errors.New("email already exists")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: string(hashedPassword),
	}

	if err := s.db.Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

// auth
func (s *AuthService) Login(input *models.LoginInput) (*models.LoginResponse, error) {
	var user models.User

	// filter by email
	if err := s.db.Where("email = ?", input.Email).First(&user).Error; err != nil {
		return nil, errors.New("invalid email or password")
	}

	// check match password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return nil, errors.New("invalid email or password")
	}

	//   token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte("your_jwt_secret"))
	if err != nil {
		return nil, err
	}

	// hidden password
	user.Password = ""

	return &models.LoginResponse{
		Token: tokenString,
		User:  user,
	}, nil
}

// list users
func (s *AuthService) GetUsers() ([]models.User, error) {
	var users []models.User
	// if err := s.db.Find(&users).Error; err != nil {
	// 	return nil, err
	// }

	result := s.db.Select("id, name, email, created_at, updated_at").Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}
