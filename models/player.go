package models

import (
    "fmt"
    "github.com/lib/pq"
    "gorm.io/gorm"
    "sort"
    "strings"
)

type Player struct {
    ID        uint `gorm:"primaryKey"`
    Name      string
    Ign       string `gorm:"index"`
    Role      string
}

type PlayerResult struct {
    ID   int32
    Ign  string
    Role string
    Team string
}

func SearchPlayers(db *gorm.DB, query string) ([]PlayerResult, []PlayerResult, []PlayerResult, error) {
    var results []PlayerResult
    if len(query) < 3 {
        return results, results, results, nil
    }

    ign, role := SimilarIgnOrRole(db, query)
    team := SimilarTeam(db, query)
    fmt.Printf("Result contains %d items\n", len(results))
    return ign, role, team, nil
}

func SimilarIgnOrRole(db *gorm.DB, query string) ([]PlayerResult, []PlayerResult) {
    var ign_results []PlayerResult
    var role_results []PlayerResult

    var ign_players []Player
    result := db.Where("ign ILIKE ?", "%"+query+"%").Find(&ign_players)
    if result.Error != nil {
        fmt.Println("IgnOrRole query is an error.")
    }

    var role_players []Player
    result = db.Where("role ILIKE ?", "%"+query+"%").Find(&role_players)
    if result.Error != nil {
        fmt.Println("IgnOrRole query is an error.")
    }

    var players []Player
    players = append(role_players, ign_players...)

    playerIDs_uncast := make([]int32, len(players))
    for i, player := range players {
        playerIDs_uncast[i] = int32(player.ID)
    }
    playerIDs := pq.Array(playerIDs_uncast)

    var teams []Team
    err := db.Find(&teams, "players && ?", playerIDs).Error
    if err != nil {
        panic(err)
    }

    playerToTeamMap := make(map[uint]string)
    for _, team := range teams {
        for _, playerID := range team.Players {
            playerToTeamMap[uint(playerID)] = team.Name
        }
    }

    for i := range ign_players {
        var res PlayerResult
        if team, exists := playerToTeamMap[ign_players[i].ID]; exists {
            res.ID = int32(ign_players[i].ID)
            res.Ign = ign_players[i].Ign
            res.Role = ign_players[i].Role
            res.Team = team
            ign_results = append(ign_results, res)
        }
    }

    for i := range role_players {
        var res PlayerResult
        if team, exists := playerToTeamMap[role_players[i].ID]; exists {
            res.ID = int32(role_players[i].ID)
            res.Ign = role_players[i].Ign
            res.Role = role_players[i].Role
            res.Team = team
            role_results = append(role_results, res)
        }
    }


    sort.Slice(ign_results, func(i, j int) bool {
        return strings.ToLower(ign_results[i].Ign) < strings.ToLower(ign_results[j].Ign)
    })

    sort.Slice(role_results, func(i, j int) bool {
        return strings.ToLower(role_results[i].Ign) < strings.ToLower(role_results[j].Ign)
    })

    return ign_results, role_results
}

func SimilarTeam(db *gorm.DB, query string) []PlayerResult {
    var results []PlayerResult

    var teams []Team
    result := db.Where("name ILIKE ?", "%"+query+"%").Find(&teams)
    if result.Error != nil {
        fmt.Println("IgnOrRole query is an error")
    }

    for _, team := range teams {

        var player_ids []uint
        for p_idx := range team.Players {
            player_ids = append(player_ids, uint(team.Players[p_idx]))
        }

        var players []Player
        db.Where("id IN ?", player_ids).Find(&players)
        for _, player := range players {
            var res PlayerResult
            res.ID = int32(player.ID)
            res.Ign = player.Ign
            res.Role = player.Role
            res.Team = team.Name
            results = append(results, res)
        }
    }

    sort.Slice(results, func(i, j int) bool {
        return strings.ToLower(results[i].Ign) < strings.ToLower(results[j].Ign)
    })

    return results
}
