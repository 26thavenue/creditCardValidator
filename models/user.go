package models

import (
    "time"

    "gorm.io/gorm"
)

type User struct {
    ID        uint           `gorm:"primaryKey;autoIncrement" json:"id"`
    Name      string         `json:"name"`
    Email     string         `json:"email"`
    CreatedAt time.Time      `json:"created_at"`
    UpdatedAt time.Time      `json:"updated_at"`
    CreditCards []*CreditCard `gorm:"foreignKey:UserID" json:"credit_cards,omitempty"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
    u.CreatedAt = time.Now()
    u.UpdatedAt = time.Now()
    return nil
}

func (u *User) BeforeUpdate(tx *gorm.DB) error {
    u.UpdatedAt = time.Now()
    return nil
}