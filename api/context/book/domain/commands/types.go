package commands

// UpdateBookCommand ..
type UpdateBookCommand struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Author    string `json:"author"`
	Publisher string `json:"publisher"`
}

// CreateBookCommand ..
type CreateBookCommand struct {
	Title     string `json:"title"`
	Author    string `json:"author"`
	Publisher string `json:"publisher"`
}
