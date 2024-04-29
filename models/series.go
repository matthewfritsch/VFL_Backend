package models

type Series struct {
    ID         uint `gorm:"primaryKey"`
    Name       string
    SeriesPath string //https://vlr.gg/{ID}
}
