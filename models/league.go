package models

import (
    "gorm.io/gorm"
    "time"
)

type League struct {
    gorm.Model
    Name      string
    Users     []User
    UserTeams []UserTeam
    StartDate time.Time
    EndDate   time.Time
    Reward    string
}
