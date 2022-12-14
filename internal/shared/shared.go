package shared

import (
	"github.com/labstack/echo/v4"
	"github.com/pejuang-awan/BE-Authentication/internal/shared/config"
	log "github.com/sirupsen/logrus"
	"go.uber.org/dig"
	"gorm.io/gorm"
)

type Deps struct {
	dig.In
	Config   *config.Config
	Logger   *log.Logger
	Server   *echo.Echo
	Database *gorm.DB
}

func (d *Deps) Close() {
	if err := d.Server.Close(); err != nil {
		d.Logger.Errorf("Failed to close the server: %v", err)
	}

	db, err := d.Database.DB()
	if err != nil {
		d.Logger.Errorf("Failed to get database connection: %v", err)
	}

	if err := db.Close(); err != nil {
		d.Logger.Errorf("Failed to close database connection: %v", err)
	}
}
