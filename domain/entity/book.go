package entity

import "time"

// Book represents Books table from database
type Book struct {
	ID          uint64    `gorm:"primary_key:auto_increment" json:"id"`
	Title       string    `gorm:"type:varchar(100)" json:"title" binding:"min:1;max=100"`
	Description string    `gorm:"type:text" json:"description"`
	UserID      uint64    `gorm:"not null" json:"-"`
	User        User      `json:"user" gorm:"foreignkey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
