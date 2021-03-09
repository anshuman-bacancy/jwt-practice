package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"column:Name"`
	Email    string `gorm:"column:Email"`
	Password string `gorm:"column:Password"`
	Role     string `gorm:"column:Role"`
}

type Config struct {
	ConnString string `mapstructure:"CONNECTION_STRING"`
	SigningKey string `mapstructure:"SIGNING_KEY"`
}
