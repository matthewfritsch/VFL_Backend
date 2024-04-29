package controllers

import (
    "fantasy-valorant/models"
    "github.com/gin-gonic/gin"
    "net/http"
)

func GetTeam(c *gin.Context) {
    userId := c.GetInt("userId") // Assuming user authentication is implemented

    var selectedPlayers []models.Player
    // TODO: Retrieve selected players for the user from the database
    _ = userId

    var availablePlayers []models.Player
    // TODO: Retrieve available players from the database

    c.HTML(http.StatusOK, "team.html", gin.H{
        "SelectedPlayers":  selectedPlayers,
        "AvailablePlayers": availablePlayers,
    })
}

func AddPlayerToTeam(c *gin.Context) {
    userId := c.GetInt("userId") // Assuming user authentication is implemented
    playerId := c.PostForm("playerId")

    // TODO: Add the player to the user's fantasy team in the database
    _ = userId
    _ = playerId

    c.Redirect(http.StatusFound, "/team")
}

func RemovePlayerFromTeam(c *gin.Context) {
    userId := c.GetInt("userId") // Assuming user authentication is implemented
    playerId := c.PostForm("playerId")

    // TODO: Remove the player from the user's fantasy team in the database
    _ = userId
    _ = playerId

    c.Redirect(http.StatusFound, "/team")
}
