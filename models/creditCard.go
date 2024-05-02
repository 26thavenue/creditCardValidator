package models

import (
    "time"

    "gorm.io/gorm"
)

type CreditCard struct {
    ID        uint       `gorm:"primaryKey;autoIncrement" json:"id"`
    Number    string     `gorm:"uniqueIndex" json:"number"`
    IsValid   bool       `json:"is_valid"`
    UserID    uint       `json:"user_id"`
    CreatedAt time.Time  `json:"created_at"`
    UpdatedAt time.Time  `json:"updated_at"`
    User      User       `gorm:"foreignKey:UserID" json:"-"`
}

func (cc *CreditCard) BeforeCreate(tx *gorm.DB) error {
    cc.CreatedAt = time.Now()
    cc.UpdatedAt = time.Now()
    return nil
}

func (cc *CreditCard) BeforeUpdate(tx *gorm.DB) error {
    cc.UpdatedAt = time.Now()
    return nil
}