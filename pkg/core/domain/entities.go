package domain

type Author struct {
	ID        int64   `json:"id,omitempty"`
	FirstName string  `json:"first_name,omitempty"`
	LastName  string  `json:"last_name,omitempty"`
	Biography string  `json:"biography,omitempty"`
	Birthday  string  `json:"birthday,omitempty"`
	Death     *string `json:"death,omitempty"`
}

type Book struct {
	ID               int64   `json:"id,omitempty"`
	Title            string  `json:"title,omitempty"`
	OriginalRelease  string  `json:"original_release,omitempty"`
	Description      string  `json:"description,omitempty"`
	ShortDescription string  `json:"short_description,omitempty"`
	Published        string  `json:"published,omitempty"`
	PublisherID      int64   `json:"publisher_id,omitempty"`
	Pages            int64   `json:"pages,omitempty"`
	AuthorID         []int64 `json:"author_id,omitempty"`
	SellerID         int64   `json:"seller_id,omitempty"`
}

type Publisher struct {
	ID          int64  `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Slogan      string `json:"slogan,omitempty"`
	Founded     string `json:"founded,omitempty"`
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
