package config

import (
	"gorm.io/driver/postgres"
	"os"
	"strings"
)

func PostgreConnection() string {
	dbConfigs := []string{
		"host=" + os.Getenv("DB_HOST"),
		"dbname=" + os.Getenv("DB_DATABASE"),
		"user=" + os.Getenv("DB_USERNAME"),
		"password=" + os.Getenv("DB_PASSWORD"),
		"sslmode=disable"}
	return strings.Join(dbConfigs, " ")
}

func PostgreConnectionLocal() string {
	dbConfigs := []string{
		"host=" + os.Getenv("DB_HOST"),
		"dbname=" + os.Getenv("DB_DATABASE"),
		"user=" + os.Getenv("DB_USERNAME"),
		"password=" + os.Getenv("DB_PASSWORD"),
		"port=" + os.Getenv("DB_PORT"),
		"sslmode=disable"}

	return strings.Join(dbConfigs, " ")
}

func PostgreConfig() postgres.Config {

	configPostgre := postgres.Config{
		PreferSimpleProtocol: false,
	}

	if os.Getenv("APP_ENV") == "local" {
		configPostgre.DSN = PostgreConnectionLocal()
	} else {
		configPostgre.DSN = PostgreConnection()
	}

	return configPostgre
}
