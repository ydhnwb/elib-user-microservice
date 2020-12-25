package persistence

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/ydhnwb/elib-user-microservice/domain/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//SetupDatabaseConnection is create a connection to database when server boot up
func SetupDatabaseConnection() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		panic("Failed to load env")
	}

	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbPort := "5432"

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", dbHost, dbUser, dbPass, dbName, dbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		println(err.Error())
		panic("Cannot connect to database!")
	}
	println("Db is connected!")
	db.AutoMigrate(&entity.User{})
	return db
}

//CloseDatabaseConnection will close connection to database
func CloseDatabaseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed when close a connection from database")
	}
	// defer dbSql.Close()
	dbSQL.Close()
}
