package models
type PlayerPerformance struct {
    ID       uint `gorm:"primaryKey"`
    PlayerID uint `gorm:"index"`
    TeamID   uint
    GameID   uint
    Name     string
    Team     string
    Kills    uint
    Deaths   uint
    Assists  uint
    ACS      uint
    P3K      uint
    P4K      uint
    P5K      uint
    V2       uint
    V3       uint
    V4       uint
    V5       uint
}
