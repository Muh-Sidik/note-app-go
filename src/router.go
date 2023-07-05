package src

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Router func(r fiber.Router)

func RouteHandler(db *gorm.DB) Router {

	note := &NoteController{db}

	return func(r fiber.Router) {
		r.Get("", note.GetAllNotes)
		r.Post("", note.CreateNote)
		r.Get("/:id", note.GetNoteById)
		r.Put("/:id", note.UpdateNote)
		r.Delete("/:id", note.DeleteNote)
	}
}
