package modules

import (
	"github.com/andreylm/basic-go-server.git/pkg/db"
	"github.com/andreylm/basic-go-server.git/pkg/router"
)

// Module - module interface
type Module interface {
	Init(r *router.Router, db db.DB) error
}
