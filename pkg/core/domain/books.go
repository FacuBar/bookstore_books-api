package domain

type BookDenormalized struct {
	Book      *Book      `json:"book"`
	Authors   []Author   `json:"authors"`
	Publisher *Publisher `json:"publisher"`
}

type PublisherDenormalized struct {
	Publisher *Publisher `json:"publisher"`
	Authors   []Author   `json:"authors"`
	Books     []Book     `json:"books"`
}

type AuthorDenormalized struct {
	Author     *Author     `json:"author"`
	Books      []Book      `json:"books"`
	Publishers []Publisher `json:"publishers"`
}

type Author struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Biography string `json:"biography"`
	Birthday  string `json:"birthday"`
	Death     string `json:"death"`
}

type Authorship struct {
	BookID   int64 `json:"book_id"`
	AuthorID int64 `json:"author_id"`
}

type Book struct {
	ID               int64   `json:"id"`
	Name             string  `json:"name"`
	OriginalRelease  string  `json:"original_release"`
	Description      string  `json:"description"`
	ShortDescription string  `json:"short_description"`
	Publised         string  `json:"publised"`
	PublisherID      int64   `json:"publisher_id"`
	Pages            int64   `json:"pages"`
	AuthorID         []int64 `json:"author_id"`
	SellerID         int64   `json:"seller_id"`
}

type Published struct {
	AuthorID    int64 `json:"author_id"`
	PublisherID int64 `json:"publisher_id"`
}

type Publisher struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Slogan      string `json:"slogan"`
	Founded     string `json:"founded"`
}

/*
	-- Although Cassandra seems like a good choice for this app
	it will not be used in the time being --
*/

// // Tables from which data will be retrieved to
// // denormalize info into 'queries tables' -BookBy...-

// type Author struct {
// 	Id             uint64 //pk
// 	Name           string
// 	Biography      string
// 	ShortBiography string
// 	// ...
// }

// type Publisher struct {
// 	Id          uint64 //pk
// 	Name        string
// 	Description string
// 	// ...
// }

// type BookById struct {
// 	// Book info that is independent of publication
// 	BookId           uint64 // pk
// 	Name             string
// 	Description      string
// 	ShortDescription string
// 	ReleaseDate      string
// 	// Publication related info
// 	PublishedDate string
// 	Pages         uint16

// 	// Denormalized data from others tables

// 	// Auhtor related info
// 	AuthorName string
// 	AuthorId   uint64

// 	// Publisher related info
// 	PublisherId   int64
// 	PublisherName string

// 	// Seller info id
// 	SellerId int64
// }

// type BookByAuthor struct {
// 	AuthorId uint64 //pk
// 	Name     string

// 	BookId           uint64
// 	BookName         string
// 	ShortDescription string

// 	PubliserName  string
// 	PublishedDate string //ck
// }

// type BookByPublisher struct {
// 	PublisherId uint64 //pk
// 	Name        string

// 	BookId           uint64
// 	BookName         string
// 	ShortDescription string

// 	AuthorName string
// 	// BookSells int64	//ck
// }
