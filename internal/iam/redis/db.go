package redis

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"github.com/roku-on-it/golang-search/internal/iam"
)

var _ iam.DB = (*DB)(nil)

type DB struct {
	Conn *redis.Client
}

// OnBoard implements iam.DB.
func (d *DB) OnBoard(ctx context.Context, display string, name string, age uint) (uuid.UUID, error) {
	// NOTE be correct and handle this, even if it's highly unlikely to fail
	uid, _ := uuid.NewV7()
	u := &User{
		ID:      uid,
		Display: display,
		Name:    name,
		Age:     age,
	}

	// NOTE redis-server >= 4.0
	err := d.Conn.HSet(ctx, fmt.Sprintf("user:%s", uid), u).Err()
	if err != nil {
		return uuid.Nil, fmt.Errorf("redis: setting new user: %w", err)
	}
	return uid, nil
}

// Search implements iam.DB.
func (d *DB) Search(ctx context.Context, display string) (iam.User, error) {
	panic("unimplemented")
}

// User implements iam.DB.
func (d *DB) User(ctx context.Context, id uuid.UUID) (iam.User, error) {
	// NOTE I'm unsure but think you can use the second arg here for the count.
	k := fmt.Sprintf("user:%s", id)

	var found iam.User
	// NOTE scan directly into a struct
	// https://redis.uptrace.dev/guide/scanning-hash-fields.html
	err := d.Conn.HGetAll(ctx, k).Scan(&found)
	if err != nil {
		return iam.User{}, err
	}
	return found, nil
}

type User struct {
	ID      uuid.UUID `redis:"id"`
	Display string    `redis:"display"`
	Name    string    `redis:"name"`
	Age     uint      `redis:"age"`
}
