package auto

import (
	"api/libs/database"
	"api/models"
	"log"
)

// Sync autogenerates the tables and records
func Sync() {
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	err = db.DropTableIfExists(&models.User{}).Error
	if err != nil {
		log.Fatal(err)
	}
	err = db.AutoMigrate(&models.User{}).Error
	if err != nil {
		log.Fatal(err)
	}

	for i, _ := range users {
		err = db.Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatal(err)
		}
		// console.Pretty(posts[i])
	}
}
