package iam

import (
	"context"
	"errors"

	"github.com/google/uuid"
)

type DB interface {
	User(ctx context.Context, id uuid.UUID) (User, error)
	Search(ctx context.Context, display string) (User, error)
	OnBoard(ctx context.Context, display, name string, age uint) (uuid.UUID, error)
}

type User struct {
	ID uuid.UUID
	// Display name used as an alias
	Display string
	// Name is the fully qualified name of a user
	Name string
	// Age is a non-negative integer value.
	Age uint
}

var (
	ErrNotFound      = errors.New("iam: not found")
	ErrDisplayExists = errors.New("iam: display exists")
)
