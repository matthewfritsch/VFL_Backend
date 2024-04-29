package models

import "gorm.io/gorm"

type User struct {
    gorm.Model
    Email        string `gorm:"uniqueIndex;not null"`
    Password     string `gorm:"not null"`
    DisplayName  string
    UserTeams    []UserTeam
}

/*

User contains an email, password, displayname, []UserTeam since they can have as many teams on as many leagues as they want
UserTeam contains a []PlayerID, Name, and League
League contains an id, name, []User, []UserTeam, startdate, enddate, reward
*/
