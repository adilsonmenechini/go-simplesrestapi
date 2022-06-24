package presenter

import (
	"github.com/adilsonmenechini/simplesrestapi/app/entity"
	"github.com/gofiber/fiber/v2"
)

type UserResponse struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"`
	Token    string `json:"token"`
}

type UserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `json:"token"`
}

type AuthResponse struct {
	Token string `json:"token"`
}

func UserSuccessResponse(data *entity.User) *fiber.Map {
	user := UserResponse{
		ID:       data.ID,
		Name:     data.Name,
		Email:    data.Email,
		Password: data.Password,
		Token:    data.Token,
	}
	return &fiber.Map{
		"status": true,
		"data":   user,
		"error":  nil,
	}
}

func AuthSuccessResponse(user *entity.User) *fiber.Map {
	return &fiber.Map{
		"status": true,
		"data":   user,
		"error":  nil,
	}
}
func UserFindResponse(data *UserResponse) *fiber.Map {
	return &fiber.Map{
		"status": true,
		"data":   data,
		"error":  nil,
	}
}

func UsersFindResponse(data []UserResponse) *fiber.Map {
	return &fiber.Map{
		"status": true,
		"data":   data,
		"error":  nil,
	}
}

func UserErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"status": false,
		"data":   "",
		"error":  err.Error(),
	}
}
