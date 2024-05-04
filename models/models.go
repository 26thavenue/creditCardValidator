package models

import "time"

type User struct {
    ID           uint           `gorm:"primaryKey"`
    Name         string         `gorm:"not null"`
    Email        string         `gorm:"not null;unique"`
    CreatedAt    time.Time
    UpdatedAt    time.Time
    CreditCards  []CreditCard   `gorm:"foreignKey:UserID"`
}

type CreditCard struct {
    ID        uint       `gorm:"primaryKey"`
    Number    string     `gorm:"not null;unique"`
    IsValid   bool       `gorm:"not null"`
    UserID    uint       `gorm:"not null"`
    CreatedAt time.Time
    UpdatedAt time.Time
    User      User       `gorm:"constraint:OnDelete:CASCADE;"`
}