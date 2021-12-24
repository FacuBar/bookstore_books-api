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

var (
	testBook = domain.Book{
		Title:            "Flow my tears, the policeman said",
		OriginalRelease:  "01-01-1970",
		Description:      "some description",
		ShortDescription: "sm descrpt",
		Published:        "20-12-2021",
		PublisherID:      12,
		Pages:            256,
		AuthorID:         []int64{0, 1},
		SellerID:         1,
	}
)

func TestSaveAuthor(t *testing.T) {
	query := regexp.QuoteMeta(saveAuthorQuery)

	death := "1982-12-1"
	testAuthor := domain.Author{
		FirstName: "Philip",
		LastName:  "Dick",
		Biography: "a weird biography ...",
		Birthday:  "16-12-1928",
		Death:     &death,
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

	t.Run("NoError", func(t *testing.T) {
		db, mock := NewMock()

		repo := booksRepository{db: db}
		book := testBook

		mock.ExpectBegin()
		mock.ExpectPrepare(queryBook).ExpectExec().WithArgs(
			book.Title,
			book.OriginalRelease,
			book.Description,
			book.ShortDescription,
			book.Published,
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

func TestGetBookByID(t *testing.T) {
	queryBook := regexp.QuoteMeta(getBookById)
	queryAuthors := regexp.QuoteMeta(getAuthorsForBook)

	t.Run("NoError", func(t *testing.T) {
		bookRow := sqlmock.NewRows([]string{
			"books.title",
			"books.original_release",
			"books.description",
			"books.short_description",
			"books.published",
			"books.pages",
			"books.seller_id",
			"publishers.id",
			"publishers.name",
		}).
			AddRow(
				testBook.Title,
				testBook.OriginalRelease,
				testBook.Description,
				testBook.ShortDescription,
				testBook.Published,
				testBook.Pages,
				testBook.SellerID,
				testBook.PublisherID,
				"penguin",
			)

		authorRows := sqlmock.NewRows([]string{
			"authors.id",
			"authors.first_name",
			"authors.last_name",
		}).AddRow(
			0,
			"Philip",
			"Dick",
		).AddRow(
			1,
			"Jorge Luis",
			"Borges",
		)

		db, mock := NewMock()
		repo := booksRepository{db: db}

		bookID := 1
		mock.ExpectBegin()
		mock.ExpectPrepare(queryBook).ExpectQuery().WithArgs(bookID).WillReturnRows(bookRow)
		mock.ExpectPrepare(queryAuthors).ExpectQuery().WithArgs(bookID).WillReturnRows(authorRows)
		mock.ExpectCommit()

		_, err := repo.GetBookById(int64(bookID))
		assert.Nil(t, err)
	})
}
