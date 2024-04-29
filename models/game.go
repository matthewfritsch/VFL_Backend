package models
type Game struct {
    ID            uint `gorm:"primaryKey"`
    Date          string
    Time          string
    MatchID       uint `gorm:"index"`
    Map           string
    RndsNotPlayed uint
}
