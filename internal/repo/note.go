package repo

import (
	"context"

	"github.com/Enthreeka/notes-server/internal/entity"
	"github.com/Enthreeka/notes-server/pkg/db"
	"github.com/Enthreeka/notes-server/pkg/logger"
)

type noteRepository struct {
	db  *db.Postgres
	log *logger.Logger
}

func NewNotesRepository(db *db.Postgres, log *logger.Logger) Note {
	return &noteRepository{
		db:  db,
		log: log,
	}
}

func (n *noteRepository) Create(ctx context.Context, notes entity.Notes) error {

	query := "INSERT INTO notes (notes) VALUES ($1)"
	_, err := n.db.Pool.Exec(ctx, query, &notes.Notes)
	return err

}

func (n *noteRepository) DeleteByID(ctx context.Context, id int) error {

	query := "DELETE FROM notes WHERE id = $1"
	_, err := n.db.Pool.Exec(ctx, query, id)
	return err
}

func (n *noteRepository) GetByID(ctx context.Context, id int) (*entity.Notes, error) {

	query := "SELECT id , notes, created_at FROM notes WHERE id = $1"
	var notes entity.Notes
	err := n.db.Pool.QueryRow(ctx, query, id).Scan(&notes.ID, &notes.Notes, &notes.CreatedAt)
	if err != nil {
		n.log.Error("QueryRow in GetByID repo failed: %v ", err)
		return nil, err
	}

	return &notes, nil
}

func (n *noteRepository) UpdateNote(ctx context.Context, id int, notes string) error {

	query := "UPDATE notes SET notes = $1 WHERE id = $2"
	_, err := n.db.Pool.Exec(ctx, query, notes, id)
	return err
}

func (n *noteRepository) GetAll(ctx context.Context) ([]entity.Notes, error) {

	query := "SELECT id, notes, created_at FROM notes"
	var notesList []entity.Notes
	rows, err := n.db.Pool.Query(ctx, query)
	if err != nil {
		n.log.Error("QueryRow in GetByID repo failed: %v ", err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var notes entity.Notes
		if err := rows.Scan(&notes.ID, &notes.Notes, &notes.CreatedAt); err != nil {
			n.log.Error("Error scanning row: %v", err)
			return nil, err
		}
		notesList = append(notesList, notes)
	}
	if err := rows.Err(); err != nil {
		n.log.Error("Error after iterating rows: %v", err)
		return nil, err
	}
	
	return notesList, nil
}
