package rest

import (
	"net/http"
	"strconv"

	"github.com/FacuBar/bookstore_books-api/pkg/core/domain"
	"github.com/FacuBar/bookstore_books-api/pkg/core/ports"
	"github.com/FacuBar/bookstore_utils-go/rest_errors"
	"github.com/gin-gonic/gin"
)

func (s *Server) handler(br ports.BooksRepositoryInterface) *gin.Engine {
	router := gin.Default()

	router.GET("/authors/:author_id", getAuthor(br))

	router.POST("/authors", createAuthor(br))
	router.POST("/publishers", createPublisher(br))
	router.POST("/books", createBook(br))

	return router
}

func createAuthor(br ports.BooksRepositoryInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		var author domain.Author
		if err := c.ShouldBindJSON(&author); err != nil {
			restErr := rest_errors.NewBadRequestError("invalid request")
			c.JSON(restErr.Status(), restErr)
			return
		}

		if err := br.SaveAuthor(&author); err != nil {
			c.JSON(err.Status(), err)
			return
		}
		c.JSON(http.StatusOK, author)
	}
}

func createPublisher(br ports.BooksRepositoryInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		var publisher domain.Publisher
		if err := c.ShouldBindJSON(&publisher); err != nil {
			restErr := rest_errors.NewBadRequestError("invalid request")
			c.JSON(restErr.Status(), restErr)
			return
		}

		if err := br.SavePublisher(&publisher); err != nil {
			c.JSON(err.Status(), err)
			return
		}
		c.JSON(http.StatusOK, publisher)
	}
}

func createBook(br ports.BooksRepositoryInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		var book domain.Book
		if err := c.ShouldBindJSON(&book); err != nil {
			restErr := rest_errors.NewBadRequestError("invalid request")
			c.JSON(restErr.Status(), restErr)
			return
		}

		if err := br.SaveBook(&book); err != nil {
			c.JSON(err.Status(), err)
			return
		}
		c.JSON(http.StatusOK, book)
	}
}

func getAuthor(br ports.BooksRepositoryInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		authorID, idErr := strconv.ParseInt(c.Param("author_id"), 10, 64)
		if idErr != nil {
			restErr := rest_errors.NewBadRequestError("invalid author id")
			c.JSON(restErr.Status(), restErr)
			return
		}

		author, err := br.GetAuthorById(authorID)
		if err != nil {
			c.JSON(err.Status(), err)
			return
		}

		c.JSON(http.StatusOK, author)
	}
}
