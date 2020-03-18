package crud

import (
	"api/models"
	"api/utils/channels"
	"errors"
	"time"

	"github.com/jinzhu/gorm"
)

// RepositoryUsersCRUD is the struct for the User CRUD
type RepositoryUsersCRUD struct {
	db    *gorm.DB
	model *gorm.DB
}

// NewRepositoryUsersCRUD returns a new repository with DB connection
func NewRepositoryUsersCRUD(db *gorm.DB) *RepositoryUsersCRUD {
	model := db.Model(&models.User{})
	return &RepositoryUsersCRUD{db, model}
}

// Create returns a new user created or an error
func (r *RepositoryUsersCRUD) Create(user models.User) (models.User, error) {
	var err error
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		err = r.model.Create(&user).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return user, nil
	}
	return models.User{}, err
}

// FindAll returns all the users from the DB
func (r *RepositoryUsersCRUD) Find(params map[string][]interface{}) ([]models.User, error) {
	var err error
	users := []models.User{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		model := r.model
		for key := range params {
			model = model.Where(key, params[key]...)
		}
		err = model.Limit(100).Find(&users).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return users, nil
	}
	return nil, err
}

// FindByID returns an user from the DB
func (r *RepositoryUsersCRUD) FindByID(uid uint32) (models.User, error) {
	var err error
	user := models.User{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		err = r.model.Where("id = ?", uid).Take(&user).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return user, nil
	}

	if gorm.IsRecordNotFoundError(err) {
		return models.User{}, errors.New("User Not Found")
	}
	return models.User{}, err
}

// Update updates an user from the DB
func (r *RepositoryUsersCRUD) Update(uid uint32, user models.User) (int64, error) {
	var rs *gorm.DB
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		rs = r.model.Where("id = ?", uid).Take(&models.User{}).UpdateColumns(
			map[string]interface{}{
				"nickname":   user.Nickname,
				"email":      user.Email,
				"updated_at": time.Now(),
			},
		)
		ch <- true
	}(done)

	if channels.OK(done) {
		if rs.Error != nil {
			return 0, rs.Error
		}

		return rs.RowsAffected, nil
	}
	return 0, rs.Error
}

// Delete removes an user from the DB
func (r *RepositoryUsersCRUD) Delete(uid uint32) (int64, error) {
	var rs *gorm.DB
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		rs = r.model.Where("id = ?", uid).Take(&models.User{}).Delete(&models.User{})
		ch <- true
	}(done)

	if channels.OK(done) {
		if rs.Error != nil {
			return 0, rs.Error
		}

		return rs.RowsAffected, nil
	}
	return 0, rs.Error
}
