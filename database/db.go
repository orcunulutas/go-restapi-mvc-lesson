package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseSettings struct {
	Host              string
	Port              string
	User              string
	Pass              string
	Dbname            string
	Sslmode           string
	Timezone          string
	ConnectionTimeout string
}

func GenerateDSN(settings DatabaseSettings) string {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s connect_timeout=%s",
		settings.Host,
		settings.Port,
		settings.User,
		settings.Pass,
		settings.Dbname,
		settings.Sslmode,
		settings.ConnectionTimeout,
	)
	return dsn

}

var DB *gorm.DB

func ConnectDatabase(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	DB = db
	return db, nil

}
