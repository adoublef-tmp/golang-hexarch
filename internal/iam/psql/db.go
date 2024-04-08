package psql

import (
	"context"

	"github.com/google/uuid"
	"github.com/roku-on-it/golang-search/internal/iam"
)

var _ iam.DB = (*DB)(nil)

type DB struct{}

// OnBoard implements iam.DB.
func (d *DB) OnBoard(ctx context.Context, display string, name string, age uint) (uuid.UUID, error) {
	panic("unimplemented")
}

// Search implements iam.DB.
func (d *DB) Search(ctx context.Context, display string) (iam.User, error) {
	panic("unimplemented")
}

// User implements iam.DB.
func (d *DB) User(ctx context.Context, id uuid.UUID) (iam.User, error) {
	panic("unimplemented")
}
