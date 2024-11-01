package config

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DbOpenConnection(host, username, password, dbName, port string) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", host, username, password, dbName, port)

	db, err := gorm.Open(postgres.Open(dsn))

	if err != nil {
		return nil, err
	}

	return db, nil
}

func LoadDatabaseEnvs() map[string]string {
	host := os.Getenv("ADEGAS_DB_HOST")
	dbName := os.Getenv("ADEGAS_DB_NAME")

	dbUsername := os.Getenv("ADEGAS_DB_USER")
	dbPassword := os.Getenv("ADEGAS_DB_PASSWORD")

	port := os.Getenv("ADEGAS_DB_PORT")

	return map[string]string{
		"host":       host,
		"dbName":     dbName,
		"dbUsername": dbUsername,
		"dbPassword": dbPassword,
		"dbPort":     port,
	}
}
