package migrations

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//Postgres

func ConnectPostgresServer() *gorm.DB {
	dsn := "host=localhost user=postgres password=tin14091998 dbname=GoReactProject port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err.Error())
	}
	err = sqlDB.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("Connected!\n")
	return db
}

func MigrateUser() {
	db := ConnectPostgresServer()
	db.AutoMigrate(&Users{})
}

func MigrateAvatar() {
	db := ConnectPostgresServer()
	db.AutoMigrate(&Avatar{})
}
