package entity

import (
	"errors"
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type User struct {
	gorm.Model
	ID       string `json:"id" gorm:"type:varchar(255);unique_index" valid:"notnull,uuid"`
	Name     string `json:"name" gorm:"type:varchar(255)" valid:"notnull"`
	Email    string `json:"email" gorm:"type:varchar(255);unique_index" valid:"notnull,email"`
	Password string `json:"-" gorm:"type:varchar(255)" valid:"notnull"`
	Token    string `json:"token" gorm:"type:varchar(255);unique_index" valid:"notnull,uuid"`
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	id := uuid.NewV4().String()
	createdAt := time.Now().UTC()
	token := uuid.NewV4().String()
	tx.Statement.SetColumn("ID", id)
	tx.Statement.SetColumn("Token", token)
	tx.Statement.SetColumn("CreatedAt", createdAt)

	return nil
}

func (user *User) BeforeSave(tx *gorm.DB) (err error) {
	updatedAt := time.Now()
	tx.Statement.SetColumn("UpdatedAt", updatedAt)
	// if user.Password != "" {
	// 	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	tx.Statement.SetColumn("Password", password)
	// }
	return nil
}
func NewUser(name string, email string, password string) (*User, error) {

	pwd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user := &User{
		Name:     name,
		Email:    email,
		Password: string(pwd),
	}
	if err = user.validate(); err != nil {
		return nil, err
	}

	return user, nil
}

func (user *User) IsCorrectPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}

func (user *User) validate() error {
	if user.Name == "" {
		return errors.New("bad user name")
	}
	if user.Email == "" {
		return errors.New("bad user email")
	}
	if user.Password == "" {
		return errors.New("bad user password")
	}

	return nil
}
