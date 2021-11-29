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

func (r booksRepository) SaveAuthor(author *domain.Author) rest_errors.RestErr {
	stmt, err := r.db.Prepare(saveAuthorQuery)
	if err != nil {
		return rest_errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	inserResult, err := stmt.Exec(author.FirstName, author.LastName, author.Biography, author.Birthday, author.Death)
	if err != nil {
		return rest_errors.NewInternalServerError(err.Error())
	}

	authorId, _ := inserResult.LastInsertId()
	author.ID = authorId

	return nil
}

const (
	getAuthorById = `-- get author
	SELECT
		first_name,
		last_name,
		biography,
		birthday,
		death
	FROM authors
	WHERE authors.id = ?;	
	`

	getBooksFromAuthor = `-- get books from author
	SELECT
		books.id,
		books.title,
		books.published,
		books.short_description,
		books.original_release
	FROM authors
	INNER JOIN authorship
		ON authorship.author_id = authors.id
	INNER JOIN books
		ON authorship.book_id = books.id
	WHERE authors.id = ?;
	`
)

func (r booksRepository) GetAuthorById(authorID int64) (*domain.AuthorDenormalized, rest_errors.RestErr) {
	tx, err := r.db.Begin()
	if err != nil {
		return nil, rest_errors.NewInternalServerError(err.Error())
	}
	defer tx.Rollback()

	var author domain.AuthorDenormalized

	authorStmt, err := tx.Prepare(getAuthorById)
	if err != nil {
		return nil, rest_errors.NewInternalServerError(err.Error())
	}
	defer authorStmt.Close()

	if err := authorStmt.QueryRow(authorID).Scan(
		&author.Author.FirstName,
		&author.Author.LastName,
		&author.Author.Biography,
		&author.Author.Birthday,
		&author.Author.Death,
	); err != nil {
		return nil, rest_errors.NewInternalServerError(err.Error())
	}

	//

	booksStmt, err := tx.Prepare(getBooksFromAuthor)
	if err != nil {
		return nil, rest_errors.NewInternalServerError(err.Error())
	}
	defer booksStmt.Close()

	rows, err := booksStmt.Query(authorID)
	if err != nil {
		return nil, rest_errors.NewInternalServerError(err.Error())
	}

	var books domain.Book
	for rows.Next() {
		if err := rows.Scan(
			&books.ID,
			&books.Title,
			&books.Published,
			&books.ShortDescription,
			&books.OriginalRelease,
		); err != nil {
			return nil, rest_errors.NewInternalServerError(err.Error())
		}
		author.Books = append(author.Books, books)
	}
	rows.Close()

	if err := tx.Commit(); err != nil {
		return nil, rest_errors.NewInternalServerError(err.Error())
	}
	return &author, nil
}

func (r booksRepository) UpdateAuthor(*domain.Author) rest_errors.RestErr {
	return nil
}

const savePublisherQuery = `-- save publisher
INSERT INTO publishers(
	name,
	description,
	slogan,
	founded
) VALUES (
	?, ?, ?, ?
);
`

func (r booksRepository) SavePublisher(publisher *domain.Publisher) rest_errors.RestErr {
	stmt, err := r.db.Prepare(savePublisherQuery)
	if err != nil {
		return rest_errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	inserResult, err := stmt.Exec(publisher.Name, publisher.Description, publisher.Slogan, publisher.Founded)
	if err != nil {
		return rest_errors.NewInternalServerError(err.Error())
	}

	publisherId, _ := inserResult.LastInsertId()
	publisher.ID = publisherId

	return nil
}

func (r booksRepository) UpdatePublisher(*domain.Publisher) rest_errors.RestErr {
	return nil
}

func (r booksRepository) GetPublisherById(publisherID int64) (*domain.PublisherDenormalized, rest_errors.RestErr) {
	return nil, nil
}

const (
	saveBookQuery = `-- save book
	INSERT INTO books(
		title,
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
	INSERT IGNORE INTO published(
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
		book.Title,
		book.OriginalRelease,
		book.Description,
		book.ShortDescription,
		book.Published,
		book.PublisherID,
		book.Pages,
		book.SellerID,
	)
	if err != nil {
		return rest_errors.NewInternalServerError(err.Error())
	}
	defer bookStmt.Close()

	bookId, _ := inserResult.LastInsertId()
	book.ID = bookId

	// TODO: would a better implementation of this use go routines?
	for k := range book.AuthorID {
		authorShipStmt, err := tx.Prepare(saveAuthorshipQuery)
		if err != nil {
			return rest_errors.NewInternalServerError(err.Error())
		}
		defer authorShipStmt.Close()

		if _, err = authorShipStmt.Exec(bookId, book.AuthorID[k]); err != nil {
			return rest_errors.NewInternalServerError(err.Error())
		}

		//

		publishedStmt, err := tx.Prepare(savePublishedQuery)
		if err != nil {
			return rest_errors.NewInternalServerError(err.Error())
		}
		defer publishedStmt.Close()

		if _, err = publishedStmt.Exec(book.AuthorID[k], book.PublisherID); err != nil {
			return rest_errors.NewInternalServerError(err.Error())
		}
	}

	tx.Commit()
	return nil
}

const (
	getBookById = ` -- get book
	SELECT 
		books.title,
		books.original_release,  
		books.description,      
		books.short_description, 
		books.published,        
		books.pages,            
		books.seller_id,
		publishers.id,
		publishers.name
	FROM books
	INNER JOIN publishers
	ON publishers.id = books.publisher_id
	WHERE books.id = ?;         
	`

	getAuthorsForBook = ` -- get authors for book
	SELECT 
		authors.id,
		authors.first_name,
		authors.last_name
	FROM authors
	INNER JOIN authorship
		ON authors.id = authorship.author_id
	WHERE authorship.book_id = ?;
	`
)

func (r booksRepository) GetBookById(bookID int64) (*domain.BookDenormalized, rest_errors.RestErr) {
	tx, err := r.db.Begin()
	if err != nil {
		return nil, rest_errors.NewInternalServerError(err.Error())
	}

	bookStmt, err := tx.Prepare(getBookById)
	if err != nil {
		return nil, rest_errors.NewInternalServerError(err.Error())
	}
	defer bookStmt.Close()

	var book domain.BookDenormalized

	if err := bookStmt.QueryRow(bookID).Scan(
		&book.Book.Title,
		&book.Book.OriginalRelease,
		&book.Book.Description,
		&book.Book.ShortDescription,
		&book.Book.Published,
		&book.Book.Pages,
		&book.Book.SellerID,
		&book.Publisher.ID,
		&book.Publisher.Name,
	); err != nil {
		return nil, rest_errors.NewInternalServerError(err.Error())
	}

	//

	authorsStmt, err := r.db.Prepare(getAuthorsForBook)
	if err != nil {
		return nil, rest_errors.NewInternalServerError(err.Error())
	}
	defer authorsStmt.Close()

	rows, err := authorsStmt.Query(bookID)
	if err != nil {
		return nil, rest_errors.NewInternalServerError(err.Error())
	}

	var author domain.Author

	for rows.Next() {
		if err := rows.Scan(
			&author.ID,
			&author.FirstName,
			&author.LastName,
		); err != nil {
			return nil, rest_errors.NewInternalServerError(err.Error())
		}

		book.Authors = append(book.Authors, author)
	}
	rows.Close()

	return &book, nil
}

func (r booksRepository) UpdateBook(book *domain.Book) rest_errors.RestErr {
	return nil
}
