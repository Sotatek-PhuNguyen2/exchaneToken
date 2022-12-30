package config

import (
	"change_money/models"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func getDatabase() *gorm.DB {
	err := godotenv.Load("././.env")
	if err != nil {
		fmt.Printf("Error loading .env file")
	}
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		Logger:                                   logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		fmt.Println("Error when connect to database " + err.Error())
		os.Exit(1)
	}
	fmt.Println("Connect to database successfully")
	err = db.AutoMigrate(
		&models.NetWork{},
		&models.EventLogs{},
		// &models.User{},
		// &models.Company{},
	)

	if err != nil {
		fmt.Println("Error when auto migrate database " + err.Error())
		os.Exit(1)
	}
	fmt.Println("Migrated database")

	return db
}

func InitDatabase() {
	DB = getDatabase()
}
