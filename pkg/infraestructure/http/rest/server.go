package rest

import (
	"context"
	"database/sql"
	"log"
	"net/http"

	"github.com/FacuBar/bookstore_books-api/pkg/infraestructure/repositories"
)

type Server struct {
	db  *sql.DB
	srv *http.Server
}

func NewServer(srv *http.Server, db *sql.DB) *Server {
	server := &Server{
		db:  db,
		srv: srv,
	}

	bookrepo := repositories.NewBooksRepo(db)

	router := server.handler(bookrepo)

	server.srv.Handler = router
	return server
}

func (s *Server) Start() {
	if err := s.srv.ListenAndServe(); err != nil {
		log.Fatalf("error while serving: %v", err)
	}
}

func (s *Server) Stop(ctx context.Context) {
	s.db.Close()

	go func() {
		if err := s.srv.Shutdown(ctx); err != nil {
			log.Fatal("Server Shutdown:", err)
		}
	}()

	log.Println("Server exiting")
}
