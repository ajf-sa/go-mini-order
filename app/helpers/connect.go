package helpers

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectDB(url string, cfg ...string) *gorm.DB {
	log := &gorm.Config{}
	if len(cfg) > 0 {
		log.Logger = logger.Default.LogMode(logger.Silent)
	}
	db, err := gorm.Open(postgres.Open(Config(url)), log)
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
