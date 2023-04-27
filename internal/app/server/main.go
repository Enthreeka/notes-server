package server

import (
	"context"
	"fmt"

	"github.com/Enthreeka/notes-server/internal/config"
	"github.com/Enthreeka/notes-server/internal/controller/http"
	"github.com/Enthreeka/notes-server/internal/repo"
	"github.com/Enthreeka/notes-server/internal/usecase"
	"github.com/Enthreeka/notes-server/pkg/db"
	"github.com/Enthreeka/notes-server/pkg/logger"
	"github.com/gofiber/fiber/v2"
)

func Run(log *logger.Logger, config *config.Config) error {

	db, err := db.NewConnect(context.Background(), config.Postgres.Url)
	if err != nil {
		return err
	}

	defer db.Close()

	app := fiber.New()

	noteRepo := repo.NewNotesRepository(db, log)
	noteService := usecase.NewNotesService(noteRepo, log)
	noteHandler := http.NewNoteHandler(noteService, log)

	api := app.Group("/api")

	v1 := api.Group("/notes")
	v1.Get("/:id", noteHandler.SearchNode)
	v1.Get("/", noteHandler.FindAll)
	v1.Post("/", noteHandler.CreateNote)
	v1.Delete("/:id", noteHandler.DeleteNotes)
	v1.Patch("/:id", noteHandler.UpdateNote)

	log.Info("Starting http server: %s:%s", config.Server.TypeServer, config.Server.Port)

	if err = app.Listen(fmt.Sprintf(":%s", config.Server.Port)); err != nil {
		log.Fatal("Server listening failed:%s", err)
	}

	return nil
}
