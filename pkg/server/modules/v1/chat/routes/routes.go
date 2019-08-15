package routes

import (
	"github.com/andreylm/basic-go-server.git/pkg/db"
	"github.com/andreylm/basic-go-server.git/pkg/router"
	"github.com/andreylm/basic-go-server.git/pkg/server/modules/v1/chat/handlers"
)

// GetRoutes - gets basic routes
func GetRoutes(db db.DB) router.Routes {
	return router.Routes{
		router.Route{Name: "Chat", Path: "/chat", Method: "GET", HandlerFunc: handlers.Chat(db)},
	}
}
