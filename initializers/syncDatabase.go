package initializers

import (
	"chat-app/models"
)

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
