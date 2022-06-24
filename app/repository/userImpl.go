package repository

import (
	"github.com/adilsonmenechini/simplesrestapi/app/entity"
	"gorm.io/gorm"
)

func UserNewRepo(db *gorm.DB) UserRepository {
	return &userrepository{
		DB: db,
	}
}

func (r *userrepository) Insert(user *entity.User) (*entity.User, error) {
	user, err := entity.NewUser(user.Name, user.Email, user.Password)
	if err != nil {
		return nil, err
	}
	err = r.DB.Create(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *userrepository) Save(user *entity.User) (*entity.User, error) {
	err := r.DB.Save(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *userrepository) FindByID(id string) (entity.User, error) {
	var user entity.User
	if err := r.DB.Debug().Find(&user, "id = ?", id).Error; err != nil {
		return user, err
	}
	return user, nil

}

func (r *userrepository) FindByAll() ([]entity.User, error) {
	var user []entity.User
	if err := r.DB.Debug().Find(&user).Error; err != nil {
		return nil, err
	}
	return user, nil

}

func (r *userrepository) FindEmail(email string) (entity.User, error) {
	var user entity.User
	if err := r.DB.Debug().First(&user, "email = ?", email).Error; err != nil {
		return user, err
	}
	return user, nil

}

func (r *userrepository) Delete(id string) error {
	user, err := r.FindByID(id)
	if err != nil {
		return err
	}
	err = r.DB.Delete(&user).Error
	if err != nil {
		return err
	}
	return nil
}
