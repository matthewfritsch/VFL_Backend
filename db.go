package main

import (
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "os"
    "strings"
)

func InitDB() *gorm.DB {
    bdsn, _ := os.ReadFile(".env")
    dsn := strings.Trim(string(bdsn), "\n")
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        return nil
    }

    db.AutoMigrate(&Series{}, &Event{}, &Match{}, &Game{}, &PlayerPerformance{}, &Team{}, &Player{})
    return db
}

func SavePlayerPerformance(db *gorm.DB, pperf PlayerPerformance) {
    var tempPerf PlayerPerformance
    db.Where(&PlayerPerformance{PlayerID:pperf.PlayerID, GameID:pperf.GameID}).First(&tempPerf)
    if tempPerf.PlayerID == pperf.PlayerID {
        pperf.ID = tempPerf.ID
        db.Updates(&pperf)
    } else {
        db.Create(&pperf)
    }
}

func SavePlayerPerformances(db *gorm.DB, pperfs *[]PlayerPerformance) {
    for idx := range *pperfs {
        SavePlayerPerformance(db, (*pperfs)[idx])
    }
}

func SaveTeam(db *gorm.DB, team Team) {
    db.Save(&team)
    db.Model(&Match{}).Where("team1_id = ? AND team1 = ?", 0, team.Name).Update("team1_id", team.ID)
    db.Model(&Match{}).Where("team2_id = ? AND team2 = ?", 0, team.Name).Update("team2_id", team.ID)
}

func SaveTeams(db *gorm.DB, teams *[]Team) {
    for _, team := range *teams {
        SaveTeam(db, team)
    }
}

