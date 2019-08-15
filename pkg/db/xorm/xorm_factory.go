package xorm

import (
	"errors"

	"github.com/andreylm/basic-go-server.git/pkg/db/xorm/drivers"

	"github.com/andreylm/basic-go-server.git/pkg/db"
)

// Factory - xorm factory
type Factory struct{}

// GetDBDriverFactory - gets driver for db
func (x *Factory) GetDBDriverFactory(config *db.ConnectionConfig) (db.DB, error) {
	switch config.DriverType {
	case db.MySQLType:
		return drivers.NewMySQLDriver(config), nil
	default:
		return nil, errors.New("Cannot get driver")
	}
}
