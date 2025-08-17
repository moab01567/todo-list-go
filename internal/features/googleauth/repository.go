package googleauth

import "gorm.io/gorm"

type LoginSessions struct {
	SessionToken string `gorm:"primarykey"`
	UserId       string `json:"user_id"`
}

type Repository struct {
	db *gorm.DB
}

func NewRepository() *Repository {
	return &Repository{}
}
