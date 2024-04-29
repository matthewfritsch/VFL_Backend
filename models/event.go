package models
//https://vlr.gg/event/matches/{ID}/?series_id=all
type Event struct {
    ID       uint `gorm:"primaryKey"`
    SeriesID uint
    Name     string
    Region   string
}
