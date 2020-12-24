package repository

import (
	"errors"

	"github.com/ydhnwb/elib-user-microservice/domain/entity"
	"github.com/ydhnwb/elib-user-microservice/infrastructure/security"
	"gorm.io/gorm"
)

//UserRepository is a contract, it is implemented on infrastructure layer
//because the we depends on gorm
type UserRepository interface {
	InsertUser(u entity.User) (entity.User, error)
	UpdateUser(u entity.User) (entity.User, error)
	VerifyCredential(email string, password string) interface{}
	FindByEmail(email string) interface{}
	Profile(userID string) entity.User
}

type userRepository struct {
	db *gorm.DB
}

//NewUserRepository creates a new instance of UserRepository
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (repo *userRepository) InsertUser(u entity.User) (entity.User, error) {
	u.Password = security.HashAndSalt([]byte(u.Password))
	err := repo.db.Save(&u).Error
	if err != nil {
		return u, err
	}
	return u, nil
}

func (repo *userRepository) UpdateUser(u entity.User) (entity.User, error) {
	if u.Password != "" {
		u.Password = security.HashAndSalt([]byte(u.Password))
	} else {
		var temp entity.User
		repo.db.Find(&temp, u.ID)
		u.Password = temp.Password
	}
	err := repo.db.Save(&u).Error
	if err != nil {
		return u, err
	}
	return u, nil
}

func (repo *userRepository) VerifyCredential(email string, password string) interface{} {
	var user entity.User
	res := repo.db.Where("email = ?", email).Take(&user)
	if res.Error == nil {
		return user
	}
	return nil
}

func (repo *userRepository) FindByEmail(email string) interface{} {
	user := entity.User{}
	err := repo.db.Where("email = ?", email).Take(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}
	return user
}

func (repo *userRepository) Profile(userID string) entity.User {
	var user entity.User
	repo.db.Find(&user, userID)
	return user
}
