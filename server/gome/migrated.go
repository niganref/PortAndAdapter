package main

import (
    "fmt"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "server/gome/model"
)



func main() {
    dsn := "host=localhost user=test dbname=paa port=xxxx sslmode=disable TimeZone=Asia/Shanghai"
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }
    db.AutoMigrate(&models.Contents{})
    fmt.Println("migrated")
}
