package drivers

import "github.com/go-xorm/xorm"

type defaultDriver struct {
	db *xorm.Engine
}

// Find - finds models
func (x *defaultDriver) Close() error {
	return x.db.Close()
}

// Find - finds models
func (x *defaultDriver) Find(findBy, objects interface{}, limit, offset int32) (err error) {
	return x.db.Find(objects, findBy)
}

// Get - finds by conditions
func (x *defaultDriver) Get(model interface{}) (err error) {
	_, err = x.db.Get(model)
	return
}

// Exists - check if model exists
func (x *defaultDriver) Exists(model interface{}) (bool, error) {
	return x.db.Get(model)
}

// Update - updates model
func (x *defaultDriver) Update(id int64, model interface{}) (err error) {
	_, err = x.db.ID(id).Update(model)
	return
}

// Store - inserts model
func (x *defaultDriver) Store(model interface{}) (err error) {
	_, err = x.db.Insert(model)
	return
}

// Delete - deletes model
func (x *MySQLDriver) Delete(id int64, model interface{}) (err error) {
	_, err = x.db.ID(id).Delete(model)
	return
}
