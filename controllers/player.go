package controllers

import (
    "fantasy-valorant/models"
    "fmt"
    "github.com/gin-gonic/gin"
    "net/http"
    "gorm.io/gorm"
)

func SearchPlayers(db *gorm.DB, c *gin.Context) {
    //really want to consider if Go's net/rpc would be a better choice than gRPC.
    //a lot of the backend work here was useful in Gin/REST because we were hosting web pages too.
    //now it seems more realistic to host API endpoints, which are detached from web page routing, in net/rpc since they perform db queries and calculate responses.
    //on the other hand, POST with replies seems reasonable, as long as it doesn't violate HTTP/1.1 laws.
    searchQuery := c.Query("query")
    if searchQuery == "" {
        // Instead of sending an error, send an empty or a special clear message
        // c.HTML(http.StatusOK, "players_search_results.html", gin.H{
        //     "players": nil,
        // })
        fmt.Println("No query")
        return
    }
    fmt.Printf("Querying for '%s'\n", searchQuery)
    ign_players, role_players, team_players, err := models.SearchPlayers(db, searchQuery)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    fmt.Println("Query complete! Passing to template...")
    c.JSON(http.StatusOK, gin.H{
        "ign_players": ign_players,
        "role_players": role_players,
        "team_players": team_players,
    })
}

func GetPlayerByID(c *gin.Context) {
    id := c.Param("id")
    var player models.Player
    // TODO: Retrieve a player by ID from the database
    _ = id

    c.JSON(http.StatusOK, player)
}
