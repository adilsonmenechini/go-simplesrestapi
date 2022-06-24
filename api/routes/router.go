package routes

import (
	"github.com/adilsonmenechini/simplesrestapi/api/handler"
	"github.com/adilsonmenechini/simplesrestapi/api/utils/db"
	"github.com/adilsonmenechini/simplesrestapi/app/repository"
	"github.com/adilsonmenechini/simplesrestapi/app/usecase"
	"github.com/gofiber/fiber/v2"
)

func UserRouter(app fiber.Router) {
	db := db.ConnectPSQL()
	r := repository.UserNewRepo(db)
	serv := usecase.UserNewService(r)
	hdr := handler.UserNewHandler(serv)
	user := app.Group("/user")
	user.Get("/", hdr.GetUsers())
	user.Post("/", hdr.AddUser())
	user.Put("/:id", hdr.UpdateUser())
	user.Get("/:id", hdr.GetUser())
	user.Delete("/:id", hdr.RemoveUser())
}
