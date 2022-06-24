package handler

import (
	"github.com/adilsonmenechini/simplesrestapi/app/usecase"
	"github.com/gofiber/fiber/v2"
)

type UserHandler interface {
	AddUser() fiber.Handler
	UpdateUser() fiber.Handler
	RemoveUser() fiber.Handler
	GetUser() fiber.Handler
	GetUsers() fiber.Handler
}

type userhandler struct {
	service usecase.UserService
}
