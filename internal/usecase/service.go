package usecase

import (
	"context"

	"github.com/Enthreeka/notes-server/internal/entity"
)

type NoteService interface {
	CreateNotes(ctx context.Context, note entity.Notes) error
	GetNotes(ctx context.Context, id int) (*entity.Notes, error)
	DeleteNotes(ctx context.Context, id int) error
	UpdateNote(ctx context.Context, id int, notes string) error
	GetAllNotes(ctx context.Context) ([]entity.Notes, error)
}
