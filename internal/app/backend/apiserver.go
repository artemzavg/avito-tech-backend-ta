package backend

import (
	"github.com/artemzavg/avito-tech-backend-ta/internal/app/backend/dbContext"
	"github.com/artemzavg/avito-tech-backend-ta/internal/app/backend/handlers/createAdvert"
	"gorm.io/gorm"
	"log"
	"net/http"
)

type ApiServer struct {
	config    *Config
	serverMux *http.ServeMux
	db        *gorm.DB
}

func New(config *Config) *ApiServer {
	return &ApiServer{
		config:    config,
		serverMux: http.NewServeMux(),
		db:        dbContext.NewContext(config.DbConnectionString),
	}
}

func (server *ApiServer) Start() error {
	log.Println("Starting API server...")

	server.serverMux.HandleFunc("/create", createAdvert.Handler{Db: server.db}.HandlerFunc)

	return http.ListenAndServe(server.config.BindAddr, server.serverMux)
}
