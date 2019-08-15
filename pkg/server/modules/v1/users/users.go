package users

import (
	"github.com/andreylm/basic-go-server.git/pkg/db"
	"github.com/andreylm/basic-go-server.git/pkg/router"
)

// Module - module
type Module struct{}

// Init - initialization
func (m *Module) Init(r *router.Router, db db.DB) error {
	return nil
}
