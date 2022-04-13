package config

import (
    "fiber-api/entities"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

var Database *gorm.DB
var DATABASE_URI string = "admin:admin@tcp(gormDB)/gorm"

func Connect() error {
    var err error

    Database, err = gorm.Open(mysql.Open(DATABASE_URI))

    if err != nil {
        panic(err)
    }

    Database.AutoMigrate(&entities.Dog{}, &entities.Login{})

    return nil
}
