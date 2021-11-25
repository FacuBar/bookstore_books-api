package domain

// Tables from which data will be retrieved to
// denormalize info into 'queries tables' -BookBy...-

type Author struct {
	Id             uint64 //pk
	Name           string
	Biography      string
	ShortBiography string
	// ...
}

type Publisher struct {
	Id          uint64 //pk
	Name        string
	Description string
	// ...
}

type BookById struct {
	// Book info that is independent of publication
	BookId           uint64 // pk
	Name             string
	Description      string
	ShortDescription string
	ReleaseDate      string
	// Publication related info
	PublishedDate string
	Pages         uint16

	// Denormalized data from others tables

	// Auhtor related info
	AuthorName string
	AuthorId   uint64

	// Publisher related info
	PublisherId   int64
	PublisherName string

	// Seller info id
	SellerId int64
}

type BookByAuthor struct {
	AuthorId uint64 //pk
	Name     string

	BookId           uint64
	BookName         string
	ShortDescription string

	PubliserName  string
	PublishedDate string //ck
}

type BookByPublisher struct {
	PublisherId uint64 //pk
	Name        string

	BookId           uint64
	BookName         string
	ShortDescription string

	AuthorName string
	// BookSells int64	//ck
}
