package user

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(dbFilePath string) (*UserRepo, error) {
	db, err := gorm.Open(sqlite.Open(dbFilePath), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	if err := db.AutoMigrate(&Account{}); err != nil {
		return nil, err
	}

	return &UserRepo{db: db}, nil
}

func (repo *UserRepo) Create(user *Account) error {
	return nil
}

func (repo *UserRepo) Get(id uint) (*Account, error) {
	return nil, nil
}

func (repo *UserRepo) Update(user *Account) error {
	return nil
}

func (repo *UserRepo) Delete(id uint) error {
	return nil
}

func (repo *UserRepo) GetAll() ([]*Account, error) {
	return nil, nil
}
