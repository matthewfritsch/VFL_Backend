package controllers

import (
    "fantasy-valorant/models"
    "fmt"
    "github.com/gin-gonic/gin"
    "net/http"
    "gorm.io/gorm"
)

func SearchPlayers(c *gin.Context, db *gorm.DB) {
    searchQuery := c.Query("query")
    if searchQuery == "" {
        // Instead of sending an error, send an empty or a special clear message
        // c.HTML(http.StatusOK, "players_search_results.html", gin.H{
        //     "players": nil,
        // })
        return
    }
    fmt.Printf("Querying for '%s'\n", searchQuery)
    ign_players, role_players, team_players, err := models.SearchPlayers(db, searchQuery)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    fmt.Println("Query complete! Passing to template...")
    // c.HTML(http.StatusOK, "players_search_results.html", gin.H{
    //     "ign_players": ign_players,
    //     "role_players": role_players,
    //     "team_players": team_players,
    // })
    // fmt.Println("Template loaded!")
    _ = ign_players
    _ = role_players
    _ = team_players
}

func GetPlayerByID(c *gin.Context) {
    id := c.Param("id")
    var player models.Player
    // TODO: Retrieve a player by ID from the database
    _ = id

    c.JSON(http.StatusOK, player)
}
