package commons

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/oneitworld-demo-crud-api-go/models"
)

func GetConnection() *gorm.DB {
	// db, error := gorm.Open("mysql", "root:Kalifornia2024$@tcp(mysql-server:3306)/sys?charset=utf8")
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
	db.AutoMigrate(&models.Audit{})
}

// Escribe la Auditoria en Base de Datos
func WriteAudit(request *http.Request, apiResponseJSON []byte, apiName string, success bool, httpStatus int, channel string) {
	db := GetConnection()
	defer db.Close()

	// Formatear la fecha y hora con zona horaria
	formattedDateTimeWithZone := time.Now().Format("2006-01-02 15:04:05 MST")

	audit := models.Audit{
		Datetime:     formattedDateTimeWithZone,
		APIName:      apiName,
		IPAddress:    string(GetIP(request)),
		URL:          string(request.URL.Path),
		HTTPMethod:   string(request.Method),
		HTTPRequest:  string(""),
		HTTPResponse: string(apiResponseJSON),
		Success:      success,
		Status:       httpStatus,
		Channel:      channel,
	}
	fmt.Println(audit)

	error := db.Save(&audit).Error

	if error != nil {
		log.Fatal(error)
		return
	}

}
