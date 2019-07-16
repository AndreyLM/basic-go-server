package drivers

import (
	"github.com/andreylm/basic-go-server.git/pkg/db/xorm/config"
	"github.com/go-xorm/xorm"
)

// MySQLDriver - mysql driver
type MySQLDriver struct {
	config *config.ConnectionConfig
	db     *xorm.Engine
}

// NewMySQLDriver - new mysql driver
func NewMySQLDriver(config *config.ConnectionConfig) *MySQLDriver {
	return &MySQLDriver{config: config}
}

// Connect - creates connection to db
func (x *MySQLDriver) Connect() (err error) {
	x.db, err = xorm.NewEngine("mysql", "user:password@tcp(localhost:3310)/test?charset=utf8")
	if err != nil {
		return
	}

	if err = x.db.Ping(); err != nil {
		return
	}
	return
}

// Find - finds models
func (x *MySQLDriver) Find(findBy, objects interface{}, limit, offset int32) (err error) {
	return x.db.Find(objects, findBy)
}

// Get - finds by conditions
func (x *MySQLDriver) Get(model interface{}) (err error) {
	_, err = x.db.Get(model)
	return
}

// Exists - check if model exists
func (x *MySQLDriver) Exists(model interface{}) (bool, error) {
	return x.db.Get(model)
}

// Update - updates model
func (x *MySQLDriver) Update(id int64, model interface{}) (err error) {
	_, err = x.db.ID(id).Update(model)
	return
}

// Store - inserts model
func (x *MySQLDriver) Store(model interface{}) (err error) {
	_, err = x.db.Insert(model)
	return
}

// Delete - deletes model
func (x *MySQLDriver) Delete(id int64, model interface{}) (err error) {
	_, err = x.db.ID(id).Delete(model)
	return
}
