package src

import (
	"time"

	"github.com/gofiber/fiber/v2"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type NoteController struct {
	DB *gorm.DB
}

func (n *NoteController) CreateNote(c *fiber.Ctx) error {
	var body CreateNoteDto

	err := c.BodyParser(&body)

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	id, err := gonanoid.New(16)

	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	err = n.DB.Create(&Note{
		ID:        id,
		Title:     body.Body,
		Tags:      body.Tags,
		Body:      body.Body,
		CreatedAt: time.Now().String(),
		UpdateAt:  time.Now().String(),
	}).Error

	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "Note created",
		"data":    fiber.Map{"notedId": id},
	})
}

func (n *NoteController) GetAllNotes(c *fiber.Ctx) error {
	var notes []Note

	err := n.DB.Find(&notes).Error

	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "Get All Notes",
		"data":    notes,
	})
}

func (n *NoteController) GetNoteById(c *fiber.Ctx) error {
	id := c.Params("id")

	var note Note

	err := n.DB.Take(&note, "id = ?", id).Error

	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "Get Note",
		"data":    note,
	})
}

func (n *NoteController) UpdateNote(c *fiber.Ctx) error {
	id := c.Params("id")
	var body UpdateNoteDto

	err := c.BodyParser(&body)

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	err = n.DB.Where("id = ?", id).Updates(&Note{
		Title:    body.Body,
		Tags:     body.Tags,
		Body:     body.Body,
		UpdateAt: time.Now().String(),
	}).Error

	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "Note updated",
		"data":    fiber.Map{"id": id},
	})
}

func (n *NoteController) DeleteNote(c *fiber.Ctx) error {
	id := c.Params("id")

	err := n.DB.Delete(&Note{}, "id = ?", id).Error

	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "Delete Note",
	})
}
