package repo

import (
	"context"

	"github.com/Enthreeka/notes-server/internal/entity"
)

type Note interface {
	Create(ctx context.Context, note entity.Note) error
	Update(ctx context.Context, id string) error
	Delete(ctx context.Context, id string) error
	GetByID(ctx context.Context, id int) (*entity.Note, error)
}
