package commons

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/oneitworld-demo-crud-api-go/models"
)

func GetConnection() *gorm.DB {
	db, error := gorm.Open("mysql", "root:root@tcp(localhost:3306)/world?charset=utf8")

	if error != nil {
		log.Fatal(error)
	}
	return db

}

func Migrate() {

	db := GetConnection()
	defer db.Close()

	log.Println("Migrating..")
	db.AutoMigrate(&models.Persona{})
}
