package http

import (
	"strconv"

	"github.com/Enthreeka/notes-server/internal/apperror"
	"github.com/Enthreeka/notes-server/internal/entity"
	"github.com/Enthreeka/notes-server/internal/usecase"
	"github.com/Enthreeka/notes-server/pkg/logger"
	"github.com/gofiber/fiber/v2"
)

type noteHandler struct {
	usecase usecase.NoteService
	log     *logger.Logger
}

func NewNoteHandler(usecase usecase.NoteService, log *logger.Logger) *noteHandler {
	return &noteHandler{
		usecase: usecase,
		log:     log,
	}
}

func (n *noteHandler) SearchNode(c *fiber.Ctx) error {
	n.log.Info("Start of note search")

	idInt, err := n.convertIdtoInt(c)
	if err != nil {
		return err
	}

	notes, err := n.usecase.GetNotes(c.Context(), idInt)
	if err != nil {
		if err == apperror.ErrNoteNotFound {
			return c.Status(fiber.StatusNotFound).JSON(apperror.ErrNoteNotFound)
		}
		return c.Status(fiber.StatusInternalServerError).JSON(map[string]interface{}{
			"Message": "Internal server error",
		})
	}

	n.log.Info("Search note completed successfully")
	return c.JSON(notes)
}

func (n *noteHandler) CreateNote(c *fiber.Ctx) error {
	n.log.Info("Start of note create")

	var notes entity.Notes
	err := c.BodyParser(&notes)
	if err != nil {
		n.log.Error("Failed to parse request body: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{
			"message": "Invalid request body",
		})
	}

	err = n.usecase.CreateNotes(c.Context(), notes)
	if err != nil {
		n.log.Error("Failed to create notes: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(apperror.ErrCreateNotes)
	}

	n.log.Info("Create note completed successfully")
	return c.Status(fiber.StatusCreated).JSON(map[string]interface{}{
		"note": notes.Notes,
	})
}

func (n *noteHandler) DeleteNotes(c *fiber.Ctx) error {
	n.log.Info("Start of note delete")

	idInt, err := n.convertIdtoInt(c)
	if err != nil {
		return err
	}

	notes, err := n.usecase.GetNotes(c.Context(), idInt)
	if err != nil {
		if err == apperror.ErrNoteNotFound {
			return c.Status(fiber.StatusNotFound).JSON(apperror.ErrNoteNotFound)
		}
		return c.Status(fiber.StatusInternalServerError).JSON(map[string]interface{}{
			"Message": "Internal server error",
		})
	}

	err = n.usecase.DeleteNotes(c.Context(), idInt)
	if err != nil {
		n.log.Error("Failed to delete notes: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(apperror.ErrDeleteNotes)
	}

	n.log.Info("Delete note completed successfully")
	return c.Status(fiber.StatusOK).JSON(map[string]interface{}{
		"deleted_id":   notes.ID,
		"deleted_note": notes.Notes,
	})
}

func (n *noteHandler) UpdateNote(c *fiber.Ctx) error {
	n.log.Info("Start of note update")

	idInt, err := n.convertIdtoInt(c)
	if err != nil {
		return err
	}

	var notes entity.Notes
	err = c.BodyParser(&notes)
	if err != nil {
		n.log.Error("Failed to parse request body: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{
			"message": "Invalid request body",
		})
	}

	err = n.usecase.UpdateNote(c.Context(), idInt, notes.Notes)
	if err != nil {
		n.log.Error("Failed to update notes: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(apperror.ErrUpdateNotes)
	}

	n.log.Info("Update note completed successfully")
	return c.Status(fiber.StatusOK).JSON(map[string]interface{}{
		"id":   idInt,
		"note": notes.Notes,
	})
}

func (n *noteHandler) FindAll(c *fiber.Ctx) error {
	n.log.Info("Start of note find all notes")

	notesList, err := n.usecase.GetAllNotes(c.Context())
	if err != nil {
		n.log.Error("Failed to find all notes: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(apperror.ErrNoteNotFound)
	}

	n.log.Info("Find all notes completed successfully")
	return c.Status(fiber.StatusOK).JSON(notesList)
}

func (n *noteHandler) convertIdtoInt(c *fiber.Ctx) (int, error) {

	id := c.Params("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		n.log.Error("Can not convert idString to idInt :%v", err)
		return 0, c.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{
			"Message": "Invalid id parameter",
		})
	}
	return idInt, nil
}
