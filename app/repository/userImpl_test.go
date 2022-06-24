package repository_test

import (
	"fmt"
	"log"
	"testing"

	"github.com/adilsonmenechini/simplesrestapi/api/utils/db"
	"github.com/adilsonmenechini/simplesrestapi/app/entity"
	"github.com/adilsonmenechini/simplesrestapi/app/repository"
	"github.com/bxcodec/faker/v3"
)

func TestUserrepository_Create(t *testing.T) {

	var usr entity.User
	usr.Name = faker.Name()
	usr.Email = faker.Email()
	usr.Password = "1234567890abc"
	db := db.ConnectMySQL()
	r := repository.UserNewRepo(db)
	user, err := r.Insert(&usr)
	if err != nil {
		log.Fatalf("Errro ## %v", err)
	}
	fmt.Println(user)
}
func TestUserrepository_FindByAll(t *testing.T) {

	db := db.ConnectMySQL()
	r := repository.UserNewRepo(db)
	userf, err := r.FindByAll()

	if err != nil {
		log.Fatalf("Errro ## %v", err)
	}
	fmt.Println(userf)
}

func TestUserrepository_FindEmail(t *testing.T) {

	db := db.ConnectMySQL()
	r := repository.UserNewRepo(db)
	userf, err := r.FindEmail("WBeLsLI@vsPPExe.info")

	if err != nil {
		log.Fatalf("Errro ## %v", err)
	}
	fmt.Println(userf)
}
