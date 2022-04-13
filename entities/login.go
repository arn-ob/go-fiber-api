package entities

import "gorm.io/gorm"

type Login struct {
    gorm.Model
    Username     string `json:"username"`
    Password     string `json:"password"`
    Name         string `json:"name"`
	Address      string `json:"address"`
}