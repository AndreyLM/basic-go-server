package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Routes - slice of routes
type Routes []Route

// Route - base route struct
type Route struct {
	Name        string
	Path        string
	Method      string
	HandlerFunc http.HandlerFunc
}

// SubRouter - sub router
type SubRouter struct {
	PathPrefix string
	Routes     Routes
	Middleware func(next http.Handler) http.Handler
}

// NewRouter - returns new router
func NewRouter() *Router {
	return &Router{
		Router: mux.NewRouter(),
	}
}

// Router - router
type Router struct {
	Router *mux.Router
}

// AttachRoutes - attaches root routes
func (r *Router) AttachRoutes(rootRoutes Routes) {
	for _, route := range rootRoutes {
		r.Router.Name(route.Name).Path(route.Path).Methods(route.Method).HandlerFunc(route.HandlerFunc)
	}
}

// AttachSubRouterWithMiddleware - attaches subroutes to main router
func (r *Router) AttachSubRouterWithMiddleware(subRouter *SubRouter) {
	MuxSubRouter := r.Router.PathPrefix(subRouter.PathPrefix).Subrouter()

	if subRouter.Middleware != nil {
		MuxSubRouter.Use(subRouter.Middleware)
	}

	for _, route := range subRouter.Routes {
		MuxSubRouter.Name(route.Name).Methods(route.Method).Path(route.Path).HandlerFunc(route.HandlerFunc)
	}

	return
}
