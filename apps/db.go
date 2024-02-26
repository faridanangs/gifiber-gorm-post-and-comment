package apps

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

func InitializeDB(dbConfig DBConfig) *gorm.DB {
	dns := "host=localhost port=5432 user=postgres password=anangs dbname=latihan sslmode=disable TimeZone=Asia/Jakarta"

	DB, err := gorm.Open(postgres.Open(dns), &gorm.Config{
		Logger:      logger.Default.LogMode(logger.Info),
		PrepareStmt: true,
	})
	if err != nil {
		log.Fatal("Error at connection to database")
	}
	DB.Exec("SET search_path TO latihan")
	return DB
}
