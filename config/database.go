package config

import (
	"database/sql"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DbConnection() *gorm.DB {
	sqlDB, err := sql.Open("pgx", "mydb_dsn")
	if err != nil {
		fmt.Printf("Error database connection: ", err)
	}
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})

	return gormDB
}
