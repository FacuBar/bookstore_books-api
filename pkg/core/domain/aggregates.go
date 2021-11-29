package domain

type BookDenormalized struct {
	Book      Book      `json:"book"`
	Authors   []Author  `json:"authors"`
	Publisher Publisher `json:"publisher"`
}

type PublisherDenormalized struct {
	Publisher Publisher `json:"publisher"`
	Authors   []Author  `json:"authors"`
	Books     []Book    `json:"books"`
}

type AuthorDenormalized struct {
	Author Author `json:"author"`
	Books  []Book `json:"books"`
}

type Published struct {
	AuthorID    int64 `json:"author_id"`
	PublisherID int64 `json:"publisher_id"`
}

type Authorship struct {
	BookID   int64 `json:"book_id"`
	AuthorID int64 `json:"author_id"`
}
