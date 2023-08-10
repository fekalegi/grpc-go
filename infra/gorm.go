package infra

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"grpc-go/config"
	"log"
)

var dbInstance *gorm.DB

func NewGormDB() *gorm.DB {
	dsn := config.PostgreConfig()

	if dbInstance == nil {
		var err error

		DBInstance, err := gorm.Open(postgres.New(dsn), &gorm.Config{
			DisableForeignKeyConstraintWhenMigrating: false,
		})
		if err != nil {
			log.Fatalln("Can't connect to db : ", err)
		}

		dbInstance = DBInstance

		return dbInstance
	} else {
		return dbInstance
	}
}
