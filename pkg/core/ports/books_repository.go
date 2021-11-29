package ports

import (
	"github.com/FacuBar/bookstore_books-api/pkg/core/domain"
	"github.com/FacuBar/bookstore_utils-go/rest_errors"
)

type BooksRepositoryInterface interface {
	SaveAuthor(*domain.Author) rest_errors.RestErr
	UpdateAuthor(*domain.Author) rest_errors.RestErr
	GetAuthorById(int64) (*domain.AuthorDenormalized, rest_errors.RestErr)

	SavePublisher(*domain.Publisher) rest_errors.RestErr
	UpdatePublisher(*domain.Publisher) rest_errors.RestErr
	GetPublisherById(int64) (*domain.PublisherDenormalized, rest_errors.RestErr)

	SaveBook(*domain.Book) rest_errors.RestErr
	UpdateBook(*domain.Book) rest_errors.RestErr
	GetBookById(int64) (*domain.BookDenormalized, rest_errors.RestErr)
}
