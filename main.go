package main

import (
    "github.com/gin-gonic/gin"
    "fantasy-valorant/routes"
)

func main() {
    gin.SetMode(gin.DebugMode)
    router := gin.Default()
    routes.SetupRouter(router)
    router.Run(":8080")
}
