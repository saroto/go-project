package config

import (
	"database/sql"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DbConnection() {
	sqlDB, err := sql.Open("pgx", "mydb_dsn")
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
}
