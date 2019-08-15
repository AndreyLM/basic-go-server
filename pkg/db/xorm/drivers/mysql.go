package drivers

import (
	"fmt"

	// mysql - mysql driver
	_ "github.com/go-sql-driver/mysql"

	"github.com/andreylm/basic-go-server.git/pkg/db"
	"github.com/go-xorm/xorm"
)

// MySQLDriver - mysql driver
type MySQLDriver struct {
	config *db.ConnectionConfig
	defaultDriver
}

// NewMySQLDriver - new mysql driver
func NewMySQLDriver(config *db.ConnectionConfig) *MySQLDriver {
	return &MySQLDriver{config: config}
}

// Connect - creates connection to db
func (x *MySQLDriver) Connect() (err error) {
	x.db, err = xorm.NewEngine(
		"mysql",
		fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?charset=utf8",
			x.config.User,
			x.config.Password,
			x.config.Host,
			x.config.Port,
			x.config.Database,
		),
	)
	if err != nil {
		return
	}

	if err = x.db.Ping(); err != nil {
		return
	}

	return
}
