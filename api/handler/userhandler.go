package handler

import (
	"net/http"

	"github.com/adilsonmenechini/simplesrestapi/app/entity"
	"github.com/adilsonmenechini/simplesrestapi/app/presenter"
	"github.com/adilsonmenechini/simplesrestapi/app/usecase"
	"github.com/google/uuid"

	"github.com/gofiber/fiber/v2"
)

func UserNewHandler(s usecase.UserService) UserHandler {
	return &userhandler{service: s}
}

func (h userhandler) AddUser() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var user entity.User
		err := c.BodyParser(&user)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.UserErrorResponse(err))
		}

		result, err := h.service.Creater(&user)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.UserErrorResponse(err))
		}
		return c.JSON(presenter.UserSuccessResponse(result))
	}
}

func (h userhandler) UpdateUser() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var userreq presenter.UserRequest
		err := c.BodyParser(&userreq)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.UserErrorResponse(err))
		}

		id := c.Params("id")
		if id == uuid.Nil.String() {
			return c.Status(http.StatusInternalServerError).JSON("Please ensure that :id is an integer")
		}
		user, err := h.service.Fetch(id)
		if err != nil {
			return err
		}
		user.Name = userreq.Name
		user.Email = userreq.Email
		user.Password = userreq.Password

		result, err := h.service.Update(&user)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.UserErrorResponse(err))
		}
		return c.JSON(presenter.UserSuccessResponse(result))
	}

}

func (h userhandler) RemoveUser() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		if id == uuid.Nil.String() {
			return c.Status(http.StatusInternalServerError).JSON("Please ensure that :id is an integer")
		}
		err := h.service.Delete(id)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.UserErrorResponse(err))
		}
		return c.JSON(&fiber.Map{
			"status": true,
			"data":   "delete successfully",
			"err":    nil,
		})
	}

}

func (h userhandler) GetUser() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		if id == uuid.Nil.String() {
			return c.Status(http.StatusInternalServerError).JSON("Please ensure that :id is an integer")
		}
		user, err := h.service.Fetch(id)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.UserErrorResponse(err))
		}
		return c.JSON(presenter.UserSuccessResponse(&user))
	}
}

func (h userhandler) GetUsers() fiber.Handler {
	return func(c *fiber.Ctx) error {
		users, err := h.service.Fetchs()
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.UserErrorResponse(err))
		}
		return c.JSON(presenter.UsersFindResponse(users))
	}
}
