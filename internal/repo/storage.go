package repo

import (
	"context"

	"github.com/Enthreeka/notes-server/internal/entity"
)

type Note interface {
	Create(ctx context.Context, note entity.Notes) error
	UpdateNote(ctx context.Context, id int, notes string) error
	DeleteByID(ctx context.Context, id int) error
	GetByID(ctx context.Context, id int) (*entity.Notes, error)
	GetAll(ctx context.Context) ([]entity.Notes, error)
}
