package usecase

import (
	"context"

	"github.com/Enthreeka/notes-server/internal/apperror"
	"github.com/Enthreeka/notes-server/internal/entity"
	"github.com/Enthreeka/notes-server/internal/repo"
	"github.com/Enthreeka/notes-server/pkg/logger"
)

type noteService struct {
	repo repo.Note
	log  *logger.Logger
}

func NewNotesService(repo repo.Note, log *logger.Logger) NoteService {
	return &noteService{
		repo: repo,
		log:  log,
	}
}

func (n *noteService) CreateNotes(ctx context.Context, note entity.Notes) error {

	err := n.repo.Create(ctx, note)
	if err != nil {
		n.log.Error("Error with create notes %v", err)
		return err
	}

	return nil
}

func (n *noteService) DeleteNotes(ctx context.Context, id int) error {

	err := n.repo.DeleteByID(ctx, id)
	if err != nil {
		n.log.Error("Error with delete notes %v", err)
		return err
	}

	return nil
}

func (n *noteService) GetNotes(ctx context.Context, id int) (*entity.Notes, error) {

	notes, err := n.repo.GetByID(ctx, id)
	if err != nil {
		n.log.Error("Error with get note %v", err)
		return nil, apperror.ErrNoteNotFound
	}

	return notes, nil
}

func (n *noteService) UpdateNote(ctx context.Context, id int, notes string) error {

	err := n.repo.UpdateNote(ctx, id, notes)
	if err != nil {
		n.log.Error("Error with update notes %v", err)
		return err
	}

	return nil
}

func (n *noteService) GetAllNotes(ctx context.Context) ([]entity.Notes, error) {

	notes, err := n.repo.GetAll(ctx)
	if err != nil {
		n.log.Error("Error with get notes %v", err)
		return nil, apperror.ErrNoteNotFound
	}

	return notes, nil
}
