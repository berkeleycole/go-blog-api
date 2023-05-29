package http

import (
	"blog-api/writings"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// Returns an instance of the Server struct
func New(writings_service writings.Service) Server {
	fmt.Printf("Init server")
	return Server{writings_service}
}

// Server struct
type Server struct {
	writings_service writings.Service
}

// Adds a method Server called Start()
func (s *Server) Start() error {
	router := mux.NewRouter()

	Routes(router, &s.writings_service)

	server := &http.Server{
		Handler:      router,
		Addr:         "0.0.0.0:80",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	return server.ListenAndServe()
}
