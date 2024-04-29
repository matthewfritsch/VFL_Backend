package models
type Player struct {
    ID        uint `gorm:"primaryKey"`
    Name      string
    Ign       string `gorm:"index"`
    Role      string
}
