package backend

import (
	"log"
	"net/http"
)

// ApiServer ...
type ApiServer struct {
	config    *Config
	serverMux *http.ServeMux
}

// New ...
func New(config *Config) *ApiServer {
	return &ApiServer{config: config, serverMux: http.NewServeMux()}
}

// Start ...
func (server *ApiServer) Start() error {
	log.Println("Starting API server...")

	server.serverMux.HandleFunc("/hello", HelloHandler)

	return http.ListenAndServe(server.config.BindAddr, server.serverMux)
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Hello World!"))
	if err != nil {
		log.Fatal(err)
	}
}
