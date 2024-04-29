package utils

import (
    supa "github.com/nedpals/supabase-go"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "gorm.io/gorm/logger"
    "os"
    "strings"
)

func ReadEnv() map[string]string {
    envvars := make(map[string]string)
    content, _ := os.ReadFile(".env")
    lines := strings.Split(string(content), "\n")
    for idx := range lines {
        line := lines[idx]
        split := strings.Split(line, "=")
        var key string
        var val string
        for idx := range len(split) {
            if idx == 0 {
                key = split[idx]
            } else {
                val += split[idx]
            }
        }
        envvars[key] = val
    }
    return envvars
}

func InitDB() *gorm.DB {
    envvars := ReadEnv()
    bdsn := envvars["PSQL_KEY"]
    dsn := strings.Trim(string(bdsn), "\n")
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
        Logger: logger.Default.LogMode(logger.Info),
    })
    if err != nil {
        return nil
    }
    return db
}

func InitSupa() *supa.Client {
    envvars := ReadEnv()
    supabaseUrl := envvars["SUPABASE_URL"]
    supabaseKey := envvars["SUPABASE_KEY"]
    supabase := supa.CreateClient(supabaseUrl, supabaseKey)
    return supabase
}


