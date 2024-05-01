package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/francopoffo/fullstack-auth/service/user"
	"github.com/gorilla/mux"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api").Subrouter()

	userHandler := user.NewHandler()

	subrouter.HandleFunc("/login", userHandler.Login)
	subrouter.HandleFunc("/register", userHandler.Register)

	log.Printf("Listening and serving HTTP on %s\n", s.addr)

	return http.ListenAndServe(s.addr, router)
}
