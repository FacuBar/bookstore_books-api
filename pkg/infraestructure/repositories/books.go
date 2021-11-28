package repositories

import (
	"context"
	"database/sql"
	"sync"

	"github.com/FacuBar/bookstore_books-api/pkg/core/domain"
	"github.com/FacuBar/bookstore_books-api/pkg/core/ports"
	"github.com/FacuBar/bookstore_utils-go/rest_errors"
)

type booksRepository struct {
	db *sql.DB
}

var (
	onceInstanceBooks sync.Once
	instanceBooks     booksRepository
)

func NewBooksRepo(db *sql.DB) ports.BooksRepositoryInterface {
	onceInstanceBooks.Do(func() {
		instanceBooks = booksRepository{db: db}
	})
	return instanceBooks
}

const saveAuthorQuery = `-- save author
INSERT INTO authors(
	first_name,
	last_name,
	biography,
	birthday,
	death
) VALUES (
	?, ?, ?, ?, ?
);
`

func (r booksRepository) SaveAuthor(*domain.Author) rest_errors.RestErr {
	stmt, err := r.db.Prepare(saveAuthorQuery)
	if err != nil {
		return rest_errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	return nil
}

func (r booksRepository) UpdateAuthor(*domain.Author) rest_errors.RestErr
func (r booksRepository) GetAuthorById(uint32) (domain.AuthorDenormalized, rest_errors.RestErr)

const savePublisherQuery = `-- save author
INSERT INTO authors(
	name,
	description,
	slogan,
	founded
) VALUES (
	?, ?, ?, ?
);
`

func (r booksRepository) SavePublisher(*domain.Publisher) rest_errors.RestErr {
	stmt, err := r.db.Prepare(savePublisherQuery)
	if err != nil {
		return rest_errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	return nil
}

func (r booksRepository) UpdatePublisher(*domain.Publisher) rest_errors.RestErr
func (r booksRepository) GetPublisherById(uint32) (domain.PublisherDenormalized, rest_errors.RestErr)

const (
	saveBookQuery = `-- save book
	INSERT INTO authors(
		name,
		original_release,
		description,
		short_description,
		published,
		publisher_id,
		pages,
		seller_id
	) VALUES (
		?, ?, ?, ?, ?, ?, ?, ?
	);
	`

	saveAuthorshipQuery = `-- save authorship
	INSERT INTO authorship(
		book_id,
		author_id
	) VALUES (
		?, ?
	);
	`

	savePublishedQuery = `-- save published
	INSERT INTO published(
		author_id,
		publisher_id
	) VALUES (
		?, ?
	)
	`
)

func (r booksRepository) SaveBook(book *domain.Book) rest_errors.RestErr {
	tx, err := r.db.BeginTx(context.Background(), nil)
	if err != nil {
		return rest_errors.NewInternalServerError(err.Error())
	}

	defer func() {
		_ = tx.Rollback()
	}()

	bookStmt, err := tx.Prepare(saveBookQuery)
	if err != nil {
		return rest_errors.NewInternalServerError(err.Error())
	}

	inserResult, err := bookStmt.Exec(
		book.Name,
		book.OriginalRelease,
		book.Description,
		book.ShortDescription,
		book.Publised,
		book.PublisherID,
		book.Pages,
		book.SellerID,
	)
	if err != nil {
		return rest_errors.NewInternalServerError(err.Error())
	}
	defer bookStmt.Close()

	bookId, _ := inserResult.LastInsertId()

	authorShipStmt, err := tx.Prepare(saveAuthorshipQuery)
	if err != nil {
		return rest_errors.NewInternalServerError(err.Error())
	}
	defer authorShipStmt.Close()

	if _, err = authorShipStmt.Exec(bookId, book.AuthorID); err != nil {
		return rest_errors.NewInternalServerError(err.Error())
	}

	publishedStmt, err := tx.Prepare(savePublishedQuery)
	if err != nil {
		return rest_errors.NewInternalServerError(err.Error())
	}
	defer publishedStmt.Close()

	if _, err = publishedStmt.Exec(book.AuthorID, book.PublisherID); err != nil {
		return rest_errors.NewInternalServerError(err.Error())
	}

	tx.Commit()
	return nil
}
func (r booksRepository) UpdateBook(*domain.Book)
func (r booksRepository) GetBookById(uint32) (domain.BookDenormalized, rest_errors.RestErr)
