package server

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/handlers"

	"github.com/andreylm/basic-go-server.git/pkg/db/xorm"
	"github.com/andreylm/basic-go-server.git/pkg/jwt"
	"github.com/andreylm/basic-go-server.git/pkg/router"

	"github.com/andreylm/basic-go-server.git/pkg/server/config"
	"github.com/andreylm/basic-go-server.git/pkg/server/routes"

	"github.com/andreylm/basic-go-server.git/pkg/db"
)

// Server - server
type Server struct {
	port   string
	router *router.Router
	DB     db.DB
}

// NewServer - creates new server
func NewServer() Server {
	return Server{}
}

// Init - initialization
func (s *Server) Init(configurations *config.ServerConfigurations) (err error) {
	log.Println("Init server...")
	s.port = configurations.Port
	xormFactory := xorm.Factory{}
	if s.DB, err = xormFactory.GetDBDriverFactory(&configurations.DbConfig); err != nil {
		return
	}
	if err = jwt.Init(configurations.KeyPrivate, configurations.KeyPublic); err != nil {
		return
	}

	s.router = router.NewRouter()
	return
}

// Start - starting server
func (s *Server) Start() error {
	log.Println("Start server...")
	if err := s.DB.Connect(); err != nil {
		return err
	}
	defer s.DB.Close()

	s.router.AttachRoutes(routes.GetRoutes(s.DB))
	if err := s.initModules(); err != nil {
		return err
	}

	return s.run(s.createDefaultHandler())
}

func (s *Server) initModules() error {
	modules := config.GetModules()

	for _, module := range modules {
		if err := module.Init(s.router, s.DB); err != nil {
			return err
		}
	}

	// subRouter := s.router.Router.PathPrefix("/v2").Subrouter()
	// subRouter.Name("test").Methods("GET").Path("/test").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("response"))
	// })
	return nil
}

func (s *Server) run(handler http.Handler) error {
	newServer := &http.Server{
		Handler:      handler,
		Addr:         "0.0.0.0:" + s.port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	return newServer.ListenAndServe()
}

func (s *Server) createDefaultHandler() http.Handler {
	handler := handlers.LoggingHandler(os.Stdout, handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PATCH", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Origin", "Cache-Control", "X-App-Token"}),
		handlers.ExposedHeaders([]string{"*"}),
		handlers.MaxAge(1000),
		handlers.AllowCredentials(),
	)(s.router.Router))

	return handlers.RecoveryHandler(handlers.PrintRecoveryStack(true))(handler)
}
