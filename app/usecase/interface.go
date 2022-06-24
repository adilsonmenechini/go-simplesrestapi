package usecase

import (
	"github.com/adilsonmenechini/simplesrestapi/app/entity"
	"github.com/adilsonmenechini/simplesrestapi/app/presenter"
)

type Writer interface {
	Creater(user *entity.User) (*entity.User, error)
	Update(user *entity.User) (*entity.User, error)
	Delete(id string) error
}

type Reader interface {
	Fetch(id string) (entity.User, error)
	Fetchs() ([]presenter.UserResponse, error)
}

type UserService interface {
	Writer
	Reader
}
