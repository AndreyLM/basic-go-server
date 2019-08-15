package routes

import (
	"github.com/andreylm/basic-go-server.git/pkg/db"
	"github.com/andreylm/basic-go-server.git/pkg/router"
	"github.com/andreylm/basic-go-server.git/pkg/server/handlers/auth"
)

// GetRoutes - gets basic routes
func GetRoutes(db db.DB) router.Routes {
	return router.Routes{
		router.Route{Name: "Auth", Path: "/registration", Method: "POST", HandlerFunc: auth.Register(db)},
		router.Route{Name: "Login", Path: "/login", Method: "POST", HandlerFunc: auth.Login(db)},
		router.Route{Name: "CheckAuth", Path: "/check-auth", Method: "GET", HandlerFunc: auth.CheckAuth(db)},
	}
}
