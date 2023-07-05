package src

type Note struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Tags      []string  `json:"tags"`
	Body      string    `json:"body"`
	CreatedAt string `json:"createdAt"`
	UpdateAt  string `json:"updatedAt"`
}
