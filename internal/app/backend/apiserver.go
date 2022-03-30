package backend

import (
	"github.com/artemzavg/avito-tech-backend-ta/internal/app/backend/dbContext"
	"github.com/artemzavg/avito-tech-backend-ta/internal/app/backend/handlers/createAdvert"
	"github.com/artemzavg/avito-tech-backend-ta/internal/app/backend/handlers/getAdvert"
	"github.com/artemzavg/avito-tech-backend-ta/internal/app/backend/handlers/getAllAdverts"
	"gorm.io/gorm"
	"log"
	"net/http"
)

type Server struct {
	config    *Config
	serverMux *http.ServeMux
	db        *gorm.DB
}

func NewServer(config *Config) *Server {
	return &Server{
		config:    config,
		serverMux: http.NewServeMux(),
		db:        dbContext.NewContext(config.DbConnectionString),
	}
}

func (server *Server) Start() error {
	log.Println("Starting API server...")

	server.serverMux.HandleFunc("/create", createAdvert.Handler{Db: server.db}.HandlerFunc)
	server.serverMux.HandleFunc("/get", getAdvert.Handler{Db: server.db}.HandlerFunc)
	server.serverMux.HandleFunc("/getAll", getAllAdverts.Handler{Db: server.db}.HandlerFunc)

	return http.ListenAndServe(server.config.BindAddr, server.serverMux)
}
