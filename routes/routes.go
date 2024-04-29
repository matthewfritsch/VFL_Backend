package routes

import (
    "fantasy-valorant/controllers"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

func SetupRouter(router *gin.Engine, db *gorm.DB) {
    router.POST("/register", controllers.RegisterUser)
    router.POST("/login", controllers.LoginUser)
    router.GET("/search", func(c *gin.Context) {
        controllers.SearchPlayers(c, db)
    })
    router.GET("/performance/:playerId", controllers.CalculatePlayerPerformance)
}
