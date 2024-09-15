package postgresql

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Connection gets connection of postgresql database
func Connection() (db *gorm.DB) {
	host := viper.Get("PGHOST")
	user := viper.Get("PGUSER")
	pass := viper.Get("PGPASSWORD")
	port := viper.Get("PGPORT")
	database := viper.Get("PGDATABASE")
	dsn := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=disable", host, port, user, pass, database)
	print(dsn + "\n")
	print("test code change 123")
	print("test code change 456")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	return db
}
