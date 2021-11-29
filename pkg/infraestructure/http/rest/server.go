package rest

import (
	"database/sql"

	"github.com/FacuBar/bookstore_books-api/pkg/infraestructure/repositories"
	"github.com/gin-gonic/gin"
)

type Server struct {
	db     *sql.DB
	router *gin.Engine
}

func NewServer(db *sql.DB) *Server {
	server := &Server{
		db: db,
	}

	bookrepo := repositories.NewBooksRepo(db)

	router := server.handler(bookrepo)

	server.router = router
	return server
}

func (s *Server) Start(address string) {
	s.router.Run(address)
}
