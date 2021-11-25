package domain

type BookById struct {
	// Book info that is independent of publication
	Id               uint64 // pk
	Name             string
	Description      string
	ShortDescription string
	ReleaseDate      string

	AuthorName string
	AuthorId   uint64

	// Publication related info
	PublisherId   int64
	PublisherName string
	PublishedDate string
	Pages         uint16

	SellerId int64
}

type BookByAuthor struct {
	Id        int64 //pk
	Name      string
	Biography string

	BookId           uint64
	BookName         string
	ShortDescription string

	PubliserName  string
	PublishedDate string //ck
}

type BookByPublisher struct {
	Id   int64 //pk
	Name string

	BookId           uint64
	BookName         string
	ShortDescription string

	AuthorName string
	// BookSells int64	//ck
}
