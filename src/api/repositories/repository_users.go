package repositories

import (
	"api/models"
)

// UserRepository is the interface User CRUD
type UserRepository interface {
	Create(models.User) (models.User, error)
	Find(map[string][]interface{}) ([]models.User, error)
	FindByID(uint32) (models.User, error)
	Update(uint32, models.User) (int64, error)
	Delete(uint32) (int64, error)
}
