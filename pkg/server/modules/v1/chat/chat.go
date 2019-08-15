package chat

import (
	"github.com/andreylm/basic-go-server.git/pkg/db"
	Router "github.com/andreylm/basic-go-server.git/pkg/router"
	"github.com/andreylm/basic-go-server.git/pkg/server/modules/v1/chat/routes"
)

// Module - module
type Module struct{}

// Init - initialization
func (m *Module) Init(r *Router.Router, db db.DB) error {
	routes := routes.GetRoutes(db)
	subRouter := &Router.SubRouter{
		PathPrefix: "/v1",
		Routes:     routes,
	}

	r.AttachSubRouterWithMiddleware(subRouter)
	return nil
}
