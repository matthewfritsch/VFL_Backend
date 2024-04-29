package main

import (
    "fantasy-valorant/routes"
    "github.com/gin-gonic/gin"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "gorm.io/gorm/logger"
    "os"
    "strings"
)

func InitDB() *gorm.DB {
    bdsn, _ := os.ReadFile(".env")
    dsn := strings.Trim(string(bdsn), "\n")
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
        Logger: logger.Default.LogMode(logger.Info),
    })
    if err != nil {
        return nil
    }
    return db
}

func main() {
    db := InitDB()
    router := gin.Default()
    routes.SetupRouter(router, db)
    router.Run(":8080")
}
