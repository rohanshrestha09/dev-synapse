package configs

import (
	"log"
	"os"

	"github.com/rohanshrestha09/dev-synapse/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitializeDatabase() {

	var err error

	if DB, err = gorm.Open(mysql.Open(os.Getenv("DSN_LOCAL")), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	}); err != nil {
		log.Fatal("Error connecting to database")
	}

	if err = DB.AutoMigrate(&models.User{}, &models.Project{}, &models.Request{}, &models.Notification{}, &models.Chat{}, &models.Developer{}); err != nil {
		log.Fatal("Error while migrating")
	}

}
