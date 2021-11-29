package repositories

import (
	"database/sql"
	"log"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/FacuBar/bookstore_books-api/pkg/core/domain"
	"github.com/stretchr/testify/assert"
)

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}
func TestSaveAuthor(t *testing.T) {
	query := regexp.QuoteMeta(saveAuthorQuery)

	testAuthor := domain.Author{
		FirstName: "Philip",
		LastName:  "Dick",
		Biography: "a weird biography ...",
		Birthday:  "16-12-1928",
		Death:     "2-3-1982",
	}

	t.Run("NoError", func(t *testing.T) {
		db, mock := NewMock()

		author := testAuthor
		repo := booksRepository{db: db}
		mock.ExpectPrepare(query).ExpectExec().WithArgs(
			author.FirstName,
			author.LastName,
			author.Biography,
			author.Birthday,
			author.Death,
		).WillReturnResult(sqlmock.NewResult(1234, 1))

		err := repo.SaveAuthor(&author)

		assert.Nil(t, err)
		assert.EqualValues(t, 1234, author.ID)
	})
}

func TestSavePublisher(t *testing.T) {
	query := regexp.QuoteMeta(savePublisherQuery)

	testPublisher := domain.Publisher{
		Name:        "Penguin",
		Description: "publisher with the penguin mascot :D",
		Slogan:      "some slogan",
		Founded:     "12-12-2000",
	}

	t.Run("NoError", func(t *testing.T) {
		db, mock := NewMock()

		repo := booksRepository{db: db}
		publisher := testPublisher

		mock.ExpectPrepare(query).ExpectExec().WithArgs(
			publisher.Name,
			publisher.Description,
			publisher.Slogan,
			publisher.Founded,
		).WillReturnResult(sqlmock.NewResult(12, 1))

		err := repo.SavePublisher(&publisher)
		assert.Nil(t, err)
		assert.EqualValues(t, 12, publisher.ID)
	})
}

func TestSaveBook(t *testing.T) {
	queryBook := regexp.QuoteMeta(saveBookQuery)
	queryAuthorShip := regexp.QuoteMeta(saveAuthorshipQuery)
	queryPublished := regexp.QuoteMeta(savePublishedQuery)

	testBook := domain.Book{
		Name:             "Flow my tears, the policeman said",
		OriginalRelease:  "01-01-1970",
		Description:      "some description",
		ShortDescription: "sm descrpt",
		Publised:         "20-12-2021",
		PublisherID:      12,
		Pages:            256,
		AuthorID:         []int64{0, 1},
		SellerID:         1,
	}

	t.Run("NoError", func(t *testing.T) {
		db, mock := NewMock()

		repo := booksRepository{db: db}
		book := testBook

		mock.ExpectBegin()
		mock.ExpectPrepare(queryBook).ExpectExec().WithArgs(
			book.Name,
			book.OriginalRelease,
			book.Description,
			book.ShortDescription,
			book.Publised,
			book.PublisherID,
			book.Pages,
			book.SellerID,
		).WillReturnResult(sqlmock.NewResult(69, 1))

		for k := range book.AuthorID {
			mock.ExpectPrepare(queryAuthorShip).ExpectExec().WithArgs(
				69, // Id of the inserted book
				book.AuthorID[k],
			).WillReturnResult(sqlmock.NewResult(1, 1))

			mock.ExpectPrepare(queryPublished).ExpectExec().WithArgs(
				book.AuthorID[k],
				book.PublisherID,
			).WillReturnResult(sqlmock.NewResult(1, 1))
		}

		mock.ExpectCommit()

		err := repo.SaveBook(&book)
		assert.Nil(t, err)
		assert.EqualValues(t, 69, book.ID)
	})
}
