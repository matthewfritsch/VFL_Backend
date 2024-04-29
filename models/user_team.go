package models

import (
    "gorm.io/gorm"
    "github.com/lib/pq"
)

type UserTeam struct {
    gorm.Model
    Name     string
    UserID   uint
    User     User
    LeagueID uint
    League   League
    PlayerIDs pq.Int64Array `gorm:"type:integer[]"`
}
