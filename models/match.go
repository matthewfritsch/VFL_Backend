package models 
type Match struct {
    ID         uint `gorm:"primaryKey"`
    EventID    uint `gorm:"index"`
    Completed  bool
    Team1      string
    Team1ID    uint
    Team1Score uint
    Team2      string
    Team2ID    uint
    Team2Score uint
}
