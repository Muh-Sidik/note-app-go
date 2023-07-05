package src

type CreateNoteDto struct {
	Title string   `json:"title"`
	Tags  []string `json:"tags,omitempty"`
	Body  string   `json:"body,omitempty"`
}

type UpdateNoteDto struct {
	CreateNoteDto
}
