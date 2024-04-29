package models
import (
    "github.com/lib/pq"
)

type Team struct {
    ID        uint `gorm:"primaryKey"`
    Name      string
    ShortName string
    Players   pq.Int32Array `gorm:"type:integer[],sort:asc"`
    Region    string
}
