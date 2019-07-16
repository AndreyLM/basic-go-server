package xorm

import (
	"errors"

	DBConfig "github.com/andreylm/basic-go-server.git/pkg/db/xorm/config"
	"github.com/andreylm/basic-go-server.git/pkg/db/xorm/drivers"

	"github.com/andreylm/basic-go-server.git/pkg/db"
)

// GetDBDriverFactory - gets driver for db
func GetDBDriverFactory(config *DBConfig.ConnectionConfig) (db.DB, error) {
	switch config.DriverType {
	case DBConfig.MySQLType:
		return drivers.NewMySQLDriver(config), nil
	default:
		return nil, errors.New("Cannot get driver")
	}
}
