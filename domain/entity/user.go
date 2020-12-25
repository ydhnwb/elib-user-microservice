package entity

import "time"

//User struct represent users table in database
type User struct {
	ID       uint64 `gorm:"primary_key:auto_increment" json:"id"`
	Name     string `gorm:"type:varchar(255)" json:"name"`
	Email    string `gorm:"uniqueIndex;type:varchar(255)" json:"email"`
	Password string `gorm:"->;<-;not null" json:"-"`
	Token    string `gorm:"-" binding:"-" json:"token,omitempty"`
	// Books     *[]Book   `json:"books,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
