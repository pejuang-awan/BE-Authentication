package di

import (
	"fmt"
	"github.com/pejuang-awan/BE-Authentication/internal/entity"
	"github.com/pejuang-awan/BE-Authentication/internal/shared/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgres(conf *config.Config) (*gorm.DB, error) {
	dbURL := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		conf.Database.Host,
		conf.Database.User,
		conf.Database.Password,
		conf.Database.DbName,
		conf.Database.Port,
		conf.Database.SSLMode,
		conf.Database.Timezone)

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(entity.User{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
