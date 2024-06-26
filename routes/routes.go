package routes

import (
    "fantasy-valorant/auth"
    "fantasy-valorant/controllers"
    "fantasy-valorant/utils"
    "github.com/gin-gonic/gin"
)

func SetupRouter(router *gin.Engine) {
    heroku_db := utils.InitDB()
    supa_auth := utils.InitSupa()
    authFunc := func(c *gin.Context) {
        auth.AuthMiddleware(supa_auth, c)
    }

    router.POST("/register", func(c *gin.Context) {
        auth.Register(supa_auth, c)
    })
    router.POST("/signin", func(c *gin.Context) {
        auth.SignIn(supa_auth, c)
    })
    router.GET("/protected/search", authFunc, func(c *gin.Context) {
        controllers.SearchPlayers(heroku_db, c)
    })
    router.GET("/performance/:playerId", controllers.CalculatePlayerPerformance)
}
